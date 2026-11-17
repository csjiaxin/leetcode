package l295

import "container/heap"

// MedianFinder partitions values into lower(max heap) and upper(min heap).
// Lower has either the same size as upper or one additional element.
type MedianFinder struct{ lower, upper *Queue[int] }

func Constructor() MedianFinder {
	return MedianFinder{NewQueue(func(a, b int) bool { return a > b }), NewQueue(func(a, b int) bool { return a < b })}
}

// AddNum takes O(log n) time. All inserted values occupy O(n) space.
func (m *MedianFinder) AddNum(num int) {
	if m.lower.Len() == 0 || num <= m.lower.Peek() {
		m.lower.Push(num)
	} else {
		m.upper.Push(num)
	}
	if m.lower.Len() > m.upper.Len()+1 {
		m.upper.Push(m.lower.Pop())
	}
	if m.upper.Len() > m.lower.Len() {
		m.lower.Push(m.upper.Pop())
	}
}

// FindMedian takes O(1) time; at least one number must have been inserted.
func (m *MedianFinder) FindMedian() float64 {
	if m.lower.Len() > m.upper.Len() {
		return float64(m.lower.Peek())
	}
	return float64(m.lower.Peek())/2 + float64(m.upper.Peek())/2
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
