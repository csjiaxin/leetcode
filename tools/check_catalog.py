#!/usr/bin/env python3
"""Check restructuring targets, internal Go imports, and selected Markdown links.

Run from any directory:
    python3 tools/check_catalog.py
    python3 tools/check_catalog.py --root . --markdown README.md --markdown dp/README.md

Only the two current entry-point READMEs are checked by default. Historical
notes can be opted in with repeated --markdown arguments; they are not silently
treated as a pristine repository-wide link catalog. This script never changes
files, imports repository Go/Python code, or invokes Git/network services.
"""

from __future__ import annotations

import argparse
import ast
from dataclasses import dataclass, field
import html
import json
import os
from pathlib import Path
import re
import sys
from typing import Iterable
from urllib.parse import unquote, urlsplit


MANIFEST = Path("algorithm_templates/reorganization.json")
DEFAULT_MARKDOWN = ("README.md", "algorithm_templates/README.md")
FENCE = re.compile(r"^[ \t]{0,3}(`{3,}|~{3,})(.*)$")
INLINE_LINK = re.compile(r"!?\[[^\]\n]*\]\(\s*")
REFERENCE_LINK = re.compile(r"(?m)^[ \t]{0,3}\[[^\]\n]+\]:[ \t]*")
SCHEME = re.compile(r"^[A-Za-z][A-Za-z0-9+.-]*:")
IDENTIFIER = re.compile(r"[A-Za-z_][A-Za-z0-9_]*")


@dataclass
class Report:
    mappings: int = 0
    go_files: int = 0
    internal_imports: int = 0
    markdown_files: int = 0
    local_links: int = 0
    errors: list[str] = field(default_factory=list)


def _relative_path(root: Path, value: object) -> Path:
    """Resolve an explicitly repo-relative manifest path without escaping root."""
    if not isinstance(value, str) or not value.strip():
        raise ValueError("path must be a nonempty string")
    path = Path(value)
    if path.is_absolute():
        raise ValueError("path must be repo-relative")
    resolved = (root / path).resolve()
    try:
        resolved.relative_to(root)
    except ValueError as exc:
        raise ValueError("path escapes the repository root") from exc
    return resolved


def _display(root: Path, path: Path) -> str:
    try:
        return path.relative_to(root).as_posix()
    except ValueError:
        return str(path)


def _go_files(root: Path) -> Iterable[Path]:
    # os.walk does not follow directory symlinks. Never inspect Git internals.
    for directory, subdirectories, files in os.walk(root):
        subdirectories[:] = sorted(name for name in subdirectories if name != ".git")
        for name in sorted(files):
            if name.endswith(".go"):
                yield Path(directory) / name


def check_mappings(root: Path, report: Report) -> None:
    manifest = root / MANIFEST
    try:
        data = json.loads(manifest.read_text(encoding="utf-8"))
    except (OSError, UnicodeError, json.JSONDecodeError) as exc:
        report.errors.append(f"{MANIFEST.as_posix()}: cannot read manifest: {exc}")
        return
    if not isinstance(data, dict) or not isinstance(data.get("mappings"), list):
        report.errors.append(f"{MANIFEST.as_posix()}: expected an object with a mappings array")
        return

    mappings: list[tuple[Path, Path]] = []
    for index, entry in enumerate(data["mappings"]):
        if not isinstance(entry, dict):
            report.errors.append(f"{MANIFEST.as_posix()}: mapping {index} must be an object")
            continue
        try:
            source = _relative_path(root, entry.get("from"))
            target = _relative_path(root, entry.get("to"))
        except ValueError as exc:
            report.errors.append(f"{MANIFEST.as_posix()}: mapping {index}: {exc}")
            continue
        mappings.append((source, target))
        report.mappings += 1
        if not target.exists():
            report.errors.append(f"mapping target does not exist: {_display(root, target)}")

    # A canonical source directory can legitimately be reused as a destination
    # or contain new variants. Case-folding also handles case-only renames on
    # case-insensitive macOS filesystems (e.g. Interval -> interval).
    target_parts = [tuple(part.casefold() for part in target.parts) for _, target in mappings]
    for source, _ in mappings:
        source_parts = tuple(part.casefold() for part in source.parts)
        reused = any(parts[: len(source_parts)] == source_parts for parts in target_parts)
        if reused:
            continue
        if source.is_file() and source.suffix == ".go":
            report.errors.append(f"original Go source still exists after move: {_display(root, source)}")
        elif source.is_dir():
            # Notes/history may remain at an old directory; only Go source and
            # tests must have moved. No baseline snapshot is required to check.
            for remaining in _go_files(source):
                report.errors.append(
                    f"original Go source still exists after move: {_display(root, remaining)}"
                )


def go_imports(source: str) -> Iterable[str]:
    """Lex Go tokens enough to read imports, ignoring comments/string contents.

    A plain import regex would incorrectly match examples inside raw strings.
    Broken practice function bodies do not prevent this lexical-only check.
    """
    tokens: list[tuple[str, str]] = []
    i = 0
    while i < len(source):
        if source.startswith("//", i):
            end = source.find("\n", i + 2)
            i = len(source) if end < 0 else end + 1
        elif source.startswith("/*", i):
            end = source.find("*/", i + 2)
            i = len(source) if end < 0 else end + 2
        elif source[i] in ('"', "`", "'"):
            quote, start = source[i], i
            i += 1
            while i < len(source):
                if quote != "`" and source[i] == "\\":
                    i += 2
                    continue
                if source[i] == quote:
                    i += 1
                    break
                i += 1
            literal = source[start:i]
            if quote == "`":
                tokens.append(("string", literal[1:-1].replace("\r", "")))
            elif quote == '"':
                try:
                    value = ast.literal_eval(literal)
                except (SyntaxError, ValueError):
                    value = literal[1:-1]
                tokens.append(("string", value))
            else:
                tokens.append(("character", literal))
        else:
            identifier = IDENTIFIER.match(source, i)
            if identifier:
                tokens.append(("identifier", identifier.group()))
                i = identifier.end()
            else:
                if not source[i].isspace():
                    tokens.append(("punctuation", source[i]))
                i += 1

    i = 0
    while i < len(tokens):
        if tokens[i] != ("identifier", "import"):
            i += 1
            continue
        i += 1
        if i < len(tokens) and tokens[i][1] == "(":
            i += 1
            while i < len(tokens) and tokens[i][1] != ")":
                if tokens[i][0] == "string":
                    yield tokens[i][1]
                i += 1
        else:
            if i < len(tokens) and (tokens[i][0] == "identifier" or tokens[i][1] == "."):
                i += 1  # Optional alias, underscore import, or dot import.
            if i < len(tokens) and tokens[i][0] == "string":
                yield tokens[i][1]
        i += 1


def check_imports(root: Path, report: Report) -> None:
    for source_file in _go_files(root):
        try:
            raw = source_file.read_bytes()
        except OSError as exc:
            report.errors.append(f"{_display(root, source_file)}: cannot read Go source: {exc}")
            continue
        if b"\0" in raw:
            continue  # Do not interpret misnamed binary artifacts as Go source.
        report.go_files += 1
        for imported in go_imports(raw.decode("utf-8", errors="replace")):
            if not imported.startswith("leetcode/"):
                continue
            report.internal_imports += 1
            try:
                target = _relative_path(root, imported[len("leetcode/") :])
            except ValueError as exc:
                report.errors.append(f"{_display(root, source_file)}: invalid import {imported!r}: {exc}")
                continue
            if not target.is_dir():
                report.errors.append(
                    f"{_display(root, source_file)}: internal import target does not exist: {imported}"
                )


def strip_markdown_code(source: str) -> str:
    """Blank fenced/inline code, retaining newlines for accurate link locations."""
    output: list[str] = []
    fence_char, fence_length = "", 0
    for line in source.splitlines(keepends=True):
        match = FENCE.match(line)
        if fence_char:
            if (
                match
                and match[1][0] == fence_char
                and len(match[1]) >= fence_length
                and not match[2].strip()
            ):
                fence_char, fence_length = "", 0
            output.append("\n" if line.endswith("\n") else "")
        elif match:
            fence_char, fence_length = match[1][0], len(match[1])
            output.append("\n" if line.endswith("\n") else "")
        else:
            output.append(line)
    text = "".join(output)
    result: list[str] = []
    i = 0
    while i < len(text):
        if text[i] != "`":
            result.append(text[i])
            i += 1
            continue
        end = i
        while end < len(text) and text[end] == "`":
            end += 1
        closing = re.search(r"(?<!`)`{" + str(end - i) + r"}(?!`)", text[end:])
        if not closing:
            result.append(text[i:end])
            i = end
            continue
        stop = end + closing.end()
        result.append("".join("\n" if ch == "\n" else " " for ch in text[i:stop]))
        i = stop
    return "".join(result)


def _link_destination(text: str, start: int) -> str:
    if start >= len(text):
        return ""
    if text[start] == "<":
        end = text.find(">", start + 1)
        return "" if end < 0 else text[start + 1 : end]
    i, depth = start, 0
    while i < len(text):
        ch = text[i]
        if ch == "\\" and i + 1 < len(text):
            i += 2
            continue
        if ch.isspace() or (ch == ")" and depth == 0):
            break
        if ch == "(":
            depth += 1
        elif ch == ")":
            depth -= 1
        i += 1
    return text[start:i]


def markdown_links(source: str) -> Iterable[tuple[int, str]]:
    text = strip_markdown_code(source)
    for expression in (INLINE_LINK, REFERENCE_LINK):
        for match in expression.finditer(text):
            destination = _link_destination(text, match.end())
            if destination:
                yield text.count("\n", 0, match.start()) + 1, destination


def check_markdown(root: Path, paths: Iterable[str], report: Report) -> None:
    for supplied in paths:
        source_file = Path(supplied)
        if not source_file.is_absolute():
            source_file = root / source_file
        source_file = source_file.resolve()
        try:
            source = source_file.read_text(encoding="utf-8")
        except (OSError, UnicodeError) as exc:
            report.errors.append(f"{_display(root, source_file)}: cannot read Markdown: {exc}")
            continue
        report.markdown_files += 1
        for line, destination in markdown_links(source):
            destination = html.unescape(re.sub(r"\\([!\"#$%&'()*+,\-./:;<=>?@\[\]^_`{|}~])", r"\1", destination))
            if destination.startswith("//") or SCHEME.match(destination):
                continue
            try:
                local = unquote(urlsplit(destination).path)
            except ValueError as exc:
                report.errors.append(f"{_display(root, source_file)}:{line}: invalid link {destination!r}: {exc}")
                continue
            if not local:
                continue  # Same-page fragment/query; fragment existence is not checked.
            report.local_links += 1
            target = root / local.lstrip("/") if local.startswith("/") else source_file.parent / local
            if not target.exists():
                report.errors.append(
                    f"{_display(root, source_file)}:{line}: local link target does not exist: {destination}"
                )


def run_checks(root: Path, markdown_paths: Iterable[str] = DEFAULT_MARKDOWN) -> Report:
    root = root.resolve()
    report = Report()
    if not root.is_dir():
        report.errors.append(f"repository root does not exist: {root}")
        return report
    check_mappings(root, report)
    check_imports(root, report)
    check_markdown(root, markdown_paths, report)
    return report


def main(argv: list[str] | None = None) -> int:
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("--root", type=Path, default=Path(__file__).resolve().parents[1])
    parser.add_argument(
        "--markdown", action="append", metavar="PATH",
        help="Markdown file to check; repeat as needed (defaults: README.md and algorithm_templates/README.md)",
    )
    args = parser.parse_args(argv)
    report = run_checks(args.root, args.markdown if args.markdown is not None else DEFAULT_MARKDOWN)
    print(
        f"Checked {report.mappings} mappings, {report.go_files} Go files, "
        f"{report.internal_imports} internal import references, "
        f"{report.markdown_files} Markdown files, {report.local_links} local links."
    )
    for error in report.errors:
        print(f"ERROR: {error}")
    print(f"Catalog check {'FAILED' if report.errors else 'passed'}: {len(report.errors)} error(s).")
    return 1 if report.errors else 0


if __name__ == "__main__":
    sys.exit(main())
