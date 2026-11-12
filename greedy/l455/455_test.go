package l455

import (
	"slices"
	"testing"
)

func TestFindContentChildren(t *testing.T) {
	for _, tt := range []struct {
		name string
		g, s []int
		want int
	}{
		{"example one", []int{1, 2, 3}, []int{1, 1}, 1},
		{"example two", []int{1, 2}, []int{1, 2, 3}, 2},
		{"no cookies", []int{1, 2}, nil, 0},
		{"unsorted", []int{3, 1, 2}, []int{2, 1}, 2},
		{"too small", []int{5, 6}, []int{1, 2, 3}, 0},
		{"duplicates", []int{2, 2, 2}, []int{2, 2}, 2},
		{"large cookie first", []int{1, 3}, []int{3, 1}, 2},
	} {
		t.Run(tt.name, func(t *testing.T) {
			if got := findContentChildren(slices.Clone(tt.g), slices.Clone(tt.s)); got != tt.want {
				t.Errorf("findContentChildren(%v, %v) = %d, want %d", tt.g, tt.s, got, tt.want)
			}
		})
	}
}
