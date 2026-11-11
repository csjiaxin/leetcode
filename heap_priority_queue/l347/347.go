package l347

import "leetcode/heap_priority_queue/internal/pqueue"

type entry struct{ value, count int }

// topKFrequent keeps the k most frequent values. Ties are deterministic.
// O(n+u log k) time, O(u+k) space for u unique values.
func topKFrequent(nums []int, k int) []int {
	counts := make(map[int]int)
	for _, x := range nums {
		counts[x]++
	}
	q := pqueue.New(func(a, b entry) bool {
		if a.count != b.count {
			return a.count < b.count
		}
		return a.value > b.value
	})
	for v, c := range counts {
		q.Push(entry{v, c})
		if q.Len() > k {
			q.Pop()
		}
	}
	out := make([]int, q.Len())
	for i := len(out) - 1; i >= 0; i-- {
		out[i] = q.Pop().value
	}
	return out
}
