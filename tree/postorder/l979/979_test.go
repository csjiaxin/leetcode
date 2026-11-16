package l979

import "testing"

func TestDistributeCoins(t *testing.T) {
	for _, tt := range []struct {
		name string
		root *TreeNode
		want int
	}{
		{"leaf", &TreeNode{Val: 1}, 0},
		{"root surplus", &TreeNode{Val: 3, Left: &TreeNode{}, Right: &TreeNode{}}, 2},
		{"sibling transfer", &TreeNode{Left: &TreeNode{Val: 3}, Right: &TreeNode{}}, 3},
		{"balanced", &TreeNode{Val: 1, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 1}}, 0},
		{"deep transfer", &TreeNode{Left: &TreeNode{Right: &TreeNode{Val: 3}}}, 3},
		{"two nodes", &TreeNode{Val: 2, Right: &TreeNode{}}, 1},
	} {
		t.Run(tt.name, func(t *testing.T) {
			if got := distributeCoins(tt.root); got != tt.want {
				t.Errorf("got %d want %d", got, tt.want)
			}
		})
	}
}
