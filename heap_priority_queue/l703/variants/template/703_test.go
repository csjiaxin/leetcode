package l703

import (
	"slices"
	"testing"
)

func TestKthLargest(t *testing.T) {
	for _, tt := range []struct {
		name             string
		k                int
		nums, adds, want []int
	}{
		{"example", 3, []int{4, 5, 8, 2}, []int{3, 5, 10, 9, 4}, []int{4, 5, 8, 8, 8}},
		{"empty initial", 1, nil, []int{3, 1, 5}, []int{3, 3, 5}},
		{"one short initial", 3, []int{1, 2}, []int{3, 4}, []int{1, 2}},
		{"duplicates", 2, []int{5, 5}, []int{5, 4, 6}, []int{5, 5, 5}},
		{"negative", 2, []int{-5, -2}, []int{-3, -1}, []int{-3, -2}},
		{"descending", 2, []int{10, 5}, []int{1, 2, 3}, []int{5, 5, 5}},
	} {
		t.Run(tt.name, func(t *testing.T) {
			stream := Constructor(tt.k, slices.Clone(tt.nums))
			for i, x := range tt.adds {
				if got := stream.Add(x); got != tt.want[i] {
					t.Errorf("Add(%d) step %d got %d want %d", x, i, got, tt.want[i])
				}
			}
		})
	}
}
