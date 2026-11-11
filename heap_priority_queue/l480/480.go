package l480

import "leetcode/heap_priority_queue/internal/pqueue"

type entry struct{ value, index int }

// medianSlidingWindow uses two heaps, index-based lazy expiration, and logical sizes.
// Stale interior entries may remain: O(n log n) worst-case time, O(n) space.
func medianSlidingWindow(nums []int, k int) []float64 {
	out := make([]float64, 0)
	if k <= 0 || k > len(nums) {
		return out
	}
	lower := pqueue.New(func(a, b entry) bool {
		if a.value != b.value {
			return a.value > b.value
		}
		return a.index > b.index
	})
	upper := pqueue.New(func(a, b entry) bool {
		if a.value != b.value {
			return a.value < b.value
		}
		return a.index < b.index
	})
	inLower := make([]bool, len(nums))
	lowSize, upSize, left := 0, 0, 0
	prune := func(q *pqueue.Queue[entry]) {
		for q.Len() > 0 && q.Peek().index < left {
			q.Pop()
		}
	}
	balance := func() {
		prune(lower)
		prune(upper)
		for lowSize > upSize+1 {
			x := lower.Pop()
			lowSize--
			upSize++
			inLower[x.index] = false
			upper.Push(x)
			prune(lower)
		}
		for lowSize < upSize {
			x := upper.Pop()
			upSize--
			lowSize++
			inLower[x.index] = true
			lower.Push(x)
			prune(upper)
		}
	}
	for i, v := range nums {
		if i >= k {
			old := i - k
			if inLower[old] {
				lowSize--
			} else {
				upSize--
			}
			left = old + 1
			balance()
		}
		x := entry{v, i}
		if lowSize == 0 || v <= lower.Peek().value {
			lower.Push(x)
			inLower[i] = true
			lowSize++
		} else {
			upper.Push(x)
			upSize++
		}
		balance()
		if i >= k-1 {
			median := float64(lower.Peek().value)
			if k%2 == 0 {
				median = median/2 + float64(upper.Peek().value)/2
			}
			out = append(out, median)
		}
	}
	return out
}
