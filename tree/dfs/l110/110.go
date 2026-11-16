package l110

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

type S struct {
	height   int
	balanced bool
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

// https://leetcode.com/problems/balanced-binary-tree/
func isBalanced(root *TreeNode) bool {
	state := PostOrder(root, S{
		balanced: true,
	}, func(node *TreeNode, left S, right S) S {
		balanced := abs(left.height-right.height) <= 1
		return S{
			height:   max(left.height, right.height) + 1,
			balanced: left.balanced && right.balanced && balanced,
		}
	})
	return state.balanced
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
