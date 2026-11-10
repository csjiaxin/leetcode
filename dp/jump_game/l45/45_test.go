package l45

import "testing"

func TestJump(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{name: "example one", nums: []int{2, 3, 1, 1, 4}, want: 2},
		{name: "example two", nums: []int{2, 3, 0, 1, 4}, want: 2},
		{name: "single zero", nums: []int{0}, want: 0},
		{name: "single positive value", nums: []int{5}, want: 0},
		{name: "two elements", nums: []int{1, 0}, want: 1},
		{name: "all unit jumps", nums: []int{1, 1, 1, 1, 1}, want: 4},
		{name: "direct jump to last index", nums: []int{4, 0, 0, 0, 0}, want: 1},
		{name: "jump range exceeds last index", nums: []int{10, 0, 0}, want: 1},
		{name: "shorter first jump is optimal", nums: []int{2, 3, 1, 1, 1}, want: 2},
		{name: "jump over zeros", nums: []int{2, 0, 2, 0, 1}, want: 2},
		{name: "extend reach within current layer", nums: []int{3, 4, 1, 1, 1, 0}, want: 2},
		{name: "multiple layers", nums: []int{2, 1, 2, 1, 2, 1, 0}, want: 3},
		{name: "do not jump again at destination", nums: []int{1, 2, 1, 0}, want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := jump(tt.nums); got != tt.want {
				t.Errorf("jump(%v) = %d, want %d", tt.nums, got, tt.want)
			}
		})
	}
}

func TestJumpLargeInput(t *testing.T) {
	tests := []struct {
		name  string
		value int
		want  int
	}{
		{name: "unit jumps", value: 1, want: 9999},
		{name: "maximum jump length", value: 1000, want: 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nums := make([]int, 10000)
			for i := range nums {
				nums[i] = tt.value
			}
			if got := jump(nums); got != tt.want {
				t.Errorf("jump(10000 values of %d) = %d, want %d", tt.value, got, tt.want)
			}
		})
	}
}
