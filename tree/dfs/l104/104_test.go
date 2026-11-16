package l104

import "testing"

func TestMaxDepth(t *testing.T) {
	for _, tt := range []struct {
		name string
		root *TreeNode
		want int
	}{
		{"empty", nil, 0}, {"leaf", &TreeNode{Val: 1}, 1},
		{"example", &TreeNode{Val: 3, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}}, 3},
		{"left chain", &TreeNode{Left: &TreeNode{Left: &TreeNode{}}}, 3},
		{"right chain", &TreeNode{Right: &TreeNode{Right: &TreeNode{Right: &TreeNode{}}}}, 4},
	} {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxDepth(tt.root); got != tt.want {
				t.Errorf("got %d want %d", got, tt.want)
			}
		})
	}
}
