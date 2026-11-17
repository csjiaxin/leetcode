package l820

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
