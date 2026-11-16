package l968

import "testing"

func TestMinCameraCover(t *testing.T) {
	for _, tt := range []struct {
		name string
		root *TreeNode
		want int
	}{
		{"leaf needs camera", &TreeNode{}, 1},
		{"two nodes", &TreeNode{Left: &TreeNode{}}, 1},
		{"three node chain", &TreeNode{Left: &TreeNode{Left: &TreeNode{}}}, 1},
		{"four node chain", &TreeNode{Right: &TreeNode{Right: &TreeNode{Right: &TreeNode{}}}}, 2},
		{"two leaves", &TreeNode{Left: &TreeNode{}, Right: &TreeNode{}}, 1},
		{"full three levels", &TreeNode{Left: &TreeNode{Left: &TreeNode{}, Right: &TreeNode{}}, Right: &TreeNode{Left: &TreeNode{}, Right: &TreeNode{}}}, 2},
		{"uncovered sibling of camera", &TreeNode{Left: &TreeNode{}, Right: &TreeNode{Left: &TreeNode{}}}, 2},
	} {
		t.Run(tt.name, func(t *testing.T) {
			if got := minCameraCover(tt.root); got != tt.want {
				t.Errorf("got %d want %d", got, tt.want)
			}
		})
	}
}
