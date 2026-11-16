package l236

import "testing"

func TestLowestCommonAncestor(t *testing.T) {
	for _, tt := range []struct {
		name       string
		p, q, want int
	}{
		{"opposite sides", 1, 2, 0}, {"ancestor", 1, 4, 1}, {"siblings", 3, 4, 1},
		{"root target", 0, 6, 0}, {"deep split", 4, 6, 0}, {"right ancestor", 2, 5, 2},
	} {
		t.Run(tt.name, func(t *testing.T) {
			nodes := make([]*TreeNode, 7)
			for i := range nodes {
				nodes[i] = &TreeNode{Val: i}
			}
			for i := 0; i < 3; i++ {
				nodes[i].Left = nodes[2*i+1]
				nodes[i].Right = nodes[2*i+2]
			}
			if got := lowestCommonAncestor(nodes[0], nodes[tt.p], nodes[tt.q]); got != nodes[tt.want] {
				t.Errorf("got node %p want %p (value %d)", got, nodes[tt.want], tt.want)
			}
		})
	}
}
