package l127

// https://leetcode.com/problems/word-ladder/
func ladderLength(beginWord string, endWord string, wordList []string) int {
	if beginWord == endWord {
		return 1
	}
	dictionary := make(map[string]bool, len(wordList))
	for _, word := range wordList {
		dictionary[word] = true
	}
	if !dictionary[endWord] {
		return 0
	}
	answer := 0
	bfs([]string{beginWord}, func(word string) []string {
		next := []string{}
		letters := []byte(word)
		for i, original := range letters {
			for replacement := byte('a'); replacement <= 'z'; replacement++ {
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
	}, func(word string, distance int) bool {
		if word == endWord {
			answer = distance + 1 // Count words, rather than transformations.
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
