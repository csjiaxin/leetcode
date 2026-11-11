package l23

import (
	"slices"
	"testing"
)

func TestMergeKLists(t *testing.T) {
	for _, tt := range []struct {
		name  string
		lists [][]int
		want  []int
	}{
		{"example", [][]int{{1, 4, 5}, {1, 3, 4}, {2, 6}}, []int{1, 1, 2, 3, 4, 4, 5, 6}},
		{"empty collection", nil, nil}, {"empty list", [][]int{{}}, nil}, {"mixed empty", [][]int{{}, {-2, 0}, {}}, []int{-2, 0}}, {"single", [][]int{{1, 1}}, []int{1, 1}},
	} {
		t.Run(tt.name, func(t *testing.T) {
			lists := make([]*ListNode, len(tt.lists))
			for i, values := range tt.lists {
				dummy := &ListNode{}
				tail := dummy
				for _, v := range values {
					tail.Next = &ListNode{Val: v}
					tail = tail.Next
				}
				lists[i] = dummy.Next
			}
			var got []int
			for node := mergeKLists(lists); node != nil; node = node.Next {
				got = append(got, node.Val)
				if len(got) > len(tt.want) {
					t.Fatalf("%s: excess nodes or cycle", t.Name())
				}
			}
			if !slices.Equal(got, tt.want) {
				t.Errorf("%s: got %v want %v", t.Name(), got, tt.want)
			}
		})
	}
}
