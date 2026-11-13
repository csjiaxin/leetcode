package l977

import (
	"slices"
	"testing"
)

func TestSortedSquares(t *testing.T) {
	for _, tt := range []struct {
		name       string
		nums, want []int
	}{
		{"example", []int{-4, -1, 0, 3, 10}, []int{0, 1, 9, 16, 100}},
		{"negative", []int{-7, -3, -1}, []int{1, 9, 49}}, {"positive", []int{1, 2, 3}, []int{1, 4, 9}},
		{"zero", []int{0}, []int{0}}, {"equal absolute", []int{-2, -2, 2, 2}, []int{4, 4, 4, 4}},
	} {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortedSquares(slices.Clone(tt.nums)); !slices.Equal(got, tt.want) {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
