package l337

import "testing"

func TestRob(t *testing.T) {
	for _, tt := range []struct {
		name string
		root *TreeNode
		want int
	}{
		{"leaf", &TreeNode{Val: 5}, 5}, {"zero", &TreeNode{}, 0},
		{"take root", &TreeNode{Val: 3, Left: &TreeNode{Val: 2, Right: &TreeNode{Val: 3}}, Right: &TreeNode{Val: 3, Right: &TreeNode{Val: 1}}}, 7},

		{"skip root", &TreeNode{Val: 3, Left: &TreeNode{Val: 4, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}}, Right: &TreeNode{Val: 5, Right: &TreeNode{Val: 1}}}, 9},
		{"chain", &TreeNode{Val: 2, Right: &TreeNode{Val: 1, Right: &TreeNode{Val: 3}}}, 5},
		{"choose child", &TreeNode{Val: 1, Left: &TreeNode{Val: 8}}, 8},
	} {
		t.Run(tt.name, func(t *testing.T) {
			if got := rob(tt.root); got != tt.want {
				t.Errorf("got %d want %d", got, tt.want)
			}
		})
	}
}
