package main

import "testing"

func TestLISPrefixStates(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		want int
	}{
		{nil, 0}, {[]int{10, 9, 2, 5, 3, 7, 101, 18}, 4},
		{[]int{0, 1, 0, 3, 2, 3}, 4}, {[]int{7, 7, 7}, 1},
		{[]int{5, 4, 3, 2, 1}, 1}, {[]int{1, 2, 3, 4}, 4},
	} {
		if got := lengthOfLIS(tc.nums); got != tc.want {
			t.Fatalf("%v: got %d, want %d", tc.nums, got, tc.want)
		}
	}
}
