package l234

import "testing"

func TestIsPalindrome(t *testing.T) {
	for _, tt := range []struct {
		name   string
		values []int
		want   bool
	}{
		{"even", []int{1, 2, 2, 1}, true}, {"different", []int{1, 2}, false},
		{"odd", []int{1, 2, 3, 2, 1}, true}, {"single", []int{1}, true},
		{"equal", []int{2, 2, 2}, true}, {"inner mismatch", []int{1, 2, 3, 1}, false},
	} {
		t.Run(tt.name, func(t *testing.T) {
			dummy := &ListNode{}
			tail := dummy
			for _, x := range tt.values {
				tail.Next = &ListNode{Val: x}
				tail = tail.Next
			}
			if got := isPalindrome(dummy.Next); got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
