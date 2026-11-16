package l437

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
	sum []int
}

// https://leetcode.com/problems/path-sum-iii/
func pathSum(root *TreeNode, targetSum int) int {
	count := 0

	PathDFS(root, S{}, func(node *TreeNode, parentState S) (state S, undo func()) {
		sum := node.Val
		if len(parentState.sum) > 0 {
			sum += parentState.sum[len(parentState.sum)-1]
			for i := 0; i < len(parentState.sum); i++ {
				if sum-parentState.sum[i] == targetSum {
					count++
				}
			}
		}
		if sum == targetSum {
			count++
		}

		s := S{}
		s.sum = append(s.sum, parentState.sum...)
		s.sum = append(s.sum, sum)
		return s, nil
	})
	return count
}
