package l1032

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
