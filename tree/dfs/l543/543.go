package l543

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

// https://leetcode.com/problems/diameter-of-binary-tree/
func diameterOfBinaryTree(root *TreeNode) int {
	type State struct {
		height, diameter int
	}

	state := PostOrder(root, State{}, func(node *TreeNode, left, right State) State {
		return State{
			height:   max(left.height, right.height) + 1,
			diameter: max(left.diameter, right.diameter, left.height+right.height),
		}
	})

	return state.diameter

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
