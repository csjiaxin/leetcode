package l129

import "testing"

func TestSumNumbers(t *testing.T) {
	for _, tt := range []struct {
		name string
		root *TreeNode
		want int
	}{
		{"empty", nil, 0},
		{"single", &TreeNode{Val: 7}, 7},
		{"zero", &TreeNode{}, 0},
		{"example one", &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}, 25},
		{"example two", &TreeNode{Val: 4, Left: &TreeNode{Val: 9, Left: &TreeNode{Val: 5}, Right: &TreeNode{Val: 1}}, Right: &TreeNode{Val: 0}}, 1026},
		{"embedded zero", &TreeNode{Val: 1, Left: &TreeNode{Val: 0, Left: &TreeNode{Val: 1}}, Right: &TreeNode{Val: 2}}, 113},
		{"leading zero", &TreeNode{Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 2}}, 3},
		{"right chain", &TreeNode{Val: 1, Right: &TreeNode{Val: 2, Right: &TreeNode{Val: 3}}}, 123},
		{"branch isolation", &TreeNode{Val: 2, Left: &TreeNode{Val: 3, Left: &TreeNode{Val: 4}}, Right: &TreeNode{Val: 5, Right: &TreeNode{Val: 6}}}, 490},
	} {
		t.Run(tt.name, func(t *testing.T) {
			for i := 0; i < 2; i++ {
				if got := sumNumbers(tt.root); got != tt.want {
					t.Fatalf("call %d: got %d, want %d", i, got, tt.want)
				}
			}
		})
	}
}
