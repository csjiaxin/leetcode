# Trie：统一走树模板与节点载荷

[返回模板总目录](../algorithm_templates/README.md)

TrieNode[V] 把节点路径与题目数据分离。TrieWalk 的 create 控制建点，onNode 维护沿途统计；Data 分别保存终止标记、次数、候选、前缀和或原单词。查找失败不创建节点，完整单词和存在前缀必须分别判断。

字符 Trie 使用字节 map，适用于本组 ASCII 词典。XOR 子系列单列31位非负整数 BitTrie，支持计数、窗口删除、最大异或和小于阈值的数量；相同“走树”外形不意味着查询规则相同。

## 复制约定

每题 `.go` 文件已内嵌所需模板，无需导入仓库的算法包。文档将固定模板只列一次，逐题展示适配器；每个适配器与它列出的模板组成一个独立代码组。不同题的同名 Constructor、类型或入口请分别使用。将列出的标准库 import 放在文件顶部，再放模板与适配器；复制提交时优先直接使用对应题目的完整 `.go`。

本组使用 Go 1.21+ 的 min/max 和 Go 1.18+ 的泛型。平台已提供的 ListNode 等类型沿用平台定义。

## 固定模板

### TrieWalk

<!-- template: TrieWalk -->
```go
// TrieNode 的 Data 保存题目需要的终止、计数或聚合信息。
type TrieNode[V any] struct {
	Next map[byte]*TrieNode[V]
	Data V
}

func NewTrieNode[V any]() *TrieNode[V] { return &TrieNode[V]{Next: make(map[byte]*TrieNode[V])} }

// TrieWalk 沿前缀走树；create 决定是否建点，onNode 可聚合沿途信息。
func TrieWalk[V any](node *TrieNode[V], key string, create bool, onNode func(*TrieNode[V])) *TrieNode[V] {
	if node == nil {
		return nil
	}
	if onNode != nil {
		onNode(node)
	}
	for i := 0; i < len(key); i++ {
		child := node.Next[key[i]]
		if child == nil {
			if !create {
				return nil
			}
			child = NewTrieNode[V]()
			node.Next[key[i]] = child
		}
		node = child
		if onNode != nil {
			onNode(node)
		}
	}
	return node
}
```

### BitTrie

<!-- template: BitTrie -->
```go
// BitTrie 按高位到低位维护非负31位整数，Count 支持窗口删除。
type bitNode struct {
	Next  [2]*bitNode
	Count int
}
type BitTrie struct{ Root *bitNode }

func NewBitTrie() *BitTrie { return &BitTrie{Root: &bitNode{}} }
func (t *BitTrie) Add(value, delta int) {
	node := t.Root
	node.Count += delta
	for bit := 30; bit >= 0; bit-- {
		digit := (value >> bit) & 1
		if node.Next[digit] == nil {
			node.Next[digit] = &bitNode{}
		}
		node = node.Next[digit]
		node.Count += delta
	}
}
func (t *BitTrie) MaxXor(value int) int {
	if t.Root.Count == 0 {
		return -1
	}
	node := t.Root
	answer := 0
	for bit := 30; bit >= 0; bit-- {
		digit := (value >> bit) & 1
		preferred := digit ^ 1
		if node.Next[preferred] != nil && node.Next[preferred].Count > 0 {
			answer |= 1 << bit
			node = node.Next[preferred]
		} else {
			node = node.Next[digit]
		}
	}
	return answer
}

// Less 统计与 value 异或后严格小于 limit 的已插入数字数量。
func (t *BitTrie) Less(value, limit int) int {
	node := t.Root
	answer := 0
	for bit := 30; bit >= 0 && node != nil; bit-- {
		digit := (value >> bit) & 1
		if (limit>>bit)&1 != 0 {
			if child := node.Next[digit]; child != nil {
				answer += child.Count
			}
			node = node.Next[digit^1]
		} else {
			node = node.Next[digit]
		}
	}
	return answer
}
```

## 题目索引

| 题目 | 模板 | 本地代码与测试 |
| --- | --- | --- |
| [14](https://leetcode.com/problems/longest-common-prefix/) | TrieWalk | [代码](l14/14.go) |
| [139](https://leetcode.com/problems/word-break/) | TrieWalk | [代码](l139/139.go) |
| [208](https://leetcode.com/problems/implement-trie-prefix-tree/) | TrieWalk | [代码](l208/208.go) · [测试](l208/208_test.go) |
| [211](https://leetcode.com/problems/design-add-and-search-words-data-structure/) | TrieWalk | [代码](l211/211.go) · [测试](l211/211_test.go) |
| [212](https://leetcode.com/problems/word-search-ii/) | TrieWalk | [代码](l212/212.go) · [测试](l212/212_test.go) |
| [421](https://leetcode.com/problems/maximum-xor-of-two-numbers-in-an-array/) | BitTrie | [代码](l421/421.go) |
| [472](https://leetcode.com/problems/concatenated-words/) | TrieWalk | [代码](l472/472.go) |
| [648](https://leetcode.com/problems/replace-words/) | TrieWalk | [代码](l648/648.go) · [测试](l648/648_test.go) |
| [677](https://leetcode.com/problems/map-sum-pairs/) | TrieWalk | [代码](l677/677.go) · [测试](l677/677_test.go) |
| [720](https://leetcode.com/problems/longest-word-in-dictionary/) | TrieWalk | [代码](l720/720.go) |
| [745](https://leetcode.com/problems/prefix-and-suffix-search/) | TrieWalk | [代码](l745/745.go) |
| [820](https://leetcode.com/problems/short-encoding-of-words/) | TrieWalk | [代码](l820/820.go) |
| [1032](https://leetcode.com/problems/stream-of-characters/) | TrieWalk | [代码](l1032/1032.go) |
| [1233](https://leetcode.com/problems/remove-sub-folders-from-the-filesystem/) | TrieWalk | [代码](l1233/1233.go) |
| [1268](https://leetcode.com/problems/search-suggestions-system/) | TrieWalk | [代码](l1268/1268.go) · [测试](l1268/1268_test.go) |
| [1707](https://leetcode.com/problems/maximum-xor-with-an-element-from-array/) | BitTrie | [代码](l1707/1707.go) |
| [1803](https://leetcode.com/problems/count-pairs-with-xor-in-a-range/) | BitTrie | [代码](l1803/1803.go) |
| [1804](https://leetcode.com/problems/implement-trie-ii-prefix-tree/) | TrieWalk | [代码](l1804/1804.go) |
| [2416](https://leetcode.com/problems/sum-of-prefix-scores-of-strings/) | TrieWalk | [代码](l2416/2416.go) |
| [2935](https://leetcode.com/problems/maximum-strong-pair-xor-ii/) | BitTrie | [代码](l2935/2935.go) |

## 逐题适配器

### [14](https://leetcode.com/problems/longest-common-prefix/)

需要模板：`TrieWalk`。

<!-- solution: trie/l14/14.go; templates: TrieWalk; imports: none -->
```go
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	root := NewTrieNode[int]()
	for _, word := range strs {
		TrieWalk(root, word, true, func(node *TrieNode[int]) { node.Data++ })
	}
	node := root
	answer := []byte{}
	for len(node.Next) == 1 {
		var value byte
		var child *TrieNode[int]
		for value, child = range node.Next {
		}
		if child.Data != len(strs) {
			break
		}
		answer = append(answer, value)
		node = child
	}
	return string(answer)
}
```

### [139](https://leetcode.com/problems/word-break/)

需要模板：`TrieWalk`。

<!-- solution: trie/l139/139.go; templates: TrieWalk; imports: none -->
```go
func wordBreak(s string, wordDict []string) bool {
	root := NewTrieNode[bool]()
	for _, word := range wordDict {
		if word != "" {
			TrieWalk(root, word, true, nil).Data = true
		}
	}
	possible := make([]bool, len(s)+1)
	possible[0] = true
	for i := 0; i < len(s); i++ {
		if !possible[i] {
			continue
		}
		node := root
		for j := i; j < len(s); j++ {
			node = node.Next[s[j]]
			if node == nil {
				break
			}
			if node.Data {
				possible[j+1] = true
			}
		}
	}
	return possible[len(s)]
}
```

### [208](https://leetcode.com/problems/implement-trie-prefix-tree/)

需要模板：`TrieWalk`。

<!-- solution: trie/l208/208.go; templates: TrieWalk; imports: none -->
```go
type Trie struct{ root *TrieNode[bool] }

func Constructor() Trie            { return Trie{NewTrieNode[bool]()} }
func (t *Trie) Insert(word string) { TrieWalk(t.root, word, true, nil).Data = true }
func (t *Trie) Search(word string) bool {
	node := TrieWalk(t.root, word, false, nil)
	return node != nil && node.Data
}
func (t *Trie) StartsWith(prefix string) bool { return TrieWalk(t.root, prefix, false, nil) != nil }
```

### [211](https://leetcode.com/problems/design-add-and-search-words-data-structure/)

需要模板：`TrieWalk`。

<!-- solution: trie/l211/211.go; templates: TrieWalk; imports: none -->
```go
type WordDictionary struct{ root *TrieNode[bool] }

func Constructor() WordDictionary             { return WordDictionary{NewTrieNode[bool]()} }
func (w *WordDictionary) AddWord(word string) { TrieWalk(w.root, word, true, nil).Data = true }
func (w *WordDictionary) Search(word string) bool {
	var visit func(*TrieNode[bool], int) bool
	visit = func(node *TrieNode[bool], i int) bool {
		if node == nil {
			return false
		}
		if i == len(word) {
			return node.Data
		}
		if word[i] != '.' {
			return visit(node.Next[word[i]], i+1)
		}
		for _, child := range node.Next {
			if visit(child, i+1) {
				return true
			}
		}
		return false
	}
	return visit(w.root, 0)
}
```

### [212](https://leetcode.com/problems/word-search-ii/)

需要模板：`TrieWalk`。

<!-- solution: trie/l212/212.go; templates: TrieWalk; imports: none -->
```go
func findWords(board [][]byte, words []string) []string {
	answer := []string{}
	if len(board) == 0 || len(board[0]) == 0 {
		return answer
	}
	root := NewTrieNode[string]()
	for _, word := range words {
		if word != "" {
			TrieWalk(root, word, true, nil).Data = word
		}
	}
	visited := make([][]bool, len(board))
	for i := range visited {
		visited[i] = make([]bool, len(board[0]))
	}
	var visit func(int, int, *TrieNode[string])
	visit = func(r, c int, node *TrieNode[string]) {
		if r < 0 || r >= len(board) || c < 0 || c >= len(board[0]) || visited[r][c] {
			return
		}
		node = node.Next[board[r][c]]
		if node == nil {
			return
		}
		if node.Data != "" {
			answer = append(answer, node.Data)
			node.Data = ""
		}
		visited[r][c] = true
		for _, d := range [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}} {
			visit(r+d[0], c+d[1], node)
		}
		visited[r][c] = false
	}
	for r := range board {
		for c := range board[r] {
			visit(r, c, root)
		}
	}
	return answer
}
```

### [421](https://leetcode.com/problems/maximum-xor-of-two-numbers-in-an-array/)

需要模板：`BitTrie`。

<!-- solution: trie/l421/421.go; templates: BitTrie; imports: none -->
```go
func findMaximumXOR(nums []int) int {
	trie := NewBitTrie()
	for _, value := range nums {
		trie.Add(value, 1)
	}
	answer := 0
	for _, value := range nums {
		answer = max(answer, trie.MaxXor(value))
	}
	return answer
}
```

### [472](https://leetcode.com/problems/concatenated-words/)

需要模板：`TrieWalk`。

<!-- solution: trie/l472/472.go; templates: TrieWalk; imports: none -->
```go
func findAllConcatenatedWordsInADict(words []string) []string {
	root := NewTrieNode[bool]()
	for _, word := range words {
		if word != "" {
			TrieWalk(root, word, true, nil).Data = true
		}
	}
	answer := []string{}
	for _, word := range words {
		pieces := make([]int, len(word)+1)
		for i := 1; i < len(pieces); i++ {
			pieces[i] = -1
		}
		for i := 0; i < len(word); i++ {
			if pieces[i] < 0 {
				continue
			}
			node := root
			for j := i; j < len(word); j++ {
				node = node.Next[word[j]]
				if node == nil {
					break
				}
				if node.Data {
					pieces[j+1] = max(pieces[j+1], pieces[i]+1)
				}
			}
		}
		if pieces[len(word)] >= 2 {
			answer = append(answer, word)
		}
	}
	return answer
}
```

### [648](https://leetcode.com/problems/replace-words/)

需要模板：`TrieWalk`。

标准库：`strings`。

<!-- solution: trie/l648/648.go; templates: TrieWalk; imports: strings -->
```go
func replaceWords(dictionary []string, sentence string) string {
	root := NewTrieNode[bool]()
	for _, word := range dictionary {
		TrieWalk(root, word, true, nil).Data = true
	}
	words := strings.Split(sentence, " ")
	for i, word := range words {
		node := root
		for j := 0; j < len(word); j++ {
			node = node.Next[word[j]]
			if node == nil {
				break
			}
			if node.Data {
				words[i] = word[:j+1]
				break
			}
		}
	}
	return strings.Join(words, " ")
}
```

### [677](https://leetcode.com/problems/map-sum-pairs/)

需要模板：`TrieWalk`。

<!-- solution: trie/l677/677.go; templates: TrieWalk; imports: none -->
```go
type MapSum struct {
	root   *TrieNode[int]
	values map[string]int
}

func Constructor() MapSum { return MapSum{NewTrieNode[int](), make(map[string]int)} }
func (m *MapSum) Insert(key string, val int) {
	delta := val - m.values[key]
	m.values[key] = val
	TrieWalk(m.root, key, true, func(node *TrieNode[int]) { node.Data += delta })
}
func (m *MapSum) Sum(prefix string) int {
	node := TrieWalk(m.root, prefix, false, nil)
	if node == nil {
		return 0
	}
	return node.Data
}
```

### [720](https://leetcode.com/problems/longest-word-in-dictionary/)

需要模板：`TrieWalk`。

<!-- solution: trie/l720/720.go; templates: TrieWalk; imports: none -->
```go
func longestWord(words []string) string {
	root := NewTrieNode[bool]()
	for _, word := range words {
		TrieWalk(root, word, true, nil).Data = true
	}
	answer := ""
	var visit func(*TrieNode[bool], []byte)
	visit = func(node *TrieNode[bool], prefix []byte) {
		for value, child := range node.Next {
			if !child.Data {
				continue
			}
			prefix = append(prefix, value)
			word := string(prefix)
			if len(word) > len(answer) || (len(word) == len(answer) && word < answer) {
				answer = word
			}
			visit(child, prefix)
			prefix = prefix[:len(prefix)-1]
		}
	}
	visit(root, nil)
	return answer
}
```

### [745](https://leetcode.com/problems/prefix-and-suffix-search/)

需要模板：`TrieWalk`。

<!-- solution: trie/l745/745.go; templates: TrieWalk; imports: none -->
```go
type WordFilter struct{ root *TrieNode[int] }

func Constructor(words []string) WordFilter {
	result := WordFilter{NewTrieNode[int]()}
	for index, word := range words {
		for start := 0; start <= len(word); start++ {
			TrieWalk(result.root, word[start:]+"{"+word, true, func(node *TrieNode[int]) { node.Data = index })
		}
	}
	return result
}
func (w *WordFilter) F(pref, suff string) int {
	node := TrieWalk(w.root, suff+"{"+pref, false, nil)
	if node == nil {
		return -1
	}
	return node.Data
}
```

### [820](https://leetcode.com/problems/short-encoding-of-words/)

需要模板：`TrieWalk`。

<!-- solution: trie/l820/820.go; templates: TrieWalk; imports: none -->
```go
func minimumLengthEncoding(words []string) int {
	root := NewTrieNode[bool]()
	for _, word := range words {
		data := []byte(word)
		for l, r := 0, len(data)-1; l < r; l, r = l+1, r-1 {
			data[l], data[r] = data[r], data[l]
		}
		TrieWalk(root, string(data), true, nil).Data = true
	}
	var visit func(*TrieNode[bool], int) int
	visit = func(node *TrieNode[bool], depth int) int {
		if len(node.Next) == 0 {
			if node.Data {
				return depth + 1
			}
			return 0
		}
		answer := 0
		for _, child := range node.Next {
			answer += visit(child, depth+1)
		}
		return answer
	}
	return visit(root, 0)
}
```

### [1032](https://leetcode.com/problems/stream-of-characters/)

需要模板：`TrieWalk`。

<!-- solution: trie/l1032/1032.go; templates: TrieWalk; imports: none -->
```go
type StreamChecker struct {
	root    *TrieNode[bool]
	history []byte
	maximum int
}

func Constructor(words []string) StreamChecker {
	result := StreamChecker{root: NewTrieNode[bool]()}
	for _, word := range words {
		result.maximum = max(result.maximum, len(word))
		data := []byte(word)
		for l, r := 0, len(data)-1; l < r; l, r = l+1, r-1 {
			data[l], data[r] = data[r], data[l]
		}
		TrieWalk(result.root, string(data), true, nil).Data = true
	}
	return result
}
func (s *StreamChecker) Query(letter byte) bool {
	s.history = append(s.history, letter)
	if len(s.history) > s.maximum {
		s.history = s.history[len(s.history)-s.maximum:]
	}
	node := s.root
	for i := len(s.history) - 1; i >= 0; i-- {
		node = node.Next[s.history[i]]
		if node == nil {
			return false
		}
		if node.Data {
			return true
		}
	}
	return node.Data
}
```

### [1233](https://leetcode.com/problems/remove-sub-folders-from-the-filesystem/)

需要模板：`TrieWalk`。

标准库：`sort`。

<!-- solution: trie/l1233/1233.go; templates: TrieWalk; imports: sort -->
```go
func removeSubfolders(folder []string) []string {
	ordered := append([]string(nil), folder...)
	sort.Strings(ordered)
	root := NewTrieNode[bool]()
	answer := []string{}
	for _, path := range ordered {
		node := root
		subfolder := false
		for i := range path {
			if node.Data && path[i] == '/' {
				subfolder = true
				break
			}
			node = node.Next[path[i]]
			if node == nil {
				break
			}
		}
		if !subfolder {
			answer = append(answer, path)
			TrieWalk(root, path, true, nil).Data = true
		}
	}
	return answer
}
```

### [1268](https://leetcode.com/problems/search-suggestions-system/)

需要模板：`TrieWalk`。

标准库：`sort`。

<!-- solution: trie/l1268/1268.go; templates: TrieWalk; imports: sort -->
```go
func suggestedProducts(products []string, searchWord string) [][]string {
	ordered := append([]string(nil), products...)
	sort.Strings(ordered)
	root := NewTrieNode[[]string]()
	for _, product := range ordered {
		TrieWalk(root, product, true, func(node *TrieNode[[]string]) {
			if len(node.Data) < 3 {
				node.Data = append(node.Data, product)
			}
		})
	}
	answer := make([][]string, len(searchWord))
	node := root
	for i := 0; i < len(searchWord); i++ {
		if node != nil {
			node = node.Next[searchWord[i]]
		}
		answer[i] = []string{}
		if node != nil {
			answer[i] = append(answer[i], node.Data...)
		}
	}
	return answer
}
```

### [1707](https://leetcode.com/problems/maximum-xor-with-an-element-from-array/)

需要模板：`BitTrie`。

标准库：`sort`。

<!-- solution: trie/l1707/1707.go; templates: BitTrie; imports: sort -->
```go
func maximizeXor(nums []int, queries [][]int) []int {
	ordered := append([]int(nil), nums...)
	sort.Ints(ordered)
	type query struct{ x, limit, index int }
	pending := make([]query, len(queries))
	for i, q := range queries {
		pending[i] = query{q[0], q[1], i}
	}
	sort.Slice(pending, func(i, j int) bool { return pending[i].limit < pending[j].limit })
	trie := NewBitTrie()
	answer := make([]int, len(queries))
	next := 0
	for _, q := range pending {
		for next < len(ordered) && ordered[next] <= q.limit {
			trie.Add(ordered[next], 1)
			next++
		}
		answer[q.index] = trie.MaxXor(q.x)
	}
	return answer
}
```

### [1803](https://leetcode.com/problems/count-pairs-with-xor-in-a-range/)

需要模板：`BitTrie`。

<!-- solution: trie/l1803/1803.go; templates: BitTrie; imports: none -->
```go
func countPairs(nums []int, low, high int) int {
	trie := NewBitTrie()
	answer := 0
	for _, value := range nums {
		answer += trie.Less(value, high+1) - trie.Less(value, low)
		trie.Add(value, 1)
	}
	return answer
}
```

### [1804](https://leetcode.com/problems/implement-trie-ii-prefix-tree/)

需要模板：`TrieWalk`。

<!-- solution: trie/l1804/1804.go; templates: TrieWalk; imports: none -->
```go
type wordCounts struct{ Pass, End int }
type Trie struct{ root *TrieNode[wordCounts] }

func Constructor() Trie { return Trie{NewTrieNode[wordCounts]()} }
func (t *Trie) Insert(word string) {
	node := TrieWalk(t.root, word, true, func(node *TrieNode[wordCounts]) { node.Data.Pass++ })
	node.Data.End++
}
func (t *Trie) CountWordsEqualTo(word string) int {
	node := TrieWalk(t.root, word, false, nil)
	if node == nil {
		return 0
	}
	return node.Data.End
}
func (t *Trie) CountWordsStartingWith(prefix string) int {
	node := TrieWalk(t.root, prefix, false, nil)
	if node == nil {
		return 0
	}
	return node.Data.Pass
}
func (t *Trie) Erase(word string) {
	if t.CountWordsEqualTo(word) == 0 {
		return
	}
	node := TrieWalk(t.root, word, false, func(node *TrieNode[wordCounts]) { node.Data.Pass-- })
	node.Data.End--
}
```

### [2416](https://leetcode.com/problems/sum-of-prefix-scores-of-strings/)

需要模板：`TrieWalk`。

<!-- solution: trie/l2416/2416.go; templates: TrieWalk; imports: none -->
```go
func sumPrefixScores(words []string) []int {
	root := NewTrieNode[int]()
	for _, word := range words {
		TrieWalk(root, word, true, func(node *TrieNode[int]) { node.Data++ })
	}
	answer := make([]int, len(words))
	for i, word := range words {
		node := root
		for j := range word {
			node = node.Next[word[j]]
			answer[i] += node.Data
		}
	}
	return answer
}
```

### [2935](https://leetcode.com/problems/maximum-strong-pair-xor-ii/)

需要模板：`BitTrie`。

标准库：`sort`。

<!-- solution: trie/l2935/2935.go; templates: BitTrie; imports: sort -->
```go
func maximumStrongPairXor(nums []int) int {
	ordered := append([]int(nil), nums...)
	sort.Ints(ordered)
	trie := NewBitTrie()
	left, answer := 0, 0
	for _, value := range ordered {
		trie.Add(value, 1)
		for ordered[left]*2 < value {
			trie.Add(ordered[left], -1)
			left++
		}
		answer = max(answer, trie.MaxXor(value))
	}
	return answer
}
```

## 验证与原有推导

```sh
go test ./trie/...
```

题解文件与文档使用相同声明和适配器；修改模板后需要同时验证每题包以及由文档重新组装的代码。原有推导、详细例子和扩展题单保存在 [整理前笔记](legacy/README.before-inline-templates.md)。
