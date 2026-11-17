package main

import "testing"

func TestTwoResourceKnapsack(t *testing.T) {
	for _, tc := range []struct {
		strs       []string
		m, n, want int
	}{
		{[]string{"10", "0001", "111001", "1", "0"}, 5, 3, 4},
		{[]string{"10", "0", "1"}, 1, 1, 2}, {[]string{"0"}, 3, 0, 1},
		{[]string{"1"}, 0, 0, 0}, {nil, 2, 2, 0},
	} {
		if got := findMaxForm(tc.strs, tc.m, tc.n); got != tc.want {
			t.Fatalf("%v (%d,%d): got %d, want %d", tc.strs, tc.m, tc.n, got, tc.want)
		}
	}
}
