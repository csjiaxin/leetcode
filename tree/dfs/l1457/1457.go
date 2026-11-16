package l1457

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

// https://leetcode.com/problems/pseudo-palindromic-paths-in-a-binary-tree/
func pseudoPalindromicPaths(root *TreeNode) int {
	count := 0
	PathDFS(root, 0, func(node *TreeNode, mask int) (int, func()) {
		// 翻转当前数字的奇偶位；叶子处至多一位为 1。
		mask ^= 1 << uint(node.Val)
		if node.Left == nil && node.Right == nil && mask&(mask-1) == 0 {
			count++
		}
		return mask, nil
	})
	return count
}
