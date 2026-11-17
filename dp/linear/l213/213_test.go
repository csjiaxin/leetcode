package l213

import "testing"

func TestRob(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{name: "example one", nums: []int{2, 3, 2}, want: 3},
		{name: "example two", nums: []int{1, 2, 3, 1}, want: 4},
		{name: "example three", nums: []int{1, 2, 3}, want: 3},
		{name: "single house", nums: []int{5}, want: 5},
		{name: "single empty house", nums: []int{0}, want: 0},
		{name: "two houses prefer first", nums: []int{9, 2}, want: 9},
		{name: "two houses prefer second", nums: []int{2, 9}, want: 9},
		{name: "all zero", nums: []int{0, 0, 0, 0}, want: 0},
		{name: "equal amounts odd count", nums: []int{5, 5, 5, 5, 5}, want: 10},
		{name: "equal amounts even count", nums: []int{5, 5, 5, 5, 5, 5}, want: 15},
		{name: "first and last cannot both be robbed", nums: []int{10, 1, 1, 10}, want: 11},
		{name: "exclude last house", nums: []int{9, 1, 8, 1, 2}, want: 17},
		{name: "exclude first house", nums: []int{2, 1, 8, 1, 9}, want: 17},
		{name: "skip both endpoints", nums: []int{1, 10, 1, 10, 1}, want: 20},
		{name: "skip largest individual house", nums: []int{1, 4, 5, 4, 1}, want: 8},
		{name: "increasing amounts", nums: []int{1, 2, 3, 4, 5, 6}, want: 12},
		{name: "decreasing amounts", nums: []int{6, 5, 4, 3, 2, 1}, want: 12},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := append([]int(nil), tt.nums...)
			if got := rob(input); got != tt.want {
				t.Errorf("rob(%v) = %d, want %d", tt.nums, got, tt.want)
			}
		})
	}
}

func TestRobMaximumInput(t *testing.T) {
	nums := make([]int, 100)
	for i := range nums {
		nums[i] = 1000
	}
	if got := rob(nums); got != 50000 {
		t.Errorf("rob(100 circular houses with 1000 each) = %d, want 50000", got)
	}
}
