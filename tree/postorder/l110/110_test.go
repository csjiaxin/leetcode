package l110

import "testing"

func TestIsBalanced(t *testing.T) {
	for _, tt := range []struct {
		name string
		root *TreeNode
		want bool
	}{
		{"empty", nil, true}, {"leaf", &TreeNode{}, true},
		{"one child", &TreeNode{Left: &TreeNode{}}, true},
		{"balanced", &TreeNode{Left: &TreeNode{}, Right: &TreeNode{Left: &TreeNode{}}}, true},
		{"chain", &TreeNode{Left: &TreeNode{Left: &TreeNode{}}}, false},
		{"balanced root but unbalanced child", &TreeNode{Left: &TreeNode{Left: &TreeNode{Left: &TreeNode{}}}, Right: &TreeNode{Right: &TreeNode{Right: &TreeNode{}}}}, false},
	} {
		t.Run(tt.name, func(t *testing.T) {
			if got := isBalanced(tt.root); got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
