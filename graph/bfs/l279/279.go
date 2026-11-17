package l279

// https://leetcode.com/problems/perfect-squares/
func numSquares(n int) int {
	squares := []int{}
	for value := 1; value*value <= n; value++ {
		squares = append(squares, value*value)
	}
	answer := -1
	bfs([]int{0}, func(sum int) []int {
		next := []int{}
		for _, square := range squares {
			if sum+square > n {
				break
			}
			next = append(next, sum+square)
		}
		return next
	}, func(sum, distance int) bool {
		if sum == n {
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
