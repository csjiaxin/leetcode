package l84

import (
	"slices"
	"testing"
)

func TestLargestRectangleArea(t *testing.T) {
	for _, tt := range []struct {
		name   string
		values []int
		want   int
	}{
		{"example", []int{2, 1, 5, 6, 2, 3}, 10}, {"two", []int{2, 4}, 4},
		{"ascending", []int{1, 2, 3, 4}, 6}, {"descending", []int{4, 3, 2, 1}, 6},
		{"equal", []int{5, 5, 5, 5}, 20}, {"zero", []int{0}, 0},
		{"split", []int{2, 0, 2}, 2}, {"valley", []int{2, 1, 2}, 3},
	} {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestRectangleArea(slices.Clone(tt.values)); got != tt.want {
				t.Errorf("got %d want %d", got, tt.want)
			}
		})
	}
}
