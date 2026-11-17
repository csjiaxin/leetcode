package l480

import "container/heap"

type entry struct{ value, index int }

// medianSlidingWindow uses two heaps, index-based lazy expiration, and logical sizes.
// Stale interior entries may remain: O(n log n) worst-case time, O(n) space.
func medianSlidingWindow(nums []int, k int) []float64 {
	out := make([]float64, 0)
	if k <= 0 || k > len(nums) {
		return out
	}
	lower := NewQueue(func(a, b entry) bool {
		if a.value != b.value {
			return a.value > b.value
		}
		return a.index > b.index
	})
	upper := NewQueue(func(a, b entry) bool {
		if a.value != b.value {
			return a.value < b.value
		}
		return a.index < b.index
	})
	inLower := make([]bool, len(nums))
	lowSize, upSize, left := 0, 0, 0
	prune := func(q *Queue[entry]) {
		for q.Len() > 0 && q.Peek().index < left {
			q.Pop()
		}
	}
	balance := func() {
		prune(lower)
		prune(upper)
		for lowSize > upSize+1 {
			x := lower.Pop()
			lowSize--
			upSize++
			inLower[x.index] = false
			upper.Push(x)
			prune(lower)
		}
		for lowSize < upSize {
			x := upper.Pop()
			upSize--
			lowSize++
			inLower[x.index] = true
			lower.Push(x)
			prune(upper)
		}
	}
	for i, v := range nums {
		if i >= k {
			old := i - k
			if inLower[old] {
				lowSize--
			} else {
				upSize--
			}
			left = old + 1
			balance()
		}
		x := entry{v, i}
		if lowSize == 0 || v <= lower.Peek().value {
			lower.Push(x)
			inLower[i] = true
			lowSize++
		} else {
			upper.Push(x)
			upSize++
		}
		balance()
		if i >= k-1 {
			median := float64(lower.Peek().value)
			if k%2 == 0 {
				median = median/2 + float64(upper.Peek().value)/2
			}
			out = append(out, median)
		}
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
