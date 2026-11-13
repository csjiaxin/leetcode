package l16

import (
	"slices"
	"testing"
)

func TestThreeSumClosest(t *testing.T) {
	for _, tt := range []struct {
		name         string
		nums         []int
		target, want int
	}{
		{"example", []int{-1, 2, 1, -4}, 1, 2},
		{"zero", []int{0, 0, 0}, 1, 0},
		{"equal", []int{1, 1, 1, 1}, -100, 3},
		{"exact", []int{-2, 0, 1, 2}, 0, 0},
		{"negative", []int{-8, -5, -3, -1}, -10, -9},
		{"three elements", []int{2, 4, 7}, 20, 13},
		{"unsorted closest below", []int{4, 0, 1, 1}, 100, 6},
		{"unsorted closest above", []int{5, -2, 3, 1}, -10, 2},
		{"duplicate fixed values", []int{-1, -1, -1, 2, 2}, 1, 0},
		{"value limits", []int{1000, -1000, 1000, -1000}, 10000, 1000},
	} {
		t.Run(tt.name, func(t *testing.T) {
			if got := threeSumClosest(slices.Clone(tt.nums), tt.target); got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}
