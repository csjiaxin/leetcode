package l433

// https://leetcode.com/problems/minimum-genetic-mutation/
func minMutation(startGene string, endGene string, bank []string) int {
	if startGene == endGene {
		return 0
	}
	dictionary := make(map[string]bool, len(bank))
	for _, gene := range bank {
		dictionary[gene] = true
	}
	if !dictionary[endGene] {
		return -1
	}
	answer := -1
	bfs([]string{startGene}, func(gene string) []string {
		next := []string{}
		letters := []byte(gene)
		for i, original := range letters {
			for _, replacement := range []byte("ACGT") {
				if replacement == original {
					continue
				}
				letters[i] = replacement
				candidate := string(letters)
				if dictionary[candidate] {
					next = append(next, candidate)
				}
			}
			letters[i] = original
		}
		return next
	}, func(gene string, distance int) bool {
		if gene == endGene {
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
