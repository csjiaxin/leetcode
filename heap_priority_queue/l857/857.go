package l857

import (
	"container/heap"
	"math"
	"sort"
)

type worker struct{ quality, wage int }

// mincostToHireWorkers scans increasing minimum wage/quality ratios.
// At each ratio a max heap retains k smallest qualities. O(n log n) time, O(n) space.
func mincostToHireWorkers(quality, wage []int, k int) float64 {
	workers := make([]worker, len(quality))
	for i := range quality {
		workers[i] = worker{quality[i], wage[i]}
	}
	sort.Slice(workers, func(i, j int) bool {
		return int64(workers[i].wage)*int64(workers[j].quality) < int64(workers[j].wage)*int64(workers[i].quality)
	})
	q := NewQueue(func(a, b int) bool { return a > b })
	var sum int64
	best := math.Inf(1)
	for _, w := range workers {
		q.Push(w.quality)
		sum += int64(w.quality)
		if q.Len() > k {
			sum -= int64(q.Pop())
		}
		if q.Len() == k {
			best = math.Min(best, float64(sum)*float64(w.wage)/float64(w.quality))
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
