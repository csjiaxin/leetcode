package l142

type ListNode struct {
	Val  int
	Next *ListNode
}

// https://leetcode.com/problems/linked-list-cycle-ii/
func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
		if slow == fast {
			p := head
			for ; p != slow; p, slow = p.Next, slow.Next {
			}
			return p
		}
	}

	return nil
}
