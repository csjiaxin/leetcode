package l621

import "container/heap"

type cooling struct{ ready, count int }

// leastInterval greedily runs the largest available frequency and queues cooldowns.
// O(T log A) time for T tasks and A distinct task kinds; O(A) space.
// Idle periods are jumped over rather than simulated one unit at a time.
func leastInterval(tasks []byte, n int) int {
	counts := make(map[byte]int)
	for _, c := range tasks {
		counts[c]++
	}
	q := NewQueue(func(a, b int) bool { return a > b })
	for _, c := range counts {
		q.Push(c)
	}
	wait := make([]cooling, 0, len(counts))
	head, now := 0, 0
	for q.Len() > 0 || head < len(wait) {
		if q.Len() == 0 {
			now = max(now, wait[head].ready)
		}
		for head < len(wait) && wait[head].ready <= now {
			q.Push(wait[head].count)
			head++
		}
		remaining := q.Pop() - 1
		now++
		if remaining > 0 {
			wait = append(wait, cooling{now + n, remaining})
		}
		// Compact consumed cooldowns to keep storage bounded by active kinds.
		if head > 0 && head*2 >= len(wait) {
			wait = append(wait[:0], wait[head:]...)
			head = 0
		}
	}
	return now
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
