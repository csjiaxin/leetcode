package l1311

import "sort"

// https://leetcode.com/problems/get-watched-videos-by-your-friends/
func watchedVideosByFriends(watchedVideos [][]string, friends [][]int, id int, level int) []string {
	frequency := map[string]int{}
	atLevel := false
	bfs([]int{id}, func(person int) []int {
		if atLevel {
			return nil // Every state at the requested depth is visited, never expanded.
		}
		return friends[person]
	}, func(person, distance int) bool {
		if distance == level {
			atLevel = true
			for _, title := range watchedVideos[person] {
				frequency[title]++
			}
		}
		return false
	})
	answer := make([]string, 0, len(frequency))
	for title := range frequency {
		answer = append(answer, title)
	}
	sort.Slice(answer, func(i, j int) bool {
		if frequency[answer[i]] != frequency[answer[j]] {
			return frequency[answer[i]] < frequency[answer[j]]
		}
		return answer[i] < answer[j]
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
