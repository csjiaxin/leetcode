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
		{"example", []int{10, 5, 2, 6}, 100, 8}, {"zero threshold", []int{1, 2, 3}, 0, 0},
		{"one threshold", []int{1, 1}, 1, 0}, {"all ones", []int{1, 1, 1}, 2, 6},
		{"strict inequality", []int{10}, 10, 0}, {"shrink", []int{1, 2, 3}, 3, 3},
	} {
		t.Run(tt.name, func(t *testing.T) {
			if got := numSubarrayProductLessThanK(slices.Clone(tt.nums), tt.k); got != tt.want {
				t.Errorf("got %d want %d", got, tt.want)
			}
		})
	}
}
