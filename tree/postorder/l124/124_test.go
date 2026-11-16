package l124

import "testing"

func TestMaxPathSum(t *testing.T) {
	for _, tt := range []struct {
		name string
		root *TreeNode
		want int
	}{
		{"three nodes", &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}, 6},
		{"below root", &TreeNode{Val: -10, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}}, 42},
		{"negative leaf", &TreeNode{Val: -3}, -3},
		{"all negative", &TreeNode{Val: -8, Left: &TreeNode{Val: -2}, Right: &TreeNode{Val: -5}}, -2},
		{"discard negative", &TreeNode{Val: 2, Left: &TreeNode{Val: -1}}, 2},
		{"zero", &TreeNode{}, 0},
	} {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxPathSum(tt.root); got != tt.want {
				t.Errorf("got %d want %d", got, tt.want)
			}
		})
	}
}
