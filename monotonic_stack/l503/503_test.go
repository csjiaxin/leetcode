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
		{"wrap around", []int{3, 2, 1}, []int{-1, 3, 3}},
		{"negative", []int{-2, -1, -3}, []int{-1, -1, -2}},
	} {
		t.Run(tt.name, func(t *testing.T) {
			if got := nextGreaterElements(slices.Clone(tt.nums)); !slices.Equal(got, tt.want) {
				t.Errorf("%s: got %v, want %v", t.Name(), got, tt.want)
			}
		})
	}
}
