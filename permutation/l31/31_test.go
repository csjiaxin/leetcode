package l31

import (
	"slices"
	"testing"
)

func TestNextPermutation(t *testing.T) {
	for _, tt := range []struct {
		name       string
		nums, want []int
	}{
		{"increasing", []int{1, 2, 3}, []int{1, 3, 2}},
		{"largest wraps", []int{3, 2, 1}, []int{1, 2, 3}},
		{"duplicates", []int{1, 1, 5}, []int{1, 5, 1}},
		{"reverse suffix", []int{1, 3, 2}, []int{2, 1, 3}},
		{"smallest larger replacement", []int{2, 3, 1, 3, 3}, []int{2, 3, 3, 1, 3}},
		{"single", []int{1}, []int{1}},
		{"all equal", []int{2, 2, 2}, []int{2, 2, 2}},
		{"zeros", []int{0, 1, 0}, []int{1, 0, 0}},
	} {
		t.Run(tt.name, func(t *testing.T) {
			got := slices.Clone(tt.nums)
			nextPermutation(got)
			if !slices.Equal(got, tt.want) {
				t.Errorf("nextPermutation(%v) = %v, want %v", tt.nums, got, tt.want)
			}
		})
	}
}
