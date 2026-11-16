package l968

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

// https://leetcode.com/problems/binary-tree-cameras/
func minCameraCover(root *TreeNode) int {
	// Larger than any legal answer (at most 1000 nodes), with safe addition.
	const inf = 1 << 20
	type State struct {
		camera    int // Current node has a camera; entire subtree is covered.
		covered   int // Covered by a child, without a camera at the current node.
		uncovered int // Only the current node is uncovered, waiting for its parent.
	}
	empty := State{camera: inf}
	state := PostOrder(root, empty, func(node *TreeNode, left, right State) State {
		return State{
			camera: min(inf, 1+min(left.camera, left.covered, left.uncovered)+
				min(right.camera, right.covered, right.uncovered)),
			covered: min(inf, left.camera+min(right.camera, right.covered),
				right.camera+min(left.camera, left.covered)),
			uncovered: min(inf, left.covered+right.covered),
		}
	})
	return min(state.camera, state.covered)
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
