package l720

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
