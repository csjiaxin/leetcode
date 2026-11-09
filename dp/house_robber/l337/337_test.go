package l337

import "testing"

func TestRob(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		want int
	}{
		{
			name: "example one",
			root: &TreeNode{
				Val:   3,
				Left:  &TreeNode{Val: 2, Right: &TreeNode{Val: 3}},
				Right: &TreeNode{Val: 3, Right: &TreeNode{Val: 1}},
			},
			want: 7,
		},
		{
			name: "example two",
			root: &TreeNode{
				Val:   3,
				Left:  &TreeNode{Val: 4, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}},
				Right: &TreeNode{Val: 5, Right: &TreeNode{Val: 1}},
			},
			want: 9,
		},
		{
			name: "single house",
			root: &TreeNode{Val: 5},
			want: 5,
		},
		{
			name: "single zero house",
			root: &TreeNode{Val: 0},
			want: 0,
		},
		{
			name: "left child worth more",
			root: &TreeNode{Val: 2, Left: &TreeNode{Val: 9}},
			want: 9,
		},
		{
			name: "parent worth more than right child",
			root: &TreeNode{Val: 9, Right: &TreeNode{Val: 2}},
			want: 9,
		},
		{
			name: "take both children",
			root: &TreeNode{Val: 5, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 4}},
			want: 8,
		},
		{
			name: "take parent instead of children",
			root: &TreeNode{Val: 10, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 4}},
			want: 10,
		},
		{
			name: "all zero",
			root: &TreeNode{Left: &TreeNode{}, Right: &TreeNode{}},
			want: 0,
		},
		{
			name: "left skewed tree",
			root: &TreeNode{Val: 4, Left: &TreeNode{Val: 1, Left: &TreeNode{Val: 5, Left: &TreeNode{Val: 2}}}},
			want: 9,
		},
		{
			name: "right skewed tree",
			root: &TreeNode{Val: 4, Right: &TreeNode{Val: 1, Right: &TreeNode{Val: 5, Right: &TreeNode{Val: 2}}}},
			want: 9,
		},
		{
			name: "optimal selection mixes depths",
			root: &TreeNode{
				Val:   0,
				Left:  &TreeNode{Val: 10, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 1}},
				Right: &TreeNode{Val: 1, Left: &TreeNode{Val: 10}, Right: &TreeNode{Val: 10}},
			},
			want: 30,
		},
		{
			name: "skip two consecutive levels",
			root: &TreeNode{Val: 10, Left: &TreeNode{Val: 1, Left: &TreeNode{Val: 1, Left: &TreeNode{Val: 10}}}},
			want: 20,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rob(tt.root); got != tt.want {
				t.Errorf("rob() = %d, want %d", got, tt.want)
			}
		})
	}
}

func TestRobMaximumSkewedTree(t *testing.T) {
	var root *TreeNode
	for i := 0; i < 10000; i++ {
		root = &TreeNode{Val: 10000, Left: root}
	}
	if got := rob(root); got != 50000000 {
		t.Errorf("rob(10000-house chain with 10000 each) = %d, want 50000000", got)
	}
}
