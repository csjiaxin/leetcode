package l864

// https://leetcode.com/problems/shortest-path-to-get-all-keys/
func shortestPathAllKeys(grid []string) int {
	rows, cols := len(grid), len(grid[0])
	start := [3]int{}
	allKeys := 0
	for r, row := range grid {
		for c := range row {
			if row[c] == '@' {
				start = [3]int{r, c, 0}
			} else if row[c] >= 'a' && row[c] <= 'z' {
				allKeys |= 1 << (row[c] - 'a')
			}
		}
	}
	directions := [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	answer := -1
	bfs([][3]int{start}, func(state [3]int) [][3]int {
		next := make([][3]int, 0, 4)
		for _, direction := range directions {
			r, c := state[0]+direction[0], state[1]+direction[1]
			if r < 0 || r >= rows || c < 0 || c >= cols || grid[r][c] == '#' {
				continue
			}
			cell, keys := grid[r][c], state[2]
			if cell >= 'A' && cell <= 'Z' && keys&(1<<(cell-'A')) == 0 {
				continue
			}
			if cell >= 'a' && cell <= 'z' {
				keys |= 1 << (cell - 'a')
			}
			next = append(next, [3]int{r, c, keys})
		}
		return next
	}, func(state [3]int, distance int) bool {
		if state[2] == allKeys {
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
