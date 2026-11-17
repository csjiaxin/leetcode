package l1345

// https://leetcode.com/problems/jump-game-iv/
func minJumps(arr []int) int {
	groups := map[int][]int{}
	for index, value := range arr {
		groups[value] = append(groups[value], index)
	}
	answer := -1
	bfs([]int{0}, func(index int) []int {
		next := make([]int, 0, len(groups[arr[index]])+2)
		if index > 0 {
			next = append(next, index-1)
		}
		if index+1 < len(arr) {
			next = append(next, index+1)
		}
		next = append(next, groups[arr[index]]...)
		// Its first expansion already offers every member at the earliest distance.
		delete(groups, arr[index])
		return next
	}, func(index, distance int) bool {
		if index == len(arr)-1 {
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
