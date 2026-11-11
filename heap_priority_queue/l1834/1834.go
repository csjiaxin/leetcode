package l1834

import (
	"leetcode/heap_priority_queue/internal/pqueue"
	"sort"
)

type task struct{ arrival, duration, index int }

// getOrder selects available tasks by processing time, then original index.
// O(n log n) time, O(n) space. The clock uses int64 for cumulative durations.
func getOrder(tasks [][]int) []int {
	ordered := make([]task, len(tasks))
	for i, t := range tasks {
		ordered[i] = task{t[0], t[1], i}
	}
	sort.Slice(ordered, func(i, j int) bool { return ordered[i].arrival < ordered[j].arrival })
	q := pqueue.New(func(a, b task) bool {
		if a.duration != b.duration {
			return a.duration < b.duration
		}
		return a.index < b.index
	})
	out := make([]int, 0, len(tasks))
	i := 0
	var now int64
	for i < len(ordered) || q.Len() > 0 {
		if q.Len() == 0 {
			now = max(now, int64(ordered[i].arrival))
		}
		for i < len(ordered) && int64(ordered[i].arrival) <= now {
			q.Push(ordered[i])
			i++
		}
		t := q.Pop()
		now += int64(t.duration)
		out = append(out, t.index)
	}
	return out
}
