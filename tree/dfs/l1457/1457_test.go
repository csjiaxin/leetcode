package l1457

import "testing"

func TestPseudoPalindromicPaths(t *testing.T) {
	for _, tt := range []struct {
		name string
		root *TreeNode
		want int
	}{
		{"empty", nil, 0},
		{"single", &TreeNode{Val: 9}, 1},
		{"example one", &TreeNode{Val: 2, Left: &TreeNode{Val: 3, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 1}}, Right: &TreeNode{Val: 1, Right: &TreeNode{Val: 1}}}, 2},
		{"example two", &TreeNode{Val: 2, Left: &TreeNode{Val: 1, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3, Right: &TreeNode{Val: 1}}}, Right: &TreeNode{Val: 1}}, 1},
		{"two odd digits", &TreeNode{Val: 1, Left: &TreeNode{Val: 2}}, 0},
		{"all even", &TreeNode{Val: 9, Left: &TreeNode{Val: 9}}, 1},
		{"one odd", &TreeNode{Val: 1, Right: &TreeNode{Val: 1, Right: &TreeNode{Val: 1}}}, 1},
		{"digit nine and branch isolation", &TreeNode{Val: 9, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 1}}, 1},
		{"four valid branches", &TreeNode{Val: 2, Left: &TreeNode{Val: 1, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 2}}, Right: &TreeNode{Val: 3, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 2}}}, 4},
	} {
		t.Run(tt.name, func(t *testing.T) {
			for i := 0; i < 2; i++ {
				if got := pseudoPalindromicPaths(tt.root); got != tt.want {
					t.Fatalf("call %d: got %d, want %d", i, got, tt.want)
				}
			}
		})
	}
}
