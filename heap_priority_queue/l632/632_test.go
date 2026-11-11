package l632

import (
	"slices"
	"testing"
)

func TestSmallestRange(t *testing.T) {
	for _, tt := range []struct {
		name string
		nums [][]int
		want []int
	}{
		{"example", [][]int{{4, 10, 15, 24, 26}, {0, 9, 12, 20}, {5, 18, 22, 30}}, []int{20, 24}},
		{"identical lists", [][]int{{1, 2, 3}, {1, 2, 3}, {1, 2, 3}}, []int{1, 1}},
		{"single list", [][]int{{-3, 0, 2}}, []int{-3, -3}},
		{"equal width chooses smaller left", [][]int{{1, 4}, {2, 5}}, []int{1, 2}},
		{"singletons", [][]int{{-10}, {0}, {10}}, []int{-10, 10}},
	} {
		t.Run(tt.name, func(t *testing.T) {
			if got := smallestRange(tt.nums); !slices.Equal(got, tt.want) {
				t.Errorf("%s: got %v want %v", t.Name(), got, tt.want)
			}
		})
	}
}
