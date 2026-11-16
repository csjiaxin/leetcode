package l113

import "slices"

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

type S struct {
	sum int
}

// https://leetcode.com/problems/path-sum-ii/
func pathSum(root *TreeNode, targetSum int) [][]int {
	result := make([][]int, 0)
	path := make([]int, 0)
	PathDFS(root, S{}, func(node *TreeNode, parentState S) (state S, undo func()) {
		s := S{
			sum: parentState.sum + node.Val,
		}
		path = append(path, node.Val)
		if node.Left == nil && node.Right == nil {
			if s.sum == targetSum {
				result = append(result, slices.Clone(path))
			}
		}
		return s, func() {
			path = path[:len(path)-1]
		}
	})
	return result
}
