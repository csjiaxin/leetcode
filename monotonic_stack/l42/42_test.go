package l42

import (
	"slices"
	"testing"
)

func TestTrap(t *testing.T) {
	for _, tt := range []struct {
		name   string
		height []int
		want   int
	}{
		{"example one", []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}, 6},
		{"example two", []int{4, 2, 0, 3, 2, 5}, 9},
		{"single", []int{5}, 0}, {"two walls only", []int{5, 5}, 0},
		{"basin", []int{3, 0, 3}, 3}, {"unequal walls", []int{3, 0, 1}, 1},
		{"flat", []int{2, 2, 2}, 0}, {"increasing", []int{1, 2, 3, 4}, 0},
		{"decreasing", []int{4, 3, 2, 1}, 0}, {"plateau bottom", []int{3, 1, 1, 3}, 4},
	} {
		t.Run(tt.name, func(t *testing.T) {
			if got := trap(slices.Clone(tt.height)); got != tt.want {
				t.Errorf("%s: got %d, want %d", t.Name(), got, tt.want)
			}
		})
	}
}
func TestTrapMaximum(t *testing.T) {
	height := make([]int, 20000)
	height[0], height[len(height)-1] = 100000, 100000
	if got := trap(height); got != 1999800000 {
		t.Errorf("%s: got %d, want 1999800000", t.Name(), got)
	}
}
