package l773

// https://leetcode.com/problems/sliding-puzzle/
func slidingPuzzle(board [][]int) int {
	encoded := make([]byte, 0, 6)
	for _, row := range board {
		for _, tile := range row {
			encoded = append(encoded, byte(tile)+'0')
		}
	}
	adjacent := [][]int{{1, 3}, {0, 2, 4}, {1, 5}, {0, 4}, {1, 3, 5}, {2, 4}}
	answer := -1
	bfs([]string{string(encoded)}, func(state string) []string {
		tiles := []byte(state)
		zero := 0
		for tiles[zero] != '0' {
			zero++
		}
		next := make([]string, 0, len(adjacent[zero]))
		for _, neighbor := range adjacent[zero] {
			tiles[zero], tiles[neighbor] = tiles[neighbor], tiles[zero]
			next = append(next, string(tiles))
			tiles[zero], tiles[neighbor] = tiles[neighbor], tiles[zero]
		}
		return next
	}, func(state string, distance int) bool {
		if state == "123450" {
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
