package l162

import "testing"

func TestFindPeakElement(t *testing.T) {
	for _, tt := range []struct {
		name string
		nums []int
	}{
		{"example", []int{1, 2, 3, 1}}, {"multiple", []int{1, 2, 1, 3, 5, 6, 4}},
		{"single", []int{1}}, {"increasing", []int{1, 2, 3}},
		{"decreasing", []int{3, 2, 1}}, {"negative", []int{-5, -2, -4}},
	} {
		t.Run(tt.name, func(t *testing.T) {
			i := findPeakElement(tt.nums)
			if i < 0 || i >= len(tt.nums) {
				t.Fatalf("invalid index %d", i)
			}
			if (i > 0 && tt.nums[i] <= tt.nums[i-1]) || (i+1 < len(tt.nums) && tt.nums[i] <= tt.nums[i+1]) {
				t.Errorf("index %d is not a peak in %v", i, tt.nums)
			}
		})
	}
}
