package l1268

import "sort"

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
