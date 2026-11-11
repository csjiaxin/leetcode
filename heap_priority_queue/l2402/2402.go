package l2402

import (
	"leetcode/heap_priority_queue/internal/pqueue"
	"sort"
)

type room struct {
	end int64
	id  int
}

// mostBooked uses an ID heap for idle rooms and an end-time/ID heap for busy rooms.
// O(m log m + (m+n) log n) time, O(m+n) space. End times may exceed int32.
func mostBooked(n int, meetings [][]int) int {
	ordered := append([][]int(nil), meetings...)
	sort.Slice(ordered, func(i, j int) bool { return ordered[i][0] < ordered[j][0] })
	idle := pqueue.New(func(a, b int) bool { return a < b })
	busy := pqueue.New(func(a, b room) bool {
		if a.end != b.end {
			return a.end < b.end
		}
		return a.id < b.id
	})
	for i := 0; i < n; i++ {
		idle.Push(i)
	}
	counts := make([]int, n)
	for _, m := range ordered {
		start, duration := int64(m[0]), int64(m[1])-int64(m[0])
		for busy.Len() > 0 && busy.Peek().end <= start {
			idle.Push(busy.Pop().id)
		}
		var r room
		if idle.Len() > 0 {
			r = room{start, idle.Pop()}
		} else {
			r = busy.Pop()
		}
		r.end += duration
		counts[r.id]++
		busy.Push(r)
	}
	best := 0
	for i := 1; i < n; i++ {
		if counts[i] > counts[best] {
			best = i
		}
	}
	return best
}
