package l215

import "leetcode/heap_priority_queue/internal/pqueue"

// findKthLargest keeps the largest k values in a minimum heap.
// O(n log k) time, O(k) space. Requires 1 <= k <= len(nums).
func findKthLargest(nums []int, k int) int {
	q := pqueue.New(func(a, b int) bool { return a < b })
	for _, x := range nums {
		q.Push(x)
		if q.Len() > k {
			q.Pop()
		}
	}
	return q.Peek()
}
