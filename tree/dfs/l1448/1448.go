package l1448

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
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

// https://leetcode.com/problems/count-good-nodes-in-binary-tree/
func goodNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	count := 0
	// 根值作为初始最大值，也适用于全负数树。
	PathDFS(root, root.Val, func(node *TreeNode, maximum int) (int, func()) {
		if node.Val >= maximum {
			count++
			maximum = node.Val
		}
		return maximum, nil
	})
	return count
}
