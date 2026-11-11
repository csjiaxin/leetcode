package l1642

import "leetcode/heap_priority_queue/internal/pqueue"

// furthestBuilding reserves ladders for the largest climbs seen so far.
// O(n log(ladders+1)) time, O(ladders+1) heap space.
func furthestBuilding(heights []int, bricks, ladders int) int {
	q := pqueue.New(func(a, b int) bool { return a < b })
	for i := 1; i < len(heights); i++ {
		d := heights[i] - heights[i-1]
		if d <= 0 {
			continue
		}
		q.Push(d)
		if q.Len() > ladders {
			bricks -= q.Pop()
		}
		if bricks < 0 {
			return i - 1
		}
	}
	return len(heights) - 1
}
