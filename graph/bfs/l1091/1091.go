package l1091

// https://leetcode.com/problems/shortest-path-in-binary-matrix/
func shortestPathBinaryMatrix(grid [][]int) int {
	n := len(grid)
	if n == 0 || len(grid[0]) == 0 || grid[0][0] != 0 || grid[n-1][n-1] != 0 {
		return -1
	}
	answer := -1
	bfs([][2]int{{0, 0}}, func(cell [2]int) [][2]int {
		next := make([][2]int, 0, 8)
		for dr := -1; dr <= 1; dr++ {
			for dc := -1; dc <= 1; dc++ {
				r, c := cell[0]+dr, cell[1]+dc
				if (dr != 0 || dc != 0) && r >= 0 && r < n && c >= 0 && c < n && grid[r][c] == 0 {
					next = append(next, [2]int{r, c})
				}
			}
		}
		return next
	}, func(cell [2]int, distance int) bool {
		if cell == [2]int{n - 1, n - 1} {
			answer = distance + 1 // This problem counts cells, including the start.
			return true
		}
		return false
	})
	return answer
}

// bfs visits each complete state once, at its minimum distance from any start.
// A true stop result ends the search before expanding that state; nil explores all.
func bfs[S comparable](starts []S, neighbors func(S) []S, stop func(S, int) bool) map[S]int {
	dist := make(map[S]int, len(starts))
	queue := make([]S, 0, len(starts))
	for _, start := range starts {
		if _, seen := dist[start]; seen {
			continue
		}
		dist[start] = 0
		queue = append(queue, start)
	}
	for head := 0; head < len(queue); head++ {
		current := queue[head]
		distance := dist[current]
		if stop != nil && stop(current, distance) {
			break
		}
		for _, next := range neighbors(current) {
			if _, seen := dist[next]; seen {
				continue
			}
			dist[next] = distance + 1
			queue = append(queue, next)
		}
	}
	return dist
}
