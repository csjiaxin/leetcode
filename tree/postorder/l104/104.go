package l104

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

// https://leetcode.com/problems/maximum-depth-of-binary-tree/
func maxDepth(root *TreeNode) int {
	return PostOrder(root, 0, func(node *TreeNode, left, right int) int {
		return 1 + max(left, right)
	})
}

func PostOrder[S any](
	node *TreeNode,
	empty S,
	merge func(node *TreeNode, left, right S) S,
) S {
	if node == nil {
		return empty
	}

	left := PostOrder(node.Left, empty, merge)
	right := PostOrder(node.Right, empty, merge)

	return merge(node, left, right)
}
