package l752

// https://leetcode.com/problems/open-the-lock/
func openLock(deadends []string, target string) int {
	dead := make(map[string]bool, len(deadends))
	for _, state := range deadends {
		dead[state] = true
	}
	if dead["0000"] {
		return -1
	}
	answer := -1
	bfs([]string{"0000"}, func(state string) []string {
		next := make([]string, 0, 8)
		digits := []byte(state)
		for i, original := range digits {
			for _, change := range []int{1, 9} {
				digits[i] = byte((int(original-'0')+change)%10) + '0'
				candidate := string(digits)
				if !dead[candidate] {
					next = append(next, candidate)
				}
			}
			digits[i] = original
		}
		return next
	}, func(state string, distance int) bool {
		if state == target {
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
