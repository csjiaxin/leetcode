package l979

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

// https://leetcode.com/problems/distribute-coins-in-binary-tree/
func distributeCoins(root *TreeNode) int {
	type S struct {
		balance, moves int
	}
	state := PostOrder(root, S{}, func(node *TreeNode, left S, right S) S {
		s := S{
			balance: node.Val + left.balance + right.balance - 1,
			moves:   abs(left.balance) + left.moves + right.moves + abs(right.balance),
		}
		return s
	})
	return state.moves
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
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
