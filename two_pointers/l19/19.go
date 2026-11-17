package l19

type ListNode struct {
	Val  int
	Next *ListNode
}

// NthPredecessor 从哑节点出发，使 runner 领先 n 步，再同步走到尾端。
func NthPredecessor(dummy *ListNode, n int) *ListNode {
	if n <= 0 {
		return nil
	}
	runner := dummy
	for step := 0; step < n; step++ {
		if runner.Next == nil {
			return nil
		}
		runner = runner.Next
	}
	previous := dummy
	for runner.Next != nil {
		runner = runner.Next
		previous = previous.Next
	}
	return previous
}

// https://leetcode.com/problems/remove-nth-node-from-end-of-list/
// 哑节点统一删除头部的情况；慢指针最后停在待删节点之前。
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}
	previous := NthPredecessor(dummy, n)
	if previous != nil {
		previous.Next = previous.Next.Next
	}
	return dummy.Next
}
