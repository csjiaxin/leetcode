package l23

import "container/heap"

type ListNode struct {
	Val  int
	Next *ListNode
}

// mergeKLists reuses nodes; the heap holds one frontier per nonempty list.
// O(N log k) time and O(k) extra space for N total nodes.
func mergeKLists(lists []*ListNode) *ListNode {
	q := NewQueue(func(a, b *ListNode) bool { return a.Val < b.Val })
	for _, node := range lists {
		if node != nil {
			q.Push(node)
		}
	}
	dummy := &ListNode{}
	tail := dummy
	for q.Len() > 0 {
		node := q.Pop()
		if node.Next != nil {
			q.Push(node.Next)
		}
		tail.Next = node
		tail = node
	}
	if tail != dummy {
		tail.Next = nil
	}
	return dummy.Next
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
