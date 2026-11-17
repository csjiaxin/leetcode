package l1804

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
