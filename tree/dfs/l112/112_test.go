package l112

import "testing"

func TestHasPathSum(t *testing.T) {
	for _, tt := range []struct {
		name   string
		root   *TreeNode
		target int
		want   bool
	}{
		{"empty", nil, 0, false},
		{"single match", &TreeNode{Val: 0}, 0, true},
		{"single mismatch", &TreeNode{Val: 1}, 0, false},
		{"example", &TreeNode{Val: 5,
			Left:  &TreeNode{Val: 4, Left: &TreeNode{Val: 11, Left: &TreeNode{Val: 7}, Right: &TreeNode{Val: 2}}},
			Right: &TreeNode{Val: 8, Left: &TreeNode{Val: 13}, Right: &TreeNode{Val: 4, Right: &TreeNode{Val: 1}}}}, 22, true},
		{"no path", &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}, 5, false},
		{"not a leaf", &TreeNode{Val: 1, Left: &TreeNode{Val: 2}}, 1, false},
		{"negative", &TreeNode{Val: -2, Right: &TreeNode{Val: -3}}, -5, true},
		{"no positive pruning", &TreeNode{Val: 5, Left: &TreeNode{Val: -4}}, 1, true},
		{"right branch independent", &TreeNode{Val: 1, Left: &TreeNode{Val: 8}, Right: &TreeNode{Val: 2}}, 3, true},
	} {
		t.Run(tt.name, func(t *testing.T) {
			for i := 0; i < 2; i++ {
				if got := hasPathSum(tt.root, tt.target); got != tt.want {
					t.Fatalf("call %d: got %v, want %v", i, got, tt.want)
				}
			}
		})
	}
}
