package l437

import (
	"math/rand"
	"testing"
)

func TestPathSum(t *testing.T) {
	for _, tt := range []struct {
		name   string
		root   *TreeNode
		target int
		want   int
	}{
		{"empty", nil, 0, 0},
		{"single zero not empty path", &TreeNode{}, 0, 1},
		{"single match", &TreeNode{Val: 5}, 5, 1},
		{"single mismatch", &TreeNode{Val: 5}, 0, 0},
		{"example one", &TreeNode{Val: 10,
			Left:  &TreeNode{Val: 5, Left: &TreeNode{Val: 3, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: -2}}, Right: &TreeNode{Val: 2, Right: &TreeNode{Val: 1}}},
			Right: &TreeNode{Val: -3, Right: &TreeNode{Val: 11}}}, 8, 3},
		{"example two", &TreeNode{Val: 5,
			Left:  &TreeNode{Val: 4, Left: &TreeNode{Val: 11, Left: &TreeNode{Val: 7}, Right: &TreeNode{Val: 2}}},
			Right: &TreeNode{Val: 8, Left: &TreeNode{Val: 13}, Right: &TreeNode{Val: 4, Left: &TreeNode{Val: 5}, Right: &TreeNode{Val: 1}}}}, 22, 3},
		{"negative", &TreeNode{Val: -2, Right: &TreeNode{Val: -3}}, -5, 1},
		{"starts below root", &TreeNode{Val: 10, Left: &TreeNode{Val: 2, Right: &TreeNode{Val: 3}}}, 5, 1},
		{"ends before leaf", &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}}}, 3, 1},
		{"duplicate prefixes", &TreeNode{Left: &TreeNode{Left: &TreeNode{}}}, 0, 6},
		{"sibling prefixes must be removed", &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 2}}, 0, 0},
		{"zero sibling isolation", &TreeNode{Left: &TreeNode{}, Right: &TreeNode{}}, 0, 5},
		{"cancelling prefixes", &TreeNode{Val: 1, Left: &TreeNode{Val: -1, Left: &TreeNode{Val: 1}}}, 1, 3},
	} {
		t.Run(tt.name, func(t *testing.T) {
			for i := 0; i < 2; i++ {
				if got := pathSum(tt.root, tt.target); got != tt.want {
					t.Fatalf("call %d: got %d, want %d", i, got, tt.want)
				}
			}
		})
	}
}

func TestPathSumInt64Prefixes(t *testing.T) {
	// 总和恰好为 2^32；32 位前缀会溢出为 0 并误报路径。
	values := []int{1000000000, 1000000000, 1000000000, 1000000000, 294967296}
	var root *TreeNode
	for i := len(values) - 1; i >= 0; i-- {
		root = &TreeNode{Val: values[i], Left: root}
	}
	if got := pathSum(root, 0); got != 0 {
		t.Fatalf("got %d paths, want 0", got)
	}
}

// 独立的 O(n²) 枚举算法作为对照，不使用前缀和表。
func brutePathSum(root *TreeNode, target int64) int {
	if root == nil {
		return 0
	}
	var startingAt func(*TreeNode, int64) int
	startingAt = func(node *TreeNode, remaining int64) int {
		if node == nil {
			return 0
		}
		remaining -= int64(node.Val)
		count := 0
		if remaining == 0 {
			count++
		}
		return count + startingAt(node.Left, remaining) + startingAt(node.Right, remaining)
	}
	return startingAt(root, target) + brutePathSum(root.Left, target) + brutePathSum(root.Right, target)
}

func TestPathSumAgainstBruteForce(t *testing.T) {
	rng := rand.New(rand.NewSource(42))
	var build func(int) *TreeNode
	build = func(depth int) *TreeNode {
		if depth == 0 || rng.Intn(4) == 0 {
			return nil
		}
		return &TreeNode{Val: rng.Intn(7) - 3, Left: build(depth - 1), Right: build(depth - 1)}
	}
	for i := 0; i < 200; i++ {
		root := build(5)
		for target := -5; target <= 5; target++ {
			want := brutePathSum(root, int64(target))
			if got := pathSum(root, target); got != want {
				t.Fatalf("tree %d target %d: got %d, want %d", i, target, got, want)
			}
		}
	}
}
