package l1020

// https://leetcode.com/problems/number-of-enclaves/
func numEnclaves(grid [][]int) int {
	rows, cols := len(grid), len(grid[0])
	starts := [][2]int{}
	land := 0
	for r := range grid {
		for c, value := range grid[r] {
			if value != 1 {
				continue
			}
			land++
			if r == 0 || r == rows-1 || c == 0 || c == cols-1 {
				starts = append(starts, [2]int{r, c})
			}
		}
	}
	directions := [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	escaped := bfs(starts, func(cell [2]int) [][2]int {
		next := make([][2]int, 0, 4)
		for _, direction := range directions {
			r, c := cell[0]+direction[0], cell[1]+direction[1]
			if r >= 0 && r < rows && c >= 0 && c < cols && grid[r][c] == 1 {
				next = append(next, [2]int{r, c})
			}
		}
		return next
	}, nil)
	return land - len(escaped)
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
