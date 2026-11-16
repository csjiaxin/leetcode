package l543

import "testing"

func TestDiameterOfBinaryTree(t *testing.T) {
	for _, tt := range []struct {
		name string
		root *TreeNode
		want int
	}{
		{"leaf", &TreeNode{}, 0}, {"one edge", &TreeNode{Left: &TreeNode{}}, 1},
		{"example", &TreeNode{Left: &TreeNode{Left: &TreeNode{}, Right: &TreeNode{}}, Right: &TreeNode{}}, 3},
		{"chain", &TreeNode{Right: &TreeNode{Right: &TreeNode{Right: &TreeNode{}}}}, 3},
		{"below root", &TreeNode{Left: &TreeNode{Left: &TreeNode{Left: &TreeNode{}}, Right: &TreeNode{Right: &TreeNode{}}}}, 4},
	} {
		t.Run(tt.name, func(t *testing.T) {
			if got := diameterOfBinaryTree(tt.root); got != tt.want {
				t.Errorf("got %d want %d", got, tt.want)
			}
		})
	}
}
