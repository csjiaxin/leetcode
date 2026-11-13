package l141

import "testing"

func TestHasCycle(t *testing.T) {
	for _, tt := range []struct {
		name   string
		values []int
		pos    int
		want   bool
	}{
		{"example", []int{3, 2, 0, -4}, 1, true}, {"head entry", []int{1, 2}, 0, true},
		{"single no cycle", []int{1}, -1, false}, {"self cycle", []int{1}, 0, true},
		{"empty", nil, -1, false}, {"same values no cycle", []int{1, 1, 1}, -1, false},
		{"tail self cycle", []int{1, 1, 1}, 2, true},
	} {
		t.Run(tt.name, func(t *testing.T) {
			nodes := make([]*ListNode, len(tt.values))
			for i, x := range tt.values {
				nodes[i] = &ListNode{Val: x}
				if i > 0 {
					nodes[i-1].Next = nodes[i]
				}
			}
			var head *ListNode
			if len(nodes) > 0 {
				head = nodes[0]
				if tt.pos >= 0 {
					nodes[len(nodes)-1].Next = nodes[tt.pos]
				}
			}
			if got := hasCycle(head); got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
