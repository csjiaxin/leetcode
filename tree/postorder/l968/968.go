package l968

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

// https://leetcode.com/problems/binary-tree-cameras/
func minCameraCover(root *TreeNode) int {
	type S struct {
		canExtend    bool
		installLight bool
		cameras      int
	}

	s := PostOrder(root, S{
		canExtend: false,
	}, func(node *TreeNode, left S, right S) S {
		canExtend := false
		installLight := false
		cameras := left.cameras + right.cameras
		if left.canExtend && right.canExtend {
			cameras -= 1
			canExtend = false
			installLight = true
		} else if left.canExtend || right.canExtend {
			canExtend = false
			installLight = true
		} else {
			if !(left.installLight || right.installLight) {
				cameras += 1
				canExtend = true
				installLight = true
			}
		}
		return S{
			canExtend:    canExtend,
			cameras:      cameras,
			installLight: installLight,
		}
	})
	return s.cameras
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
