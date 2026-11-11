package l84

import (
	"slices"
	"testing"
)

func TestLargestRectangleArea(t *testing.T) {
	for _, tt := range []struct {
		name    string
		heights []int
		want    int
	}{
		{"example one", []int{2, 1, 2}, 3},
		{"example one", []int{2, 1, 5, 6, 2, 3}, 10},
		// {"example two", []int{2, 4}, 4},
		// {"single", []int{7}, 7},
		// {"zero", []int{0}, 0}, {"equal", []int{2, 2, 2}, 6},
		// {"increasing", []int{1, 2, 3, 4}, 6}, {"decreasing", []int{4, 3, 2, 1}, 6},
		// {"zero separates", []int{2, 0, 2}, 2}, {"wide low rectangle", []int{2, 1, 2}, 3},
	} {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestRectangleArea(slices.Clone(tt.heights)); got != tt.want {
				t.Errorf("%s: got %d, want %d", t.Name(), got, tt.want)
			}
		})
	}
}
func TestLargestRectangleAreaMaximum(t *testing.T) {
	heights := make([]int, 100000)
	for i := range heights {
		heights[i] = 10000
	}
	if got := largestRectangleArea(heights); got != 1000000000 {
		t.Errorf("%s: got %d, want 1000000000", t.Name(), got)
	}
}
