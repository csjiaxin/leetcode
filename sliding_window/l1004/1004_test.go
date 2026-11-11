package l1004

import (
	"slices"
	"testing"
)

func TestLongestOnes(t *testing.T) {
	for _, tt := range []struct {
		name    string
		nums    []int
		k, want int
	}{
		{"example one", []int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}, 2, 6},
		{"example two", []int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1}, 3, 10},
		{"no flips", []int{1, 1, 0, 1}, 0, 2},
		{"all zeros no flips", []int{0, 0, 0}, 0, 0},
		{"all zeros", []int{0, 0, 0}, 2, 2},
		{"all ones", []int{1, 1, 1}, 0, 3},
		{"flip whole array", []int{0, 1, 0}, 3, 3},
		{"single zero", []int{0}, 1, 1},
	} {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestOnes(slices.Clone(tt.nums), tt.k); got != tt.want {
				t.Errorf("longestOnes(%v, %d) = %d, want %d", tt.nums, tt.k, got, tt.want)
			}
		})
	}
}
