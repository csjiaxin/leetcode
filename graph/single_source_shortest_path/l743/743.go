package l743

import "container/heap"

// https://leetcode.com/problems/network-delay-time/
func networkDelayTime(times [][]int, n int, k int) int {
	type edge struct {
		to     int
		weight int64
	}
	graph := make([][]edge, n)
	for _, e := range times {
		graph[e[0]-1] = append(graph[e[0]-1], edge{e[1] - 1, int64(e[2])})
	}
	const limit = int64(1<<63 - 1)
	best := BestLabels(k-1, int64(0), func(u int, distance int64, emit func(int, int64)) {
		for _, e := range graph[u] {
			if e.weight <= limit-distance {
				emit(e.to, distance+e.weight)
			}
		}
	}, func(a, b int64) bool { return a < b })
	if len(best) != n {
		return -1
	}
	var answer int64
	for _, distance := range best {
		answer = max(answer, distance)
	}
	return int(answer)
}

// BestLabels 要求扩展标签不会优于当前标签；better 同时定义优先级和严格改善。
func BestLabels[S comparable, V comparable](start S, initial V, next func(S, V, func(S, V)), better func(V, V) bool) map[S]V {
	best := map[S]V{start: initial}
	queue := &labelQueue[S, V]{better: better}
	heap.Push(queue, pathLabel[S, V]{start, initial})
	for queue.Len() > 0 {
		current := heap.Pop(queue).(pathLabel[S, V])
		if best[current.node] != current.value {
			continue
		}
		next(current.node, current.value, func(node S, value V) {
			previous, seen := best[node]
			if !seen || better(value, previous) {
				best[node] = value
				heap.Push(queue, pathLabel[S, V]{node, value})
			}
		})
	}
	return best
}

type pathLabel[S comparable, V comparable] struct {
	node  S
	value V
}
type labelQueue[S comparable, V comparable] struct {
	items  []pathLabel[S, V]
	better func(V, V) bool
}

func (q labelQueue[S, V]) Len() int           { return len(q.items) }
func (q labelQueue[S, V]) Less(i, j int) bool { return q.better(q.items[i].value, q.items[j].value) }
func (q labelQueue[S, V]) Swap(i, j int)      { q.items[i], q.items[j] = q.items[j], q.items[i] }
func (q *labelQueue[S, V]) Push(value any)    { q.items = append(q.items, value.(pathLabel[S, V])) }
func (q *labelQueue[S, V]) Pop() any {
	last := len(q.items) - 1
	value := q.items[last]
	q.items = q.items[:last]
	return value
}
