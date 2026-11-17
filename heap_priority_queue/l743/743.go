package l743

import "container/heap"

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
	q := NewQueue(func(a, b state) bool { return a.distance < b.distance }, state{k, 0})
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

// 以下泛型堆模板与题解放在同一个文件，便于直接提交。
type heapData[T any] struct {
	items []T
	less  func(T, T) bool
}

func (d heapData[T]) Len() int           { return len(d.items) }
func (d heapData[T]) Less(i, j int) bool { return d.less(d.items[i], d.items[j]) }
func (d heapData[T]) Swap(i, j int)      { d.items[i], d.items[j] = d.items[j], d.items[i] }
func (d *heapData[T]) Push(x any)        { d.items = append(d.items, x.(T)) }
func (d *heapData[T]) Pop() any {
	n := len(d.items) - 1
	x := d.items[n]
	var zero T
	d.items[n] = zero
	d.items = d.items[:n]
	return x
}

// Queue must not be copied after use. less(a,b) means a has higher priority.
// Peek and Pop require a nonempty queue.
type Queue[T any] struct{ d heapData[T] }

// NewQueue copies the initial values and builds the heap in O(n).
func NewQueue[T any](less func(T, T) bool, values ...T) *Queue[T] {
	q := &Queue[T]{d: heapData[T]{items: append([]T(nil), values...), less: less}}
	heap.Init(&q.d)
	return q
}
func (q *Queue[T]) Len() int { return q.d.Len() }
func (q *Queue[T]) Peek() T  { return q.d.items[0] }
func (q *Queue[T]) Push(x T) { heap.Push(&q.d, x) }
func (q *Queue[T]) Pop() T   { return heap.Pop(&q.d).(T) }
