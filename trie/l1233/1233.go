package l1233

import "sort"

func removeSubfolders(folder []string) []string {
	ordered := append([]string(nil), folder...)
	sort.Strings(ordered)
	root := NewTrieNode[bool]()
	answer := []string{}
	for _, path := range ordered {
		node := root
		subfolder := false
		for i := range path {
			if node.Data && path[i] == '/' {
				subfolder = true
				break
			}
			node = node.Next[path[i]]
			if node == nil {
				break
			}
		}
		if !subfolder {
			answer = append(answer, path)
			TrieWalk(root, path, true, nil).Data = true
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
