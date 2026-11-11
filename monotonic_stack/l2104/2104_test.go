package l2104

import (
	"slices"
	"testing"
)

func TestSubArrayRanges(t *testing.T) {
	for _, tt := range []struct {
		name string
		nums []int
		want int64
	}{
		{"example one", []int{1, 2, 3}, 4}, {"example two", []int{1, 3, 3}, 4},
		{"example three", []int{4, -2, -3, 4, 1}, 59}, {"single", []int{7}, 0},
		{"equal", []int{2, 2, 2}, 0}, {"negative", []int{-3, -2, -1}, 4},
		{"large difference", []int{-1000000000, 1000000000, -1000000000}, 6000000000},
	} {
		t.Run(tt.name, func(t *testing.T) {
			if got := subArrayRanges(slices.Clone(tt.nums)); got != tt.want {
				t.Errorf("%s: got %d, want %d", t.Name(), got, tt.want)
			}
		})
	}
}
func TestSubArrayRangesMaximum(t *testing.T) {
	nums := make([]int, 1000)
	for i := range nums {
		if i%2 == 0 {
			nums[i] = -1000000000
		} else {
			nums[i] = 1000000000
		}
	}
	const want int64 = 999000000000000
	if got := subArrayRanges(nums); got != want {
		t.Errorf("%s: got %d, want %d", t.Name(), got, want)
	}
}
