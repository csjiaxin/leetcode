package l992

import (
	"slices"
	"testing"
)

func TestSubarraysWithKDistinct(t *testing.T) {
	for _, tt := range []struct {
		name    string
		nums    []int
		k, want int
	}{
		{"example", []int{1, 2, 1, 2, 3}, 2, 7}, {"three", []int{1, 2, 1, 3, 4}, 3, 3},
		{"all equal", []int{1, 1, 1}, 1, 6}, {"too many required", []int{1, 1, 1}, 2, 0},
		{"single", []int{1}, 1, 1}, {"all distinct", []int{1, 2, 3}, 2, 2},
	} {
		t.Run(tt.name, func(t *testing.T) {
			if got := subarraysWithKDistinct(slices.Clone(tt.nums), tt.k); got != tt.want {
				t.Errorf("got %d want %d", got, tt.want)
			}
		})
	}
}
