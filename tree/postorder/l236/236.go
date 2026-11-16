package l236

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

type S *TreeNode

// https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-tree/
// The problem guarantees that p and q are distinct nodes present in the tree.
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	state := PostOrder(root, nil, func(node *TreeNode, left S, right S) S {
		if node == p || node == q {
			return node
		}
		if left != nil && right != nil {
			return node
		}
		if left != nil {
			return left
		}
		return right
	})
	return state
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
