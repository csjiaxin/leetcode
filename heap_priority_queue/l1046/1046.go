package l1046

import "leetcode/heap_priority_queue/internal/pqueue"

// lastStoneWeight smashes the two largest stones. O(n log n) time, O(n) space.
func lastStoneWeight(stones []int) int {
	q := pqueue.New(func(a, b int) bool { return a > b }, stones...)
	for q.Len() > 1 {
		a, b := q.Pop(), q.Pop()
		if a > b {
			q.Push(a - b)
		}
	}
	if q.Len() == 0 {
		return 0
	}
	return q.Pop()
}
