package l630

import (
	"container/heap"
	"slices"
)

// scheduleCourse sorts deadlines; when over budget, discard the longest course.
// O(n log n) time, O(n) space. The input is not reordered.
func scheduleCourse(courses [][]int) int {
	slices.SortFunc(courses, func(a []int, b []int) int { return a[1] - b[1] })
	total := 0
	q := NewQueue(func(a, b []int) bool { return a[0] > b[0] })

	drop := 0
	for i := 0; i < len(courses); i++ {
		c := courses[i]
		q.Push(c)
		total += c[0]
		if total > c[1] {
			removed := q.Pop()
			total -= removed[0]
			drop++
		}
	}
	return len(courses) - drop
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
