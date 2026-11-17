package l239

import (
	"slices"
	"testing"
)

func TestMaxSlidingWindow(t *testing.T) {
	for _, tt := range []struct {
		name string
		nums []int
		k    int
		want []int
	}{
		{"example", []int{1, 3, -1, -3, 5, 3, 6, 7}, 3, []int{3, 3, 5, 5, 6, 7}},
		{"single", []int{1}, 1, []int{1}},
		{"window one", []int{3, -1, 2}, 1, []int{3, -1, 2}},
		{"full window", []int{2, 5, 1}, 3, []int{5}},
		{"equal", []int{2, 2, 2}, 2, []int{2, 2}},
		{"equal negative maximum retained", []int{-1, -1, -2, -3}, 3, []int{-1, -1}},
		{"equal negative maximum expires", []int{-1, -1, -2, -3}, 2, []int{-1, -1, -2}},
		{"expire maximum", []int{5, 4, 3, 2}, 2, []int{5, 4, 3}},
		{"new maximum", []int{1, 2, 3, 4}, 2, []int{2, 3, 4}},
		{"negative", []int{-4, -2, -5, -1}, 2, []int{-2, -2, -1}},
	} {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSlidingWindow(slices.Clone(tt.nums), tt.k); !slices.Equal(got, tt.want) {
				t.Errorf("%s: got %v, want %v", t.Name(), got, tt.want)
			}
		})
	}
}
func TestMaxSlidingWindowMaximum(t *testing.T) {
	nums := make([]int, 100000)
	for i := range nums {
		nums[i] = i % 10001
	}
	got := maxSlidingWindow(nums, len(nums))
	if !slices.Equal(got, []int{10000}) {
		t.Errorf("%s: got %v, want [10000]", t.Name(), got)
	}
}
