package l1962

import "leetcode/heap_priority_queue/internal/pqueue"

// minStoneSum removes floor(pile/2) from the largest pile each time.
// O(n+k log n) time, O(n) space.
func minStoneSum(piles []int, k int) int {
	q := pqueue.New(func(a,b int) bool { return a>b }, piles...)
	total := 0
	for _, x := range piles { total += x }
	for ; k>0 && q.Len()>0; k-- {
		x := q.Pop()
		total -= x/2
		q.Push(x-x/2)
	}
	return total
}
