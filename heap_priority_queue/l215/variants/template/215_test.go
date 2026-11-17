package l215

import (
	"slices"
	"testing"
)

func TestFindKthLargest(t *testing.T) {
	for _, tt := range []struct {
		name    string
		nums    []int
		k, want int
	}{
		{"example", []int{3, 2, 1, 5, 6, 4}, 2, 5}, {"duplicates", []int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4, 4},
		{"single", []int{1}, 1, 1}, {"equal", []int{5, 5, 5}, 2, 5},
		{"negative", []int{-1, -5, -2}, 2, -2}, {"smallest", []int{1, 3, 2}, 3, 1},
	} {
		t.Run(tt.name, func(t *testing.T) {
			if got := findKthLargest(slices.Clone(tt.nums), tt.k); got != tt.want {
				t.Errorf("got %d want %d", got, tt.want)
			}
		})
	}
}
