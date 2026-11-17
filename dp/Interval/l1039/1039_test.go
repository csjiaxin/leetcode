package main

import "testing"

func TestTriangulationIntervalBoundaries(t *testing.T) {
	for _, tc := range []struct {
		values []int
		want   int
	}{
		{[]int{1, 2, 3}, 6}, {[]int{3, 7, 4, 5}, 144}, {[]int{1, 3, 1, 4, 1, 5}, 13},
	} {
		if got := minScoreTriangulation(tc.values); got != tc.want {
			t.Fatalf("%v: got %d, want %d", tc.values, got, tc.want)
		}
	}
}
