package l23

import "leetcode/heap_priority_queue/internal/pqueue"

type ListNode struct {
	Val  int
	Next *ListNode
}

// mergeKLists reuses nodes; the heap holds one frontier per nonempty list.
// O(N log k) time and O(k) extra space for N total nodes.
func mergeKLists(lists []*ListNode) *ListNode {
	q := pqueue.New(func(a, b *ListNode) bool { return a.Val < b.Val })
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
