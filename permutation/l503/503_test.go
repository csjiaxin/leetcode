package l503

import (
	"slices"
	"testing"
)

func TestNextGreaterElements(t *testing.T) {
	for _, tt := range []struct {
		name       string
		nums, want []int
	}{
		{"example one", []int{1, 2, 1}, []int{2, -1, 2}},
		{"example two", []int{1, 2, 3, 4, 3}, []int{2, 3, 4, -1, 4}},
		{"single", []int{5}, []int{-1}},
		{"equal", []int{2, 2, 2}, []int{-1, -1, -1}},
		{"decreasing wraps", []int{5, 4, 3, 2, 1}, []int{-1, 5, 5, 5, 5}},
		{"negative values", []int{-2, -1, -3}, []int{-1, -1, -2}},
		{"first greater not maximum", []int{2, 1, 3, 4}, []int{3, 3, 4, -1}},
	} {
		t.Run(tt.name, func(t *testing.T) {
			got := nextGreaterElements(slices.Clone(tt.nums))
			if !slices.Equal(got, tt.want) {
				t.Errorf("nextGreaterElements(%v) = %v, want %v", tt.nums, got, tt.want)
			}
		})
	}
}
