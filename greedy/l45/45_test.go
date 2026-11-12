package l45

import (
	"slices"
	"testing"
)

func TestJump(t *testing.T) {
	// LeetCode guarantees that the last index is reachable.
	for _, tt := range []struct {
		name string
		nums []int
		want int
	}{
		{"example one", []int{2, 3, 1, 1, 4}, 2},
		{"example two", []int{2, 3, 0, 1, 4}, 2},
		{"single", []int{0}, 0},
		{"two elements", []int{1, 0}, 1},
		{"one jump", []int{4, 0, 0, 0, 0}, 1},
		{"every step", []int{1, 1, 1, 1}, 3},
		{"best next position not farthest", []int{2, 4, 1, 1, 1, 0}, 2},
		{"zeros inside layer", []int{3, 0, 0, 2, 0, 0}, 2},
	} {
		t.Run(tt.name, func(t *testing.T) {
			if got := jump(slices.Clone(tt.nums)); got != tt.want {
				t.Errorf("jump(%v) = %d, want %d", tt.nums, got, tt.want)
			}
		})
	}
}
