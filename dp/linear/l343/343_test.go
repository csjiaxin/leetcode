package main

import "testing"

func TestIntegerBreakRequiresSplit(t *testing.T) {
	for _, tc := range []struct{ n, want int }{{2, 1}, {3, 2}, {4, 4}, {10, 36}, {58, 1549681956}} {
		if got := integerBreak(tc.n); got != tc.want {
			t.Fatalf("n=%d: got %d, want %d", tc.n, got, tc.want)
		}
	}
}
