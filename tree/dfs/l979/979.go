package l979

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

// https://leetcode.com/problems/distribute-coins-in-binary-tree/
func distributeCoins(root *TreeNode) int {
	type State struct {
		balance, moves int
	}
	state := PostOrder(root, State{}, func(node *TreeNode, left, right State) State {
		return State{
			balance: node.Val - 1 + left.balance + right.balance,
			moves:   left.moves + right.moves + abs(left.balance) + abs(right.balance),
		}
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
