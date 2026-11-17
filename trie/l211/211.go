package l211

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
