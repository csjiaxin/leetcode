package l934

// https://leetcode.com/problems/shortest-bridge/
func shortestBridge(grid [][]int) int {
	rows, cols := len(grid), len(grid[0])
	start := [2]int{}
	found := false
	for r := 0; r < rows && !found; r++ {
		for c, value := range grid[r] {
			if value == 1 {
				start, found = [2]int{r, c}, true
				break
			}
		}
	}
	directions := [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	adjacent := func(cell [2]int) [][2]int {
		next := make([][2]int, 0, 4)
		for _, direction := range directions {
			r, c := cell[0]+direction[0], cell[1]+direction[1]
			if r >= 0 && r < rows && c >= 0 && c < cols {
				next = append(next, [2]int{r, c})
			}
		}
		return next
	}
	island := bfs([][2]int{start}, func(cell [2]int) [][2]int {
		next := [][2]int{}
		for _, neighbor := range adjacent(cell) {
			if grid[neighbor[0]][neighbor[1]] == 1 {
				next = append(next, neighbor)
			}
		}
		return next
	}, nil)
	starts := make([][2]int, 0, len(island))
	for cell := range island {
		starts = append(starts, cell)
	}
	answer := -1
	bfs(starts, adjacent, func(cell [2]int, distance int) bool {
		if _, firstIsland := island[cell]; grid[cell[0]][cell[1]] == 1 && !firstIsland {
			answer = distance - 1 // The final land cell is not a flipped sea cell.
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
