package l962

import (
	"slices"
	"testing"
)

func TestMaxWidthRamp(t *testing.T) {
	for _, tt := range []struct {
		name string
		nums []int
		want int
	}{
		{"example one", []int{6, 0, 8, 2, 1, 5}, 4}, {"example two", []int{9, 8, 1, 0, 1, 9, 4, 0, 4, 1}, 7},
		{"equal pair", []int{2, 2}, 1}, {"decreasing", []int{5, 4, 3, 2, 1}, 0},
		{"increasing", []int{1, 2, 3, 4}, 3}, {"equal endpoints", []int{5, 1, 2, 5}, 3},
		{"discard dominated left", []int{2, 3, 1, 4}, 3},
	} {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxWidthRamp(slices.Clone(tt.nums)); got != tt.want {
				t.Errorf("%s: got %d, want %d", t.Name(), got, tt.want)
			}
		})
	}
}
func TestMaxWidthRampMaximum(t *testing.T) {
	nums := make([]int, 50000)
	if got := maxWidthRamp(nums); got != 49999 {
		t.Errorf("%s: got %d, want 49999", t.Name(), got)
	}
}
