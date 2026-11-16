package l113

import (
	"reflect"
	"testing"
)

func TestPathSum(t *testing.T) {
	for _, tt := range []struct {
		name   string
		root   *TreeNode
		target int
		want   [][]int
	}{
		{"empty", nil, 0, [][]int{}},
		{"single match", &TreeNode{Val: 5}, 5, [][]int{{5}}},
		{"single mismatch", &TreeNode{Val: 5}, 0, [][]int{}},
		{"example", &TreeNode{Val: 5,
			Left:  &TreeNode{Val: 4, Left: &TreeNode{Val: 11, Left: &TreeNode{Val: 7}, Right: &TreeNode{Val: 2}}},
			Right: &TreeNode{Val: 8, Left: &TreeNode{Val: 13}, Right: &TreeNode{Val: 4, Left: &TreeNode{Val: 5}, Right: &TreeNode{Val: 1}}}},
			22, [][]int{{5, 4, 11, 2}, {5, 8, 4, 5}}},
		{"no path", &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}, 5, [][]int{}},
		{"not a leaf", &TreeNode{Val: 1, Left: &TreeNode{Val: 2}}, 1, [][]int{}},
		{"negative", &TreeNode{Val: -2, Right: &TreeNode{Val: -3}}, -5, [][]int{{-2, -3}}},
		{"no positive pruning", &TreeNode{Val: 5, Left: &TreeNode{Val: -4}}, 1, [][]int{{5, -4}}},
		{"duplicate paths", &TreeNode{Left: &TreeNode{}, Right: &TreeNode{}}, 0, [][]int{{0, 0}, {0, 0}}},
		{"snapshot and sibling isolation", &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}, Right: &TreeNode{Val: 3, Left: &TreeNode{Val: 2}}}, 6, [][]int{{1, 2, 3}, {1, 3, 2}}},
	} {
		t.Run(tt.name, func(t *testing.T) {
			for i := 0; i < 2; i++ {
				got := pathSum(tt.root, tt.target)
				// 模板按左到右遍历，必须保留每条路径内部的祖先顺序。
				if !reflect.DeepEqual(got, tt.want) {
					t.Fatalf("call %d: got %v, want %v", i, got, tt.want)
				}
			}
		})
	}
}

func TestPathSnapshotsIndependent(t *testing.T) {
	root := &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}, Right: &TreeNode{Val: 3, Left: &TreeNode{Val: 2}}}
	got := pathSum(root, 6)
	if len(got) != 2 || len(got[0]) != 3 || len(got[1]) != 3 {
		t.Fatalf("unexpected paths %v", got)
	}
	got[0][0] = 99
	if got[1][0] != 1 || root.Val != 1 {
		t.Fatalf("paths or input share mutable storage: %v", got)
	}
	want := [][]int{{1, 2, 3}, {1, 3, 2}}
	if next := pathSum(root, 6); !reflect.DeepEqual(next, want) {
		t.Fatalf("later call got %v, want %v", next, want)
	}
}
