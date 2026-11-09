package l337

// TreeNode is a node in a binary tree.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// rob returns the maximum amount that can be robbed without robbing both
// a house and its directly connected child.
func rob(root *TreeNode) int {

	var dfs func(node *TreeNode) (skip, take int)
	dfs = func(node *TreeNode) (skip, take int) {
		if node == nil {
			return 0, 0
		}
		lskip, ltake := dfs(node.Left)
		rskip, rtake := dfs(node.Right)

		return max(lskip, ltake) + max(rskip, rtake), node.Val + lskip + rskip
	}
	skip, take := dfs(root)
	return max(skip, take)
}
