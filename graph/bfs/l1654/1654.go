package l1654

// https://leetcode.com/problems/minimum-jumps-to-reach-home/
func minimumJumps(forbidden []int, a int, b int, x int) int {
	blocked := make(map[int]bool, len(forbidden))
	maxForbidden := 0
	for _, position := range forbidden {
		blocked[position] = true
		if position > maxForbidden {
			maxForbidden = position
		}
	}
	// Beyond this bound, a detour can be shortened without crossing a forbidden point.
	upper := maxForbidden + a
	if x > upper {
		upper = x
	}
	upper += b
	if blocked[0] {
		return -1
	}
	answer := -1
	// The second field is 1 exactly when the last jump was backward.
	bfs([][2]int{{0, 0}}, func(state [2]int) [][2]int {
		next := make([][2]int, 0, 2)
		if forward := state[0] + a; forward <= upper && !blocked[forward] {
			next = append(next, [2]int{forward, 0})
		}
		if backward := state[0] - b; state[1] == 0 && backward >= 0 && !blocked[backward] {
			next = append(next, [2]int{backward, 1})
		}
		return next
	}, func(state [2]int, distance int) bool {
		if state[0] == x {
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
