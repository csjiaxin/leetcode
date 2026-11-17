package l301

// https://leetcode.com/problems/remove-invalid-parentheses/
func removeInvalidParentheses(s string) []string {
	valid := func(text string) bool {
		balance := 0
		for _, character := range text {
			if character == '(' {
				balance++
			} else if character == ')' {
				balance--
				if balance < 0 {
					return false
				}
			}
		}
		return balance == 0
	}
	answer := []string{}
	foundDepth := -1
	bfs([]string{s}, func(text string) []string {
		if foundDepth >= 0 {
			return nil // Keep checking this depth, but generate no deeper states.
		}
		next := []string{}
		for i := 0; i < len(text); i++ {
			if text[i] != '(' && text[i] != ')' {
				continue
			}
			if i > 0 && text[i] == text[i-1] {
				continue // Removing either identical adjacent bracket gives the same string.
			}
			next = append(next, text[:i]+text[i+1:])
		}
		return next
	}, func(text string, distance int) bool {
		if foundDepth >= 0 && distance > foundDepth {
			return true // Do not even validate the already queued deeper states.
		}
		if valid(text) {
			foundDepth = distance
			answer = append(answer, text)
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
