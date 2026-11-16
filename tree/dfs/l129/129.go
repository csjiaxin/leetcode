package l129

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

// https://leetcode.com/problems/sum-root-to-leaf-numbers/
func sumNumbers(root *TreeNode) int {
	total := 0
	PathDFS(root, 0, func(node *TreeNode, value int) (int, func()) {
		value = value*10 + node.Val
		if node.Left == nil && node.Right == nil {
			total += value
		}
		return value, nil
	})
	return total
}
