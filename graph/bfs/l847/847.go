package l847

// https://leetcode.com/problems/shortest-path-visiting-all-nodes/
func shortestPathLength(graph [][]int) int {
	n := len(graph)
	starts := make([][2]int, n)
	for node := range graph {
		starts[node] = [2]int{node, 1 << node}
	}
	all := (1 << n) - 1
	answer := -1
	bfs(starts, func(state [2]int) [][2]int {
		next := make([][2]int, 0, len(graph[state[0]]))
		for _, node := range graph[state[0]] {
			next = append(next, [2]int{node, state[1] | (1 << node)})
		}
		return next
	}, func(state [2]int, distance int) bool {
		if state[1] == all {
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
