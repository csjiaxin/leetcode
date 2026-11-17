package l407

import "container/heap"

type cell struct{ row, col, height int }

// trapRainWater expands the lowest boundary; visited cells are marked on insertion.
// O(mn log(mn)) time, O(mn) space.
func trapRainWater(heightMap [][]int) int {
	m := len(heightMap)
	if m == 0 {
		return 0
	}
	n := len(heightMap[0])
	if m < 3 || n < 3 {
		return 0
	}
	visited := make([][]bool, m)
	for r := range visited {
		visited[r] = make([]bool, n)
	}
	q := NewQueue(func(a, b cell) bool { return a.height < b.height })
	add := func(r, c int) {
		if !visited[r][c] {
			visited[r][c] = true
			q.Push(cell{r, c, heightMap[r][c]})
		}
	}
	for r := 0; r < m; r++ {
		add(r, 0)
		add(r, n-1)
	}
	for c := 0; c < n; c++ {
		add(0, c)
		add(m-1, c)
	}
	dirs := [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	water := 0
	for q.Len() > 0 {
		s := q.Pop()
		for _, d := range dirs {
			r, c := s.row+d[0], s.col+d[1]
			if r < 0 || r >= m || c < 0 || c >= n || visited[r][c] {
				continue
			}
			visited[r][c] = true
			water += max(0, s.height-heightMap[r][c])
			q.Push(cell{r, c, max(s.height, heightMap[r][c])})
		}
	}
	return water
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
