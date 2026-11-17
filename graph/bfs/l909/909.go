package l909

// https://leetcode.com/problems/snakes-and-ladders/
func snakesAndLadders(board [][]int) int {
	n := len(board)
	answer := -1
	bfs([]int{1}, func(square int) []int {
		next := make([]int, 0, 6)
		for roll := 1; roll <= 6 && square+roll <= n*n; roll++ {
			landing := square + roll
			rowFromBottom := (landing - 1) / n
			row, col := n-1-rowFromBottom, (landing-1)%n
			if rowFromBottom%2 == 1 {
				col = n - 1 - col
			}
			if board[row][col] != -1 {
				landing = board[row][col] // Exactly one snake/ladder per roll.
			}
			next = append(next, landing)
		}
		return next
	}, func(square, distance int) bool {
		if square == n*n {
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
