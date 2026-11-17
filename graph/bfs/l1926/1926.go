package l1926

// https://leetcode.com/problems/nearest-exit-from-entrance-in-maze/
func nearestExit(maze [][]byte, entrance []int) int {
	rows, cols := len(maze), len(maze[0])
	start := [2]int{entrance[0], entrance[1]}
	directions := [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	answer := -1
	bfs([][2]int{start}, func(cell [2]int) [][2]int {
		next := make([][2]int, 0, 4)
		for _, direction := range directions {
			r, c := cell[0]+direction[0], cell[1]+direction[1]
			if r >= 0 && r < rows && c >= 0 && c < cols && maze[r][c] == '.' {
				next = append(next, [2]int{r, c})
			}
		}
		return next
	}, func(cell [2]int, distance int) bool {
		if cell != start && (cell[0] == 0 || cell[0] == rows-1 || cell[1] == 0 || cell[1] == cols-1) {
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
