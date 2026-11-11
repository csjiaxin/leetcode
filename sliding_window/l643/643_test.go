package l643

import (
	"math"
	"slices"
	"testing"
)

func TestFindMaxAverage(t *testing.T) {
	for _, tt := range []struct {
		name string
		nums []int
		k    int
		want float64
	}{
		// {"example", []int{1, 12, -5, -6, 50, 3}, 4, 12.75},
		// {"single", []int{5}, 1, 5},
		{"all negative", []int{-5, -2, -8}, 2, -3.5},
		{"whole array", []int{1, 2, 3}, 3, 2},
		{"length one", []int{-3, 7, 2}, 1, 7},
		{"best last window", []int{1, 2, 10, 10}, 2, 10},
		{"fraction", []int{1, 0, 0}, 3, 1.0 / 3},
		{"zeros", []int{0, 0, 0}, 2, 0},
	} {
		t.Run(tt.name, func(t *testing.T) {
			got := findMaxAverage(slices.Clone(tt.nums), tt.k)
			if math.IsNaN(got) || math.IsInf(got, 0) || math.Abs(got-tt.want) > 1e-5 {
				t.Errorf("findMaxAverage(%v, %d) = %v, want %v", tt.nums, tt.k, got, tt.want)
			}
		})
	}
}
