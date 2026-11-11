package l1753

import "leetcode/heap_priority_queue/internal/pqueue"

// maximumScore always pairs the two largest nonempty piles.
// O(a+b+c) time, O(1) space (the heap has at most three elements).
func maximumScore(a, b, c int) int {
	q := pqueue.New(func(a, b int) bool { return a > b }, a, b, c)
	score := 0
	for q.Len() > 1 {
		x, y := q.Pop(), q.Pop()
		if y == 0 {
			break
		}
		score++
		if x > 1 {
			q.Push(x - 1)
		}
		if y > 1 {
			q.Push(y - 1)
		}
	}
	return score
}
