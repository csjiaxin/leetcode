package l876

type ListNode struct {
	Val  int
	Next *ListNode
}

// Middle 在偶数长度链表上返回第二个中点。
func Middle(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

// https://leetcode.com/problems/middle-of-the-linked-list/
// 快指针每次两步，慢指针每次一步。
func middleNode(head *ListNode) *ListNode { return Middle(head) }
