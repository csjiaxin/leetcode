package l55

import (
	"slices"
	"testing"
)

func TestCanJump(t *testing.T) {
	for _, tt := range []struct {
		name string
		nums []int
		want bool
	}{
		{"example reachable", []int{2, 3, 1, 1, 4}, true},
		{"example blocked", []int{3, 2, 1, 0, 4}, false},
		{"single zero", []int{0}, true},
		{"blocked at start", []int{0, 1}, false},
		{"exact reach", []int{1, 1, 0}, true},
		{"jump over zero", []int{2, 0, 0}, true},
		{"unreachable large value", []int{1, 0, 100}, false},
		{"jump beyond end", []int{5, 0, 0}, true},
	} {
		t.Run(tt.name, func(t *testing.T) {
			if got := canJump(slices.Clone(tt.nums)); got != tt.want {
				t.Errorf("canJump(%v) = %v, want %v", tt.nums, got, tt.want)
			}
		})
	}
}
