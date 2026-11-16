package l124

import "math"

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

type S struct {
	maxpath, maxCanConnect int
}

// https://leetcode.com/problems/binary-tree-maximum-path-sum/
// The problem guarantees a nonempty tree.
func maxPathSum(root *TreeNode) int {

	state := PostOrder(root, S{
		maxpath: math.MinInt,
	}, func(node *TreeNode, left S, right S) S {
		maxpath := node.Val
		if left.maxCanConnect > 0 {
			maxpath += left.maxCanConnect
		}
		if right.maxCanConnect > 0 {
			maxpath += right.maxCanConnect
		}

		return S{
			maxpath:       max(left.maxpath, right.maxpath, maxpath),
			maxCanConnect: max(left.maxCanConnect, right.maxCanConnect, 0) + node.Val,
		}
	})
	return state.maxpath
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
