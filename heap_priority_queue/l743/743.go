package l743

import "leetcode/heap_priority_queue/internal/pqueue"

type edge struct{ to, weight int }
type state struct{ node, distance int }

// networkDelayTime uses Dijkstra with lazy outdated-state removal.
// O(n+E log(E+1)) time, O(n+E) space; all edge weights are nonnegative.
func networkDelayTime(times [][]int, n, k int) int {
	graph := make([][]edge, n+1)
	for _, e := range times {
		graph[e[0]] = append(graph[e[0]], edge{e[1], e[2]})
	}
	inf := int(^uint(0) >> 1)
	dist := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dist[i] = inf
	}
	dist[k] = 0
	q := pqueue.New(func(a, b state) bool { return a.distance < b.distance }, state{k, 0})
	for q.Len() > 0 {
		s := q.Pop()
		if s.distance != dist[s.node] {
			continue
		}
		for _, e := range graph[s.node] {
			d := s.distance + e.weight
			if d < dist[e.to] {
				dist[e.to] = d
				q.Push(state{e.to, d})
			}
		}
	}
	answer := 0
	for i := 1; i <= n; i++ {
		if dist[i] == inf {
			return -1
		}
		answer = max(answer, dist[i])
	}
	return answer
}
