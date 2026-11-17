package l14

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
