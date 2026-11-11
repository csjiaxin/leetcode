package l630

import (
	"leetcode/heap_priority_queue/internal/pqueue"
	"sort"
)

// scheduleCourse sorts deadlines; when over budget, discard the longest course.
// O(n log n) time, O(n) space. The input is not reordered.
func scheduleCourse(courses [][]int) int {
	ordered := append([][]int(nil), courses...)
	sort.Slice(ordered, func(i, j int) bool { return ordered[i][1] < ordered[j][1] })
	q := pqueue.New(func(a, b int) bool { return a > b })
	var elapsed int64
	for _, c := range ordered {
		q.Push(c[0])
		elapsed += int64(c[0])
		if elapsed > int64(c[1]) {
			elapsed -= int64(q.Pop())
		}
	}
	return q.Len()
}
