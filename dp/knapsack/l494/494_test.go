package l494

import "testing"

func TestFindTargetSumWays(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   int
	}{
		{name: "example one", nums: []int{1, 1, 1, 1, 1}, target: 3, want: 5},
		{name: "example two", nums: []int{1}, target: 1, want: 1},
		{name: "single negative sign", nums: []int{1}, target: -1, want: 1},
		{name: "single unreachable target", nums: []int{1}, target: 0, want: 0},
		{name: "single zero", nums: []int{0}, target: 0, want: 2},
		{name: "zero cannot reach nonzero", nums: []int{0}, target: 1, want: 0},
		{name: "all zeros", nums: []int{0, 0, 0, 0, 0}, target: 0, want: 32},
		{name: "zeros multiply ways", nums: []int{0, 0, 1}, target: 1, want: 4},
		{name: "zero after nonzero", nums: []int{1, 0}, target: -1, want: 2},
		{name: "target exceeds total", nums: []int{1, 2, 3}, target: 7, want: 0},
		{name: "target below negative total", nums: []int{1, 2, 3}, target: -7, want: 0},
		{name: "parity mismatch", nums: []int{1, 2, 3}, target: 1, want: 0},
		{name: "zero target", nums: []int{1, 2, 3}, target: 0, want: 2},
		{name: "negative target", nums: []int{1, 1, 1, 1, 1}, target: -3, want: 5},
		{name: "all positive signs", nums: []int{1, 2, 3}, target: 6, want: 1},
		{name: "all negative signs", nums: []int{1, 2, 3}, target: -6, want: 1},
		{name: "equal values at distinct indices", nums: []int{2, 2, 2, 2}, target: 0, want: 6},
		{name: "cannot reuse an element", nums: []int{1, 3}, target: 0, want: 0},
		{name: "maximum total", nums: []int{1000}, target: -1000, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nums := append([]int(nil), tt.nums...)
			if got := findTargetSumWays(nums, tt.target); got != tt.want {
				t.Errorf("findTargetSumWays(%v, %d) = %d, want %d", tt.nums, tt.target, got, tt.want)
			}
		})
	}
}

func TestFindTargetSumWaysMaximumLength(t *testing.T) {
	tests := []struct {
		name   string
		value  int
		target int
		want   int
	}{
		{name: "twenty zeros", value: 0, target: 0, want: 1 << 20},
		{name: "twenty ones", value: 1, target: 0, want: 184756},
		{name: "maximum length and total", value: 50, target: 0, want: 184756},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nums := make([]int, 20)
			for i := range nums {
				nums[i] = tt.value
			}
			if got := findTargetSumWays(nums, tt.target); got != tt.want {
				t.Errorf("findTargetSumWays(20 values of %d, %d) = %d, want %d", tt.value, tt.target, got, tt.want)
			}
		})
	}
}
