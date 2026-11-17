package l1514

import "container/heap"

// https://leetcode.com/problems/path-with-maximum-probability/
func maxProbability(n int, edges [][]int, succProb []float64, startNode int, endNode int) float64 {
	type edge struct {
		to          int
		probability float64
	}
	graph := make([][]edge, n)
	for i, e := range edges {
		graph[e[0]] = append(graph[e[0]], edge{e[1], succProb[i]})
		graph[e[1]] = append(graph[e[1]], edge{e[0], succProb[i]})
	}
	best := BestLabels(startNode, 1.0, func(u int, probability float64, emit func(int, float64)) {
		for _, e := range graph[u] {
			emit(e.to, probability*e.probability)
		}
	}, func(a, b float64) bool { return a > b })
	return best[endNode]
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
