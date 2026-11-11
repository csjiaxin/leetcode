package l713

import (
	"slices"
	"testing"
)

func TestNumSubarrayProductLessThanK(t *testing.T) {
	for _, tt := range []struct {
		name    string
		nums    []int
		k, want int
	}{
		{"example", []int{10, 5, 2, 6}, 100, 8},
		{"zero threshold", []int{1, 2, 3}, 0, 0},
		{"one threshold", []int{1, 1, 1}, 1, 0},
		{"all ones", []int{1, 1, 1}, 2, 6},
		{"strict inequality", []int{2, 5}, 10, 2},
		{"all too large", []int{5, 6}, 5, 0},
		{"shrink several times", []int{1, 2, 3}, 3, 3},
		{"single qualifies", []int{3}, 4, 1},
	} {
		t.Run(tt.name, func(t *testing.T) {
			if got := numSubarrayProductLessThanK(slices.Clone(tt.nums), tt.k); got != tt.want {
				t.Errorf("numSubarrayProductLessThanK(%v, %d) = %d, want %d", tt.nums, tt.k, got, tt.want)
			}
		})
	}
}
