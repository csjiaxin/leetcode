package l55

import "testing"

func TestCanJump(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want bool
	}{
		{name: "example one", nums: []int{2, 3, 1, 1, 4}, want: true},
		{name: "example two", nums: []int{3, 2, 1, 0, 4}, want: false},
		{name: "single zero", nums: []int{0}, want: true},
		{name: "single positive value", nums: []int{5}, want: true},
		{name: "cannot leave start", nums: []int{0, 1}, want: false},
		{name: "reach zero at last index", nums: []int{1, 0}, want: true},
		{name: "all unit jumps", nums: []int{1, 1, 1, 1}, want: true},
		{name: "zero blocks path", nums: []int{1, 0, 1}, want: false},
		{name: "jump over zero", nums: []int{2, 0, 0}, want: true},
		{name: "exact reach across zeros", nums: []int{3, 0, 0, 0}, want: true},
		{name: "jump can exceed last index", nums: []int{10, 0, 0}, want: true},
		{name: "unreachable large jump", nums: []int{1, 0, 100, 0}, want: false},
		{name: "choose shorter first jump", nums: []int{2, 3, 0, 0, 0}, want: true},
		{name: "extend reach before dead end", nums: []int{2, 0, 2, 0, 0}, want: true},
		{name: "reach stops before last index", nums: []int{2, 0, 1, 0, 1}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canJump(tt.nums); got != tt.want {
				t.Errorf("canJump(%v) = %t, want %t", tt.nums, got, tt.want)
			}
		})
	}
}

func TestCanJumpLargeInput(t *testing.T) {
	for _, reachable := range []bool{true, false} {
		name := "reachable"
		if !reachable {
			name = "blocked near end"
		}
		t.Run(name, func(t *testing.T) {
			nums := make([]int, 10000)
			for i := range nums {
				nums[i] = 1
			}
			if !reachable {
				nums[len(nums)-2] = 0
			}
			if got := canJump(nums); got != reachable {
				t.Errorf("canJump(10000 elements, %s) = %t, want %t", name, got, reachable)
			}
		})
	}
}
