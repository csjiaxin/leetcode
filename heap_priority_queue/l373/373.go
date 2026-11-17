package l373

import "container/heap"

type pair struct {
	sum  int64
	i, j int
}

// kSmallestPairs merges rows nums1[i]+nums2[j] without generating the product.
// O(k log min(k,len(nums1))) time, O(min(k,len(nums1))) heap space.
func kSmallestPairs(nums1, nums2 []int, k int) [][]int {
	out := make([][]int, 0)
	if len(nums1) == 0 || len(nums2) == 0 || k <= 0 {
		return out
	}
	q := NewQueue(func(a, b pair) bool {
		if a.sum != b.sum {
			return a.sum < b.sum
		}
		if a.i != b.i {
			return a.i < b.i
		}
		return a.j < b.j
	})
	for i := 0; i < len(nums1) && i < k; i++ {
		q.Push(pair{int64(nums1[i]) + int64(nums2[0]), i, 0})
	}
	for len(out) < k && q.Len() > 0 {
		x := q.Pop()
		out = append(out, []int{nums1[x.i], nums2[x.j]})
		if x.j+1 < len(nums2) {
			q.Push(pair{int64(nums1[x.i]) + int64(nums2[x.j+1]), x.i, x.j + 1})
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
