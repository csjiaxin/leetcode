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
		{"zero budget", []int{1, 0, 1, 1, 0}, 0, 2}, {"zeros", []int{0, 0, 0}, 2, 2},
		{"ones", []int{1, 1, 1}, 0, 3}, {"single zero", []int{0}, 0, 0},
		{"whole array", []int{0, 1, 0}, 3, 3},
	} {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestOnes(slices.Clone(tt.nums), tt.k); got != tt.want {
				t.Errorf("got %d want %d", got, tt.want)
			}
		})
	}
}
