package l283

import (
	"slices"
	"testing"
)

func TestMoveZeroes(t *testing.T) {
	for _, tt := range []struct {
		name       string
		nums, want []int
	}{
		{"example", []int{0, 1, 0, 3, 12}, []int{1, 3, 12, 0, 0}},
		{"single zero", []int{0}, []int{0}}, {"no zero", []int{1, 2}, []int{1, 2}},
		{"all zero", []int{0, 0, 0}, []int{0, 0, 0}}, {"stable negatives", []int{0, -1, 0, -2, 3}, []int{-1, -2, 3, 0, 0}},
		{"single nonzero", []int{7}, []int{7}},
		{"leading zeroes", []int{0, 0, 1, 2}, []int{1, 2, 0, 0}},
		{"trailing zeroes", []int{1, 2, 0, 0}, []int{1, 2, 0, 0}},
		{"alternating", []int{1, 0, 2, 0, 3, 0}, []int{1, 2, 3, 0, 0, 0}},
		{"repeated nonzeroes", []int{0, 2, 0, 2, 1}, []int{2, 2, 1, 0, 0}},
	} {
		t.Run(tt.name, func(t *testing.T) {
			got := slices.Clone(tt.nums)
			moveZeroes(got)
			if !slices.Equal(got, tt.want) {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
