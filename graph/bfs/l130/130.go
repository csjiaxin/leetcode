package l130

// https://leetcode.com/problems/surrounded-regions/
func solve(board [][]byte) {
	if len(board) == 0 || len(board[0]) == 0 {
		return
	}
	rows, cols := len(board), len(board[0])
	starts := [][2]int{}
	for r := range board {
		for c, value := range board[r] {
			if value == 'O' && (r == 0 || r == rows-1 || c == 0 || c == cols-1) {
				starts = append(starts, [2]int{r, c})
			}
		}
	}
	directions := [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	safe := bfs(starts, func(cell [2]int) [][2]int {
		next := make([][2]int, 0, 4)
		for _, direction := range directions {
			r, c := cell[0]+direction[0], cell[1]+direction[1]
			if r >= 0 && r < rows && c >= 0 && c < cols && board[r][c] == 'O' {
				next = append(next, [2]int{r, c})
			}
		}
		return next
	}, nil)
	for r := range board {
		for c, value := range board[r] {
			if _, escaped := safe[[2]int{r, c}]; value == 'O' && !escaped {
				board[r][c] = 'X'
			}
		}
	}
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
