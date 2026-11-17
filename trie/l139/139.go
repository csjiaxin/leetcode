package l139

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
