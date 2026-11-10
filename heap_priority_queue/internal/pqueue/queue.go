// Package pqueue provides a typed priority queue backed by container/heap.
package pqueue

import "container/heap"

type data[T any] struct {
	items []T
	less func(T, T) bool
}

func (d data[T]) Len() int { return len(d.items) }
func (d data[T]) Less(i, j int) bool { return d.less(d.items[i], d.items[j]) }
func (d data[T]) Swap(i, j int) { d.items[i], d.items[j] = d.items[j], d.items[i] }
func (d *data[T]) Push(x any) { d.items = append(d.items, x.(T)) }
func (d *data[T]) Pop() any {
	n := len(d.items)-1
	x := d.items[n]
	var zero T
	d.items[n] = zero
	d.items = d.items[:n]
	return x
}

// Queue must not be copied after use. less(a,b) means a has higher priority.
// Peek and Pop require a nonempty queue.
type Queue[T any] struct { d data[T] }

// New copies the initial values and builds the heap in O(n).
func New[T any](less func(T,T) bool, values ...T) *Queue[T] {
	q := &Queue[T]{d: data[T]{items: append([]T(nil), values...), less: less}}
	heap.Init(&q.d)
	return q
}
func (q *Queue[T]) Len() int { return q.d.Len() }
func (q *Queue[T]) Peek() T { return q.d.items[0] }
func (q *Queue[T]) Push(x T) { heap.Push(&q.d, x) }
func (q *Queue[T]) Pop() T { return heap.Pop(&q.d).(T) }
