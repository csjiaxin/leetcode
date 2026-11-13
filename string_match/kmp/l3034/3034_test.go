package l3034

import "testing"

func TestCountMatchingSubarrays(t *testing.T) {
	tests := []struct {
		name          string
		nums, pattern []int
		want          int
	}{
		{"increasing overlaps", []int{1, 2, 3, 4, 5, 6}, []int{1, 1}, 4},
		{"mixed example", []int{1, 4, 4, 1, 3, 5, 5, 3}, []int{1, 0, -1}, 2},
		{"decreasing", []int{4, 3, 2, 1}, []int{-1, -1}, 2},
		{"equal", []int{5, 5, 5, 5}, []int{0, 0}, 2},
		{"absent", []int{1, 2, 3}, []int{-1}, 0},
		{"whole array", []int{1, 3, 3, 2}, []int{1, 0, -1}, 1},
		{"minimum size", []int{2, 1}, []int{-1}, 1},
		{"alternating overlaps", []int{1, 2, 1, 2, 1}, []int{1, -1}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countMatchingSubarrays(tt.nums, tt.pattern); got != tt.want {
				t.Errorf("countMatchingSubarrays(%v, %v) = %d, want %d", tt.nums, tt.pattern, got, tt.want)
			}
		})
	}
}
