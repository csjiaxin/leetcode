package l1129

// https://leetcode.com/problems/shortest-path-with-alternating-colors/
func shortestAlternatingPaths(n int, redEdges [][]int, blueEdges [][]int) []int {
	adjacent := [2][][]int{make([][]int, n), make([][]int, n)}
	for color, edges := range [][][]int{redEdges, blueEdges} {
		for _, edge := range edges {
			adjacent[color][edge[0]] = append(adjacent[color][edge[0]], edge[1])
		}
	}
	// Seeding both last colors allows either color for the first edge.
	dist := bfs([][2]int{{0, 0}, {0, 1}}, func(state [2]int) [][2]int {
		color := 1 - state[1]
		next := make([][2]int, 0, len(adjacent[color][state[0]]))
		for _, node := range adjacent[color][state[0]] {
			next = append(next, [2]int{node, color})
		}
		return next
	}, nil)
	answer := make([]int, n)
	for node := range answer {
		answer[node] = -1
		for color := 0; color < 2; color++ {
			if distance, found := dist[[2]int{node, color}]; found && (answer[node] == -1 || distance < answer[node]) {
				answer[node] = distance
			}
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
