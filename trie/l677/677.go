package l677

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
