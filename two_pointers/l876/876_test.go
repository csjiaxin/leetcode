package l876

import "testing"

func TestMiddleNode(t *testing.T) {
	for _, tt := range []struct {
		name   string
		values []int
		index  int
	}{
		{"odd", []int{1, 2, 3, 4, 5}, 2}, {"even second middle", []int{1, 2, 3, 4, 5, 6}, 3},
		{"single", []int{1}, 0}, {"two", []int{1, 2}, 1}, {"equal values", []int{7, 7, 7, 7}, 2},
	} {
		t.Run(tt.name, func(t *testing.T) {
			nodes := make([]*ListNode, len(tt.values))
			for i, x := range tt.values {
				nodes[i] = &ListNode{Val: x}
				if i > 0 {
					nodes[i-1].Next = nodes[i]
				}
			}
			if got := middleNode(nodes[0]); got != nodes[tt.index] {
				t.Errorf("got node %p, want original node %p", got, nodes[tt.index])
			}
		})
	}
}
