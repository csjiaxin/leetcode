package l1707

import "sort"

func maximizeXor(nums []int, queries [][]int) []int {
	ordered := append([]int(nil), nums...)
	sort.Ints(ordered)
	type query struct{ x, limit, index int }
	pending := make([]query, len(queries))
	for i, q := range queries {
		pending[i] = query{q[0], q[1], i}
	}
	sort.Slice(pending, func(i, j int) bool { return pending[i].limit < pending[j].limit })
	trie := NewBitTrie()
	answer := make([]int, len(queries))
	next := 0
	for _, q := range pending {
		for next < len(ordered) && ordered[next] <= q.limit {
			trie.Add(ordered[next], 1)
			next++
		}
		answer[q.index] = trie.MaxXor(q.x)
	}
	return answer
}

// BitTrie 按高位到低位维护非负31位整数，Count 支持窗口删除。
type bitNode struct {
	Next  [2]*bitNode
	Count int
}
type BitTrie struct{ Root *bitNode }

func NewBitTrie() *BitTrie { return &BitTrie{Root: &bitNode{}} }
func (t *BitTrie) Add(value, delta int) {
	node := t.Root
	node.Count += delta
	for bit := 30; bit >= 0; bit-- {
		digit := (value >> bit) & 1
		if node.Next[digit] == nil {
			node.Next[digit] = &bitNode{}
		}
		node = node.Next[digit]
		node.Count += delta
	}
}
func (t *BitTrie) MaxXor(value int) int {
	if t.Root.Count == 0 {
		return -1
	}
	node := t.Root
	answer := 0
	for bit := 30; bit >= 0; bit-- {
		digit := (value >> bit) & 1
		preferred := digit ^ 1
		if node.Next[preferred] != nil && node.Next[preferred].Count > 0 {
			answer |= 1 << bit
			node = node.Next[preferred]
		} else {
			node = node.Next[digit]
		}
	}
	return answer
}

// Less 统计与 value 异或后严格小于 limit 的已插入数字数量。
func (t *BitTrie) Less(value, limit int) int {
	node := t.Root
	answer := 0
	for bit := 30; bit >= 0 && node != nil; bit-- {
		digit := (value >> bit) & 1
		if (limit>>bit)&1 != 0 {
			if child := node.Next[digit]; child != nil {
				answer += child.Count
			}
			node = node.Next[digit^1]
		} else {
			node = node.Next[digit]
		}
	}
	return answer
}
