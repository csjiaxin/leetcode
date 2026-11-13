package l142

import "testing"

func TestDetectCycle(t *testing.T) {
	for _, tt := range []struct {
		name   string
		values []int
		pos    int
	}{
		{"example", []int{3, 2, 0, -4}, 1}, {"head entry", []int{1, 2}, 0},
		{"no cycle", []int{1}, -1}, {"self cycle", []int{1}, 0}, {"empty", nil, -1},
		{"same values", []int{7, 7, 7, 7}, 2}, {"diagram", []int{1, 2, 3, 4, 5}, 2},
		{"tail self cycle", []int{1, 2, 3}, 2},
	} {
		t.Run(tt.name, func(t *testing.T) {
			nodes := make([]*ListNode, len(tt.values))
			for i, x := range tt.values {
				nodes[i] = &ListNode{Val: x}
				if i > 0 {
					nodes[i-1].Next = nodes[i]
				}
			}
			var head, want *ListNode
			if len(nodes) > 0 {
				head = nodes[0]
				if tt.pos >= 0 {
					want = nodes[tt.pos]
					nodes[len(nodes)-1].Next = want
				}
			}
			if got := detectCycle(head); got != want {
				t.Errorf("got node %p, want entry %p at index %d", got, want, tt.pos)
			}
		})
	}
}
