from pathlib import Path
import re, subprocess
root=Path('dp/game')
s=(root/'readme.md').read_text()
blocks=re.findall(r'```go\n(.*?)```',s,re.S)
common=blocks[0]
def normalized(text):
    out=subprocess.run(['gofmt'],input=text,text=True,capture_output=True,check=True).stdout
    return re.sub(r'\s+','',re.sub(r'//[^\n]*','',out))
templates=[]
for p in root.glob('l*/*.go'):
    if p.name.endswith('_test.go'):continue
    text=p.read_text();start=text.index('func MemoGame[');end=text.index('\n}',start)+2
    templates.append(normalized(text[start:end]))
assert len(templates)==7 and len(set(templates))==1
assert templates[0]==normalized(common)
for block in blocks[1:]:
    name=re.search(r'func (\w+)',block).group(1)
    assert any(normalized(block) in normalized(p.read_text()) for p in root.glob('l*/*.go') if not p.name.endswith('_test.go')),name
text=re.sub(r'```.*?```','',s,flags=re.S)
for link in re.findall(r'\]\(([^)]+)\)',text):
    if link.startswith('#'):assert f'id="{link[1:]}"' in text,link
    elif not link.startswith('http'):assert (root/link).exists(),link
assert s.count('```')%2==0
print('Seven identical templates; README Go snippets match source; all local links and anchors valid.')
