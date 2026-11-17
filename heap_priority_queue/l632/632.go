package l632

import "container/heap"

type item struct{ value, row, col int }

// smallestRange advances the current minimum while tracking the maximum.
// O(N log k) time, O(k) space for N total values and k lists.
func smallestRange(nums [][]int) []int {
	q := NewQueue(func(a, b item) bool { return a.value < b.value })
	high := nums[0][0]
	for r, row := range nums {
		q.Push(item{row[0], r, 0})
		high = max(high, row[0])
	}
	lo, hi := q.Peek().value, high
	for {
		x := q.Pop()
		width, best := int64(high)-int64(x.value), int64(hi)-int64(lo)
		if width < best || width == best && x.value < lo {
			lo, hi = x.value, high
		}
		if x.col+1 == len(nums[x.row]) {
			break
		}
		v := nums[x.row][x.col+1]
		high = max(high, v)
		q.Push(item{v, x.row, x.col + 1})
	}
	return []int{lo, hi}
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
