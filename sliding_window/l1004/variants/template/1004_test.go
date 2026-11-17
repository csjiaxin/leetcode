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
		{"example", []int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}, 2, 6},
		{"all ones", []int{1, 1, 1}, 0, 3}, {"all zeros", []int{0, 0, 0}, 2, 2},
		{"no flips", []int{1, 0, 1, 1}, 0, 2}, {"full budget", []int{0, 1, 0}, 3, 3},
		{"single blocked", []int{0}, 0, 0}, {"single flipped", []int{0}, 1, 1},
	} {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestOnes(slices.Clone(tt.nums), tt.k); got != tt.want {
				t.Errorf("got %d want %d", got, tt.want)
			}
		})
	}
}
