package l1631

import "container/heap"

type cell struct{ row, col, effort int }

// minimumEffortPath minimizes the maximum edge difference on a path.
// O(mn log(mn)) time, O(mn) space using Dijkstra.
func minimumEffortPath(heights [][]int) int {
	m, n := len(heights), len(heights[0])
	inf := int(^uint(0) >> 1)
	dist := make([][]int, m)
	for r := range dist {
		dist[r] = make([]int, n)
		for c := range dist[r] {
			dist[r][c] = inf
		}
	}
	dist[0][0] = 0
	q := NewQueue(func(a, b cell) bool { return a.effort < b.effort }, cell{0, 0, 0})
	dirs := [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	for q.Len() > 0 {
		s := q.Pop()
		if s.effort != dist[s.row][s.col] {
			continue
		}
		if s.row == m-1 && s.col == n-1 {
			return s.effort
		}
		for _, d := range dirs {
			r, c := s.row+d[0], s.col+d[1]
			if r < 0 || r >= m || c < 0 || c >= n {
				continue
			}
			difference := heights[r][c] - heights[s.row][s.col]
			if difference < 0 {
				difference = -difference
			}
			effort := max(s.effort, difference)
			if effort < dist[r][c] {
				dist[r][c] = effort
				q.Push(cell{r, c, effort})
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
