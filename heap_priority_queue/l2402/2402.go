package l2402

import (
	"container/heap"
	"sort"
)

type room struct {
	end int64
	id  int
}

// mostBooked uses an ID heap for idle rooms and an end-time/ID heap for busy rooms.
// O(m log m + (m+n) log n) time, O(m+n) space. End times may exceed int32.
func mostBooked(n int, meetings [][]int) int {
	ordered := append([][]int(nil), meetings...)
	sort.Slice(ordered, func(i, j int) bool { return ordered[i][0] < ordered[j][0] })
	idle := NewQueue(func(a, b int) bool { return a < b })
	busy := NewQueue(func(a, b room) bool {
		if a.end != b.end {
			return a.end < b.end
		}
		return a.id < b.id
	})
	for i := 0; i < n; i++ {
		idle.Push(i)
	}
	counts := make([]int, n)
	for _, m := range ordered {
		start, duration := int64(m[0]), int64(m[1])-int64(m[0])
		for busy.Len() > 0 && busy.Peek().end <= start {
			idle.Push(busy.Pop().id)
		}
		var r room
		if idle.Len() > 0 {
			r = room{start, idle.Pop()}
		} else {
			r = busy.Pop()
		}
		r.end += duration
		counts[r.id]++
		busy.Push(r)
	}
	best := 0
	for i := 1; i < n; i++ {
		if counts[i] > counts[best] {
			best = i
		}
	}
	return best
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
