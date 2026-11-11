package l973

import "leetcode/heap_priority_queue/internal/pqueue"

type point struct {
	coords   []int
	distance int64
}

// kClosest keeps k smallest squared distances. O(n log k) time, O(k) heap space.
func kClosest(points [][]int, k int) [][]int {
	q := pqueue.New(func(a, b point) bool { return a.distance > b.distance })
	for _, p := range points {
		x, y := int64(p[0]), int64(p[1])
		q.Push(point{p, x*x + y*y})
		if q.Len() > k {
			q.Pop()
		}
	}
	out := make([][]int, q.Len())
	for i := range out {
		out[i] = append([]int(nil), q.Pop().coords...)
	}
	return out
}
