package l141

type ListNode struct {
	Val  int
	Next *ListNode
}

// Floyd 用确定后继函数寻找环内相遇点；next 的 bool=false 表示轨迹结束。
func Floyd[S comparable](start S, next func(S) (S, bool)) (S, bool) {
	slow, fast := start, start
	for {
		var ok bool
		fast, ok = next(fast)
		if !ok {
			var zero S
			return zero, false
		}
		fast, ok = next(fast)
		if !ok {
			var zero S
			return zero, false
		}
		slow, ok = next(slow)
		if !ok {
			var zero S
			return zero, false
		}
		if slow == fast {
			return slow, true
		}
	}
}

func NodeStep(node *ListNode) (*ListNode, bool) {
	if node == nil || node.Next == nil {
		return nil, false
	}
	return node.Next, true
}

// https://leetcode.com/problems/linked-list-cycle/
// 先推进后比较；有限链表结束时不能把两个 nil 当成相遇环。
func hasCycle(head *ListNode) bool { _, found := Floyd(head, NodeStep); return found }
