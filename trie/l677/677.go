package l677

type wordNode struct {
	children [26]*wordNode
	isEnd    bool
	val      int
}

// https://leetcode.com/problems/map-sum-pairs/
type MapSum struct {
	root *wordNode
}

func Constructor() MapSum {
	return MapSum{
		root: &wordNode{},
	}
}

func (m *MapSum) Insert(key string, val int) {
	node := m.root
	for i := 0; i < len(key); i++ {
		c := key[i] - 'a'
		if node.children[c] == nil {
			node.children[c] = &wordNode{}
		}
		node = node.children[c]
	}
	node.isEnd = true
	node.val = val
}

func (m *MapSum) Sum(prefix string) int {
	node := m.root
	for _, c := range prefix {
		if node != nil {
			node = node.children[c-'a']
		} else {
			return 0
		}
	}
	ret := 0
	var dfs func(c *wordNode)
	dfs = func(c *wordNode) {
		if c == nil {
			return
		}
		ret += c.val
		for _, child := range c.children {
			if child != nil {
				dfs(child)
			}
		}
	}
	dfs(node)
	return ret
}
