package l542

// https://leetcode.com/problems/01-matrix/
func updateMatrix(mat [][]int) [][]int {
	if len(mat) == 0 || len(mat[0]) == 0 {
		return mat
	}
	rows, cols := len(mat), len(mat[0])
	starts := [][2]int{}
	for r := range mat {
		for c, value := range mat[r] {
			if value == 0 {
				starts = append(starts, [2]int{r, c})
			}
		}
	}
	directions := [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	dist := bfs(starts, func(cell [2]int) [][2]int {
		next := make([][2]int, 0, 4)
		for _, direction := range directions {
			r, c := cell[0]+direction[0], cell[1]+direction[1]
			if r >= 0 && r < rows && c >= 0 && c < cols {
				next = append(next, [2]int{r, c})
			}
		}
		return next
	}, nil)
	answer := make([][]int, rows)
	for r := range answer {
		answer[r] = make([]int, cols)
		for c := range answer[r] {
			answer[r][c] = dist[[2]int{r, c}]
		}
	}
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
