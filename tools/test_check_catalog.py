"""Focused, isolated tests for the stdlib-only catalog checker."""

from contextlib import redirect_stdout
import io
import json
from pathlib import Path
import tempfile
import unittest

try:
    from . import check_catalog
except ImportError:  # Support direct execution and unittest discovery with -s tools.
    import check_catalog


class CatalogTests(unittest.TestCase):
    def setUp(self) -> None:
        self.temporary = tempfile.TemporaryDirectory()
        self.addCleanup(self.temporary.cleanup)
        self.root = Path(self.temporary.name).resolve()
        self.write("README.md", "# Catalog\n")
        self.write("algorithm_templates/README.md", "# Templates\n")
        self.manifest([])

    def write(self, name: str, content: str) -> None:
        path = self.root / name
        path.parent.mkdir(parents=True, exist_ok=True)
        path.write_text(content, encoding="utf-8")

    def manifest(self, mappings: list[dict[str, str]]) -> None:
        self.write("algorithm_templates/reorganization.json", json.dumps({"mappings": mappings}))

    def test_valid_links_imports_and_ignored_code(self) -> None:
        self.write("canonical/solution.go", "package canonical\n")
        self.write("docs/a b.md", "# Exists\n")
        self.write("docs/name(with parens).md", "# Parenthesized\n")
        self.manifest([{"from": "old", "to": "canonical"}])
        self.write("README.md", """[directory](docs/)
[encoded](<docs/a%20b.md#part>)
[query](docs/a%20b.md?view=raw#part)
[root relative](/docs/a%20b.md)
[parenthesized](<docs/name(with parens).md>)
[anchor](#missing-anchor)
[web](https://example.invalid/missing)
[mail](mailto:nobody@example.invalid)
[protocol relative](//example.invalid/missing)
`[inline code](missing-inline.md)`
``[code with ` tick](missing-double.md)``
```markdown
[fenced](missing-fence.md)
```
~~~markdown
[tilde fenced](missing-tilde.md)
~~~
[reference]: docs/a%20b.md "Title"
""")
        self.write("client/client.go", '''package client
// import "leetcode/missing-comment"
/* import "leetcode/missing-block" */
import (
    alias "leetcode/canonical"
    _ "example.invalid/external"
)
const example = `
import "leetcode/missing-raw-string"
`
const another = "import \\\"leetcode/missing-quoted-string\\\""
''')
        self.write(".git/misleading.go", 'import "leetcode/missing-git"')
        (self.root / "binary.go").write_bytes(b'\0import "leetcode/missing-binary"')
        # This old, unrelated note is deliberately outside the default scope.
        self.write("legacy.md", "[outdated](missing-history.md)")

        report = check_catalog.run_checks(self.root)
        self.assertEqual(report.errors, [])
        self.assertEqual(report.mappings, 1)
        self.assertEqual(report.go_files, 2)
        self.assertEqual(report.internal_imports, 1)
        self.assertEqual(report.markdown_files, 2)
        self.assertEqual(report.local_links, 6)

    def test_broken_link_and_import_are_reported(self) -> None:
        self.write("README.md", "# Index\n[broken](missing.md#anchor)\n")
        self.write("solution.go", 'package example\nimport . "leetcode/missing_package"\n')
        report = check_catalog.run_checks(self.root)
        self.assertEqual(len(report.errors), 2)
        self.assertTrue(any("README.md:2:" in error for error in report.errors))
        self.assertTrue(any("leetcode/missing_package" in error for error in report.errors))

    def test_missing_mapping_target_and_residual_go_source(self) -> None:
        self.manifest([{"from": "old", "to": "new"}])
        self.write("old/solution_test.go", "package old\n")
        self.write("old/history.md", "Retained notes are allowed.\n")
        report = check_catalog.run_checks(self.root)
        self.assertEqual(len(report.errors), 2)
        self.assertTrue(any("mapping target" in error for error in report.errors))
        self.assertTrue(any("old/solution_test.go" in error for error in report.errors))

    def test_reused_source_directory_is_not_required_to_disappear(self) -> None:
        self.manifest([{"from": "canonical", "to": "canonical/variants/original"}])
        self.write("canonical/solution.go", "package canonical\n")
        self.write("canonical/variants/original/solution.go", "package original\n")
        report = check_catalog.run_checks(self.root)
        self.assertEqual(report.errors, [])

    def test_case_only_rename_is_not_a_residual_source(self) -> None:
        self.manifest([{"from": "Interval", "to": "interval"}])
        self.write("Interval/source.go", "package interval\n")
        (self.root / "interval").mkdir(exist_ok=True)
        report = check_catalog.run_checks(self.root)
        self.assertEqual(report.errors, [])

    def test_reference_links_and_repeatable_markdown_scope(self) -> None:
        self.write("docs/one.md", "[ref]: missing-reference.md\n")
        self.write("docs/two.md", "[other](missing-other.md)\n")
        output = io.StringIO()
        with redirect_stdout(output):
            result = check_catalog.main([
                "--root", str(self.root),
                "--markdown", "docs/one.md", "--markdown", "docs/two.md",
            ])
        self.assertEqual(result, 1)
        self.assertIn("2 Markdown files", output.getvalue())
        self.assertIn("2 error(s)", output.getvalue())
        self.assertIn("missing-reference.md", output.getvalue())
        self.assertIn("missing-other.md", output.getvalue())

    def test_manifest_schema_and_escaping_paths_are_errors(self) -> None:
        self.manifest([{"from": "old", "to": "../outside"}])
        report = check_catalog.run_checks(self.root)
        self.assertTrue(any("escapes" in error for error in report.errors))
        self.write("algorithm_templates/reorganization.json", '{"mappings": {}}')
        report = check_catalog.run_checks(self.root)
        self.assertTrue(any("mappings array" in error for error in report.errors))

    def test_single_alias_dot_and_raw_import_forms(self) -> None:
        source = '''package p
import "leetcode/one"
import _ "leetcode/two"
import . `leetcode/three`
import alias "external.example/ignored"
'''
        self.assertEqual(list(check_catalog.go_imports(source)), [
            "leetcode/one", "leetcode/two", "leetcode/three", "external.example/ignored",
        ])


if __name__ == "__main__":
    unittest.main()
