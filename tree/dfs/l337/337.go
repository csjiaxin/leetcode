package l337

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

// https://leetcode.com/problems/house-robber-iii/
type S struct {
	skip, rob int
}

func rob(root *TreeNode) int {
	s := PostOrder(root, S{}, func(node *TreeNode, left S, right S) S {
		return S{
			skip: max(left.rob, left.skip) + max(right.rob, right.skip),
			rob:  node.Val + left.skip + right.skip,
		}
	})
	return max(s.rob, s.skip)
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
