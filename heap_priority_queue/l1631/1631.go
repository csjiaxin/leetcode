package l1631

import "leetcode/heap_priority_queue/internal/pqueue"

type cell struct{ row, col, effort int }

// minimumEffortPath minimizes the maximum edge difference on a path.
// O(mn log(mn)) time, O(mn) space using Dijkstra.
func minimumEffortPath(heights [][]int) int {
	m, n := len(heights), len(heights[0])
	inf := int(^uint(0) >> 1)
	dist := make([][]int, m)
	for r := range dist {
		dist[r] = make([]int, n)
		for c := range dist[r] {
			dist[r][c] = inf
		}
	}
	dist[0][0] = 0
	q := pqueue.New(func(a, b cell) bool { return a.effort < b.effort }, cell{0, 0, 0})
	dirs := [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	for q.Len() > 0 {
		s := q.Pop()
		if s.effort != dist[s.row][s.col] {
			continue
		}
		if s.row == m-1 && s.col == n-1 {
			return s.effort
		}
		for _, d := range dirs {
			r, c := s.row+d[0], s.col+d[1]
			if r < 0 || r >= m || c < 0 || c >= n {
				continue
			}
			difference := heights[r][c] - heights[s.row][s.col]
			if difference < 0 {
				difference = -difference
			}
			effort := max(s.effort, difference)
			if effort < dist[r][c] {
				dist[r][c] = effort
				q.Push(cell{r, c, effort})
			}
		}
	}
	return 0
}
