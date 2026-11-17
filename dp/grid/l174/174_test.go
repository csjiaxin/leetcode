package main

import "testing"

func TestReverseHealthRequirement(t *testing.T) {
	for _, tc := range []struct {
		grid [][]int
		want int
	}{
		{[][]int{{-2, -3, 3}, {-5, -10, 1}, {10, 30, -5}}, 7},
		{[][]int{{0}}, 1}, {[][]int{{-3}}, 4}, {[][]int{{5}}, 1},
		{[][]int{{-2, -3}}, 6}, {[][]int{{-2}, {-3}}, 6},
	} {
		if got := calculateMinimumHP(tc.grid); got != tc.want {
			t.Fatalf("%v: got %d, want %d", tc.grid, got, tc.want)
		}
	}
}
