package l1631

import "container/heap"

// https://leetcode.com/problems/path-with-minimum-effort/
func minimumEffortPath(heights [][]int) int {
	if len(heights) == 0 || len(heights[0]) == 0 {
		return 0
	}
	rows, cols := len(heights), len(heights[0])
	directions := [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	best := BestLabels([2]int{0, 0}, int64(0), func(at [2]int, effort int64, emit func([2]int, int64)) {
		for _, direction := range directions {
			r, c := at[0]+direction[0], at[1]+direction[1]
			if r < 0 || r >= rows || c < 0 || c >= cols {
				continue
			}
			delta := int64(heights[r][c]) - int64(heights[at[0]][at[1]])
			if delta < 0 {
				delta = -delta
			}
			emit([2]int{r, c}, max(effort, delta))
		}
	}, func(a, b int64) bool { return a < b })
	return int(best[[2]int{rows - 1, cols - 1}])
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
