package l88

import (
	"slices"
	"testing"
)

func TestMerge(t *testing.T) {
	for _, tt := range []struct {
		name string
		a    []int
		m    int
		b    []int
		n    int
		want []int
	}{
		{"example", []int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3, []int{1, 2, 2, 3, 5, 6}},
		{"empty second", []int{1}, 1, nil, 0, []int{1}},
		{"empty first", []int{0}, 0, []int{1}, 1, []int{1}},
		{"second smaller", []int{4, 5, 0, 0}, 2, []int{1, 2}, 2, []int{1, 2, 4, 5}},
		{"duplicates negatives", []int{-3, -1, 0, 0, 0}, 2, []int{-3, 0, 2}, 3, []int{-3, -3, -1, 0, 2}},
	} {
		t.Run(tt.name, func(t *testing.T) {
			got := slices.Clone(tt.a)
			merge(got, tt.m, slices.Clone(tt.b), tt.n)
			if !slices.Equal(got, tt.want) {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
