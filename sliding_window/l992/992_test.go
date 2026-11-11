package l992

import (
	"slices"
	"testing"
)

func TestSubarraysWithKDistinct(t *testing.T) {
	for _, tt := range []struct {
		name    string
		nums    []int
		k, want int
	}{
		{"example one", []int{1, 2, 1, 2, 3}, 2, 7},
		{"example two", []int{1, 2, 1, 3, 4}, 3, 3},
		{"single", []int{1}, 1, 1},
		{"all equal", []int{1, 1, 1}, 1, 6},
		{"one distinct runs", []int{1, 1, 2, 2}, 1, 6},
		{"all distinct pairs", []int{1, 2, 3}, 2, 2},
		{"whole array only", []int{1, 2, 3}, 3, 1},
		{"insufficient distinct", []int{1, 1, 1}, 2, 0},
	} {
		t.Run(tt.name, func(t *testing.T) {
			if got := subarraysWithKDistinct(slices.Clone(tt.nums), tt.k); got != tt.want {
				t.Errorf("subarraysWithKDistinct(%v, %d) = %d, want %d", tt.nums, tt.k, got, tt.want)
			}
		})
	}
}

func TestSubarraysWithKDistinctMaximum(t *testing.T) {
	const n = 20000
	nums := make([]int, n)
	for i := range nums {
		nums[i] = 1
	}
	want := int(int64(n) * int64(n+1) / 2)
	if got := subarraysWithKDistinct(nums, 1); got != want {
		t.Errorf("maximum equal input: got %d, want %d", got, want)
	}
}
