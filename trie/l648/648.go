package l648

import "strings"

func replaceWords(dictionary []string, sentence string) string {
	root := NewTrieNode[bool]()
	for _, word := range dictionary {
		TrieWalk(root, word, true, nil).Data = true
	}
	words := strings.Split(sentence, " ")
	for i, word := range words {
		node := root
		for j := 0; j < len(word); j++ {
			node = node.Next[word[j]]
			if node == nil {
				break
			}
			if node.Data {
				words[i] = word[:j+1]
				break
			}
		}
	}
	return strings.Join(words, " ")
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
