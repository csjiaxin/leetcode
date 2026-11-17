package l234

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

func ReverseList(head *ListNode) *ListNode {
	var previous *ListNode
	for head != nil {
		next := head.Next
		head.Next = previous
		previous = head
		head = next
	}
	return previous
}

// https://leetcode.com/problems/palindrome-linked-list/
// 反转后半段进行比较，返回前恢复链表；额外空间 O(1)。
func isPalindrome(head *ListNode) bool {
	second := ReverseList(Middle(head))
	left, right := head, second
	answer := true
	for right != nil {
		if left.Val != right.Val {
			answer = false
			break
		}
		left = left.Next
		right = right.Next
	}
	ReverseList(second)
	return answer
}
