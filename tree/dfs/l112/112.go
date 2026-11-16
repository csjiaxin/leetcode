package l112

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}
type S struct {
	sum int
}

// PathDFS 传递路径状态，并在离开节点时撤销共享状态。
func PathDFS[S any](node *TreeNode, parentState S, enter func(node *TreeNode, parentState S) (state S, undo func())) {
	if node == nil {
		return
	}
	state, undo := enter(node, parentState)
	if undo != nil {
		defer undo()
	}
	PathDFS(node.Left, state, enter)
	PathDFS(node.Right, state, enter)
}

// https://leetcode.com/problems/path-sum/
func hasPathSum(root *TreeNode, targetSum int) bool {
	found := false
	PathDFS(root, S{}, func(node *TreeNode, parentState S) (state S, undo func()) {
		s := S{
			sum: parentState.sum + node.Val,
		}
		if node.Left == nil && node.Right == nil {
			if s.sum == targetSum {
				found = true
			}
		}
		return s, nil
	})
	return found
}
