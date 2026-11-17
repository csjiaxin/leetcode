package l1293

// https://leetcode.com/problems/shortest-path-in-a-grid-with-obstacles-elimination/
func shortestPath(grid [][]int, k int) int {
	rows, cols := len(grid), len(grid[0])
	directions := [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	answer := -1
	// The remaining budget belongs in the state; position alone is insufficient.
	bfs([][3]int{{0, 0, k}}, func(state [3]int) [][3]int {
		next := make([][3]int, 0, 4)
		for _, direction := range directions {
			r, c := state[0]+direction[0], state[1]+direction[1]
			if r < 0 || r >= rows || c < 0 || c >= cols {
				continue
			}
			remaining := state[2] - grid[r][c]
			if remaining >= 0 {
				next = append(next, [3]int{r, c, remaining})
			}
		}
		return next
	}, func(state [3]int, distance int) bool {
		if state[0] == rows-1 && state[1] == cols-1 {
			answer = distance
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
