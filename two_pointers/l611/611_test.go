package l611

import (
	"slices"
	"testing"
)

func TestTriangleNumber(t *testing.T) {
	for _, tt := range []struct {
		name string
		nums []int
		want int
	}{
		{"example", []int{2, 2, 3, 4}, 3}, {"all valid", []int{4, 2, 3, 4}, 4},
		{"zeros", []int{0, 0, 0}, 0}, {"with zero", []int{0, 2, 2, 3}, 1},
		{"degenerate", []int{1, 2, 3}, 0}, {"short", []int{1, 2}, 0},
	} {
		t.Run(tt.name, func(t *testing.T) {
			if got := triangleNumber(slices.Clone(tt.nums)); got != tt.want {
				t.Errorf("got %d want %d", got, tt.want)
			}
		})
	}
}
