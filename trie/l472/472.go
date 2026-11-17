package l472

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
