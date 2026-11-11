package l1053

import (
	"slices"
	"testing"
)

func TestPrevPermOpt1(t *testing.T) {
	for _, tt := range []struct {
		name      string
		arr, want []int
	}{
		{"decreasing", []int{3, 2, 1}, []int{3, 1, 2}},
		{"already smallest", []int{1, 1, 5}, []int{1, 1, 5}},
		{"largest smaller replacement", []int{1, 9, 4, 6, 7}, []int{1, 7, 4, 6, 9}},
		{"leftmost equal replacement", []int{3, 1, 1, 3}, []int{1, 3, 1, 3}},
		{"one swap not suffix reversal", []int{3, 1, 2}, []int{2, 1, 3}},
		{"single", []int{1}, []int{1}},
		{"all equal", []int{2, 2, 2}, []int{2, 2, 2}},
		{"increasing", []int{1, 2, 3}, []int{1, 2, 3}},
	} {
		t.Run(tt.name, func(t *testing.T) {
			got := prevPermOpt1(slices.Clone(tt.arr))
			if !slices.Equal(got, tt.want) {
				t.Errorf("prevPermOpt1(%v) = %v, want %v", tt.arr, got, tt.want)
			}
		})
	}
}
