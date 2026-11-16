package l1448

import (
	"reflect"
	"testing"
)

func TestGoodNodes(t *testing.T) {
	for _, tt := range []struct {
		name string
		root *TreeNode
		want int
	}{
		{"empty", nil, 0},
		{"single", &TreeNode{Val: -5}, 1},
		{"example one", &TreeNode{Val: 3, Left: &TreeNode{Val: 1, Left: &TreeNode{Val: 3}}, Right: &TreeNode{Val: 4, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 5}}}, 4},
		{"example two", &TreeNode{Val: 3, Left: &TreeNode{Val: 3, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 2}}}, 3},
		{"negative", &TreeNode{Val: -5, Left: &TreeNode{Val: -10}, Right: &TreeNode{Val: -3}}, 2},
		{"equal", &TreeNode{Val: 2, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 2}}, 3},
		{"branch isolation", &TreeNode{Val: 2, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 3}}, 3},
		{"decreasing", &TreeNode{Val: 3, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 1}}}, 1},
	} {
		t.Run(tt.name, func(t *testing.T) {
			for i := 0; i < 2; i++ {
				if got := goodNodes(tt.root); got != tt.want {
					t.Fatalf("call %d: got %d, want %d", i, got, tt.want)
				}
			}
		})
	}
}

func TestPathDFSLifecycle(t *testing.T) {
	root := &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}}, Right: &TreeNode{Val: 3}}
	var events, path []int
	parents := map[int]int{1: 0, 2: 1, 4: 3, 3: 1}
	PathDFS(root, 0, func(node *TreeNode, parent int) (int, func()) {
		if parent != parents[node.Val] {
			t.Errorf("node %d: parent state %d, want %d", node.Val, parent, parents[node.Val])
		}
		path = append(path, node.Val)
		events = append(events, node.Val)
		return parent + node.Val, func() {
			if path[len(path)-1] != node.Val {
				t.Fatalf("undo node %d with path %v", node.Val, path)
			}
			events = append(events, -node.Val)
			path = path[:len(path)-1]
		}
	})
	want := []int{1, 2, 4, -4, -2, 3, -3, -1}
	if !reflect.DeepEqual(events, want) || len(path) != 0 {
		t.Fatalf("events %v, want %v; restored path %v", events, want, path)
	}
	PathDFS[int](nil, 0, func(*TreeNode, int) (int, func()) {
		t.Fatal("nil root must not invoke enter")
		return 0, nil
	})
}
