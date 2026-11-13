package l42

import "testing"

func TestTrap(t *testing.T) {
	for _, tt := range []struct {
		name   string
		height []int
		want   int
	}{
		{"example", []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}, 6},
		{"diagram", []int{4, 2, 0, 3, 2, 5, 1}, 9},
		{"single", []int{1}, 0},
		{"increasing", []int{1, 2, 3, 4}, 0},
		{"decreasing", []int{4, 3, 2, 1}, 0},
		{"equal", []int{3, 3, 3}, 0},
		{"basin", []int{3, 0, 3}, 3},
	} {
		t.Run(tt.name, func(t *testing.T) {
			if got := trap(tt.height); got != tt.want {
				t.Errorf("trap(%v)=%d, want %d", tt.height, got, tt.want)
			}
		})
	}
}
