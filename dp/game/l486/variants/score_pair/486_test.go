package l486

import "testing"

func TestPredictTheWinner(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want bool
	}{
		{name: "example zero", nums: []int{2, 4, 55, 6, 8}, want: false},
		{name: "example one", nums: []int{1, 5, 2}, want: false},
		{name: "example two", nums: []int{1, 5, 233, 7}, want: true},
		{name: "single positive value", nums: []int{7}, want: true},
		{name: "single zero is a tie", nums: []int{0}, want: true},
		{name: "two equal values tie", nums: []int{3, 3}, want: true},
		{name: "two values prefer left", nums: []int{9, 2}, want: true},
		{name: "two values prefer right", nums: []int{2, 9}, want: true},
		{name: "all zeros", nums: []int{0, 0, 0, 0, 0}, want: true},
		{name: "equal values odd length", nums: []int{2, 2, 2, 2, 2}, want: true},
		{name: "equal values even length tie", nums: []int{2, 2, 2, 2}, want: true},
		{name: "large inaccessible middle", nums: []int{1, 100, 1}, want: false},
		{name: "zero endpoints", nums: []int{0, 1, 0}, want: false},
		{name: "tie with odd length", nums: []int{1, 2, 1}, want: true},
		{name: "winning left endpoint", nums: []int{10, 1, 2}, want: true},
		{name: "winning right endpoint", nums: []int{2, 1, 10}, want: true},
		{name: "taking largest endpoint is not optimal", nums: []int{8, 15, 3, 7}, want: true},
		{name: "maximum value in middle", nums: []int{0, 10000000, 0}, want: false},
		{name: "maximum value at endpoint", nums: []int{10000000, 0, 0}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nums := append([]int(nil), tt.nums...)
			if got := predictTheWinner(nums); got != tt.want {
				t.Errorf("%s: predictTheWinner(%v) = %t, want %t", t.Name(), tt.nums, got, tt.want)
			}
		})
	}
}

func TestPredictTheWinnerMaximumLength(t *testing.T) {
	tests := []struct {
		name  string
		value int
	}{
		{name: "all zeros tie", value: 0},
		{name: "all ones tie", value: 1},
		{name: "maximum values tie", value: 10000000},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nums := make([]int, 20)
			for i := range nums {
				nums[i] = tt.value
			}
			if got := predictTheWinner(nums); !got {
				t.Errorf("%s: predictTheWinner(20 values of %d) = false, want true", t.Name(), tt.value)
			}
		})
	}
}
