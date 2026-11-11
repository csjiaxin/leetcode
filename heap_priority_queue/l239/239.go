package l239

import "leetcode/heap_priority_queue/internal/pqueue"

type entry struct{ value, index int }

// maxSlidingWindow removes expired entries when they reach the heap top.
// Stale interior entries can accumulate: O(n log n) time, O(n) auxiliary space.
func maxSlidingWindow(nums []int, k int) []int {
	out := make([]int, 0)
	if k <= 0 || k > len(nums) {
		return out
	}
	q := pqueue.New(func(a, b entry) bool {
		if a.value != b.value {
			return a.value > b.value
		}
		return a.index > b.index
	})
	for i, v := range nums {
		q.Push(entry{v, i})
		for q.Peek().index <= i-k {
			q.Pop()
		}
		if i >= k-1 {
			out = append(out, q.Peek().value)
		}
	}
	return out
}
