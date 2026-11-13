package l80

import (
	"slices"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	for _, tt := range []struct {
		name       string
		nums, want []int
	}{
		{"example", []int{1, 1, 1, 2, 2, 3}, []int{1, 1, 2, 2, 3}},
		{"groups", []int{0, 0, 1, 1, 1, 1, 2, 3, 3}, []int{0, 0, 1, 1, 2, 3, 3}},
		{"single", []int{1}, []int{1}}, {"equal", []int{2, 2, 2, 2}, []int{2, 2}},
		{"already valid", []int{-1, -1, 0, 1, 1}, []int{-1, -1, 0, 1, 1}},
	} {
		t.Run(tt.name, func(t *testing.T) {
			nums := slices.Clone(tt.nums)
			k := removeDuplicates(nums)
			if k < 0 || k > len(nums) {
				t.Fatalf("invalid length %d", k)
			}
			if k != len(tt.want) || !slices.Equal(nums[:k], tt.want) {
				t.Errorf("got %v (k=%d), want %v", nums[:k], k, tt.want)
			}
		})
	}
}
