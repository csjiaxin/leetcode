package l871

import "leetcode/heap_priority_queue/internal/pqueue"

// minRefuelStops retroactively selects the largest passed station when needed.
// O(n log n) time, O(n) space. Stations must be sorted by position.
func minRefuelStops(target, startFuel int, stations [][]int) int {
	q := pqueue.New(func(a, b int) bool { return a > b })
	reach := int64(startFuel)
	i, stops := 0, 0
	for reach < int64(target) {
		for i < len(stations) && int64(stations[i][0]) <= reach {
			q.Push(stations[i][1])
			i++
		}
		if q.Len() == 0 {
			return -1
		}
		reach += int64(q.Pop())
		stops++
	}
	return stops
}
