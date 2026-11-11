package l703

import "leetcode/heap_priority_queue/internal/pqueue"

// KthLargest retains only the largest k values seen so far.
type KthLargest struct {
	k int
	q *pqueue.Queue[int]
}

func Constructor(k int, nums []int) KthLargest {
	s := KthLargest{k: k, q: pqueue.New(func(a, b int) bool { return a < b })}
	for _, x := range nums {
		s.Add(x)
	}
	return s
}

// Add takes O(log k) time; total storage is O(k).
func (s *KthLargest) Add(val int) int {
	s.q.Push(val)
	if s.q.Len() > s.k {
		s.q.Pop()
	}
	return s.q.Peek()
}
