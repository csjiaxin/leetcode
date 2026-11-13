package l211

type wordNode struct {
	children [26]*wordNode
	isEnd    bool
}

// https://leetcode.com/problems/design-add-and-search-words-data-structure/
type WordDictionary struct {
	root *wordNode
}

func Constructor() WordDictionary {
	return WordDictionary{
		root: &wordNode{},
	}
}

func (d *WordDictionary) AddWord(word string) {
	node := d.root
	for i := 0; i < len(word); i++ {
		c := word[i] - 'a'
		if node.children[c] == nil {
			node.children[c] = &wordNode{}
		}
		node = node.children[c]
	}
	node.isEnd = true
}

func (d *WordDictionary) Search(word string) bool {
	N := len(word)
	var dfs func(*wordNode, int) bool
	dfs = func(node *wordNode, i int) bool {
		if node == nil {
			return false
		}
		if i == N {
			return node.isEnd
		}
		c := word[i]
		if c != '.' {
			return dfs(node.children[c-'a'], i+1)
		}

		for _, child := range node.children {
			if child != nil && dfs(child, i+1) {
				return true
			}
		}
		return false
	}

	return dfs(d.root, 0)
}
