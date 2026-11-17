package l1514

import "container/heap"

type edge struct {
	to          int
	probability float64
}
type state struct {
	node        int
	probability float64
}

// maxProbability uses a maximum heap; probabilities in [0,1] cannot improve along a path.
// O(n+E log(E+1)) time, O(n+E) space.
func maxProbability(n int, edges [][]int, succProb []float64, start_node, end_node int) float64 {
	graph := make([][]edge, n)
	for i, e := range edges {
		graph[e[0]] = append(graph[e[0]], edge{e[1], succProb[i]})
		graph[e[1]] = append(graph[e[1]], edge{e[0], succProb[i]})
	}
	best := make([]float64, n)
	best[start_node] = 1
	q := NewQueue(func(a, b state) bool { return a.probability > b.probability }, state{start_node, 1})
	for q.Len() > 0 {
		s := q.Pop()
		if s.probability < best[s.node] {
			continue
		}
		if s.node == end_node {
			return s.probability
		}
		for _, e := range graph[s.node] {
			p := s.probability * e.probability
			if p > best[e.to] {
				best[e.to] = p
				q.Push(state{e.to, p})
			}
		}
	}
	return 0
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
