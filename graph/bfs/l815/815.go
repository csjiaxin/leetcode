package l815

// https://leetcode.com/problems/bus-routes/
func numBusesToDestination(routes [][]int, source int, target int) int {
	if source == target {
		return 0
	}
	stopRoutes := map[int][]int{}
	for route, stops := range routes {
		for _, stop := range stops {
			stopRoutes[stop] = append(stopRoutes[stop], route)
		}
	}
	expandedRoute := make([]bool, len(routes))
	answer := -1
	// One edge connects two stops reachable by boarding the same bus.
	bfs([]int{source}, func(stop int) []int {
		next := []int{}
		for _, route := range stopRoutes[stop] {
			if expandedRoute[route] {
				continue
			}
			expandedRoute[route] = true
			next = append(next, routes[route]...)
		}
		return next
	}, func(stop, distance int) bool {
		if stop == target {
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
