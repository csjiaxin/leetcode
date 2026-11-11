package l2558

import (
	"leetcode/heap_priority_queue/internal/pqueue"
	"math"
)

// pickGifts uses a maximum heap. O(n+k log n) time, O(n) space.
func pickGifts(gifts []int, k int) int64 {
	q := pqueue.New(func(a, b int) bool { return a > b }, gifts...)
	for ; k > 0 && q.Len() > 0; k-- {
		x := q.Pop()
		r := int(math.Sqrt(float64(x)))
		// Correct rounding at perfect squares using integer comparisons.
		for r > 0 && r > x/r {
			r--
		}
		for r+1 <= x/(r+1) {
			r++
		}
		q.Push(r)
	}
	var total int64
	for q.Len() > 0 {
		total += int64(q.Pop())
	}
	return total
}
