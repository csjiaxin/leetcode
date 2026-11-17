package l1834

import (
	"container/heap"
	"sort"
)

type task struct{ arrival, duration, index int }

// getOrder selects available tasks by processing time, then original index.
// O(n log n) time, O(n) space. The clock uses int64 for cumulative durations.
func getOrder(tasks [][]int) []int {
	ordered := make([]task, len(tasks))
	for i, t := range tasks {
		ordered[i] = task{t[0], t[1], i}
	}
	sort.Slice(ordered, func(i, j int) bool { return ordered[i].arrival < ordered[j].arrival })
	q := NewQueue(func(a, b task) bool {
		if a.duration != b.duration {
			return a.duration < b.duration
		}
		return a.index < b.index
	})
	out := make([]int, 0, len(tasks))
	i := 0
	var now int64
	for i < len(ordered) || q.Len() > 0 {
		if q.Len() == 0 {
			now = max(now, int64(ordered[i].arrival))
		}
		for i < len(ordered) && int64(ordered[i].arrival) <= now {
			q.Push(ordered[i])
			i++
		}
		t := q.Pop()
		now += int64(t.duration)
		out = append(out, t.index)
	}
	return out
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
