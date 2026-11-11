package l632

import "leetcode/heap_priority_queue/internal/pqueue"

type item struct{ value, row, col int }

// smallestRange advances the current minimum while tracking the maximum.
// O(N log k) time, O(k) space for N total values and k lists.
func smallestRange(nums [][]int) []int {
	q := pqueue.New(func(a, b item) bool { return a.value < b.value })
	high := nums[0][0]
	for r, row := range nums {
		q.Push(item{row[0], r, 0})
		high = max(high, row[0])
	}
	lo, hi := q.Peek().value, high
	for {
		x := q.Pop()
		width, best := int64(high)-int64(x.value), int64(hi)-int64(lo)
		if width < best || width == best && x.value < lo {
			lo, hi = x.value, high
		}
		if x.col+1 == len(nums[x.row]) {
			break
		}
		v := nums[x.row][x.col+1]
		high = max(high, v)
		q.Push(item{v, x.row, x.col + 1})
	}
	return []int{lo, hi}
}
