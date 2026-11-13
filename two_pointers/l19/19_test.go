package l19

import (
	"slices"
	"testing"
)

func TestRemoveNthFromEnd(t *testing.T) {
	for _, tt := range []struct {
		name   string
		values []int
		n      int
		want   []int
	}{
		{"example", []int{1, 2, 3, 4, 5}, 2, []int{1, 2, 3, 5}},
		{"single", []int{1}, 1, nil}, {"head", []int{1, 2, 3}, 3, []int{2, 3}},
		{"tail", []int{1, 2, 3}, 1, []int{1, 2}}, {"two", []int{1, 2}, 2, []int{2}},
		{"same values", []int{2, 2, 2}, 2, []int{2, 2}},
	} {
		t.Run(tt.name, func(t *testing.T) {
			dummy := &ListNode{}
			tail := dummy
			for _, x := range tt.values {
				tail.Next = &ListNode{Val: x}
				tail = tail.Next
			}
			var got []int
			for node := removeNthFromEnd(dummy.Next, tt.n); node != nil; node = node.Next {
				got = append(got, node.Val)
				if len(got) > len(tt.want) {
					t.Fatal("excess nodes or cycle")
				}
			}
			if !slices.Equal(got, tt.want) {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
