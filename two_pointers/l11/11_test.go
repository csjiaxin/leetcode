package l11

import "testing"

func TestMaxArea(t *testing.T) {
	for _, tt := range []struct {
		name   string
		height []int
		want   int
	}{
		{"example", []int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49},
		{"two bars", []int{1, 1}, 1},
		{"zeros", []int{0, 0}, 0},
		{"equal", []int{4, 4, 4, 4}, 12},
		{"increasing", []int{1, 2, 3, 4, 5}, 6},
		{"tall endpoints", []int{5, 1, 1, 5}, 15},
	} {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxArea(tt.height); got != tt.want {
				t.Errorf("maxArea(%v)=%d, want %d", tt.height, got, tt.want)
			}
		})
	}
}
