package l1882

import "container/heap"

type server struct {
	end        int64
	weight, id int
}

// assignTasks maintains idle weight/ID priority and busy end/weight/ID priority.
// O((n+m) log n) time, O(n+m) space including output; cumulative times use int64.
func assignTasks(servers, tasks []int) []int {
	byWeight := func(a, b server) bool {
		if a.weight != b.weight {
			return a.weight < b.weight
		}
		return a.id < b.id
	}
	idle := NewQueue(byWeight)
	busy := NewQueue(func(a, b server) bool {
		if a.end != b.end {
			return a.end < b.end
		}
		return byWeight(a, b)
	})
	for i, w := range servers {
		idle.Push(server{weight: w, id: i})
	}
	out := make([]int, len(tasks))
	var now int64
	for i, duration := range tasks {
		now = max(now, int64(i))
		if idle.Len() == 0 {
			now = max(now, busy.Peek().end)
		}
		// Release every simultaneous completion before choosing by weight.
		for busy.Len() > 0 && busy.Peek().end <= now {
			idle.Push(busy.Pop())
		}
		s := idle.Pop()
		out[i] = s.id
		s.end = now + int64(duration)
		busy.Push(s)
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
