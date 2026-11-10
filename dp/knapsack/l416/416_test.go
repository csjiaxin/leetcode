package l416

import "testing"

func TestCanPartition(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want bool
	}{
		{name: "example one", nums: []int{1, 5, 11, 5}, want: true},
		{name: "example two", nums: []int{1, 2, 3, 5}, want: false},
		{name: "single element", nums: []int{1}, want: false},
		{name: "single even element", nums: []int{2}, want: false},
		{name: "two equal elements", nums: []int{3, 3}, want: true},
		{name: "two unequal elements even total", nums: []int{1, 3}, want: false},
		{name: "even total without partition", nums: []int{2, 2, 2, 2, 2}, want: false},
		{name: "cannot reuse an element", nums: []int{1, 2, 5}, want: false},
		{name: "element exceeds half total", nums: []int{1, 2, 7}, want: false},
		{name: "different subset sizes", nums: []int{1, 2, 3, 6}, want: true},
		{name: "combination required", nums: []int{2, 3, 4, 5}, want: true},
		{name: "duplicate values", nums: []int{2, 2, 3, 3}, want: true},
		{name: "maximum element values", nums: []int{100, 100}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nums := append([]int(nil), tt.nums...)
			if got := canPartition(nums); got != tt.want {
				t.Errorf("canPartition(%v) = %t, want %t", tt.nums, got, tt.want)
			}
		})
	}
}

func TestCanPartitionLargeInput(t *testing.T) {
	tests := []struct {
		name  string
		count int
		want  bool
	}{
		{name: "maximum length and total", count: 200, want: true},
		{name: "even total but no partition", count: 199, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nums := make([]int, tt.count)
			for i := range nums {
				nums[i] = 100
			}
			if got := canPartition(nums); got != tt.want {
				t.Errorf("canPartition(%d values of 100) = %t, want %t", tt.count, got, tt.want)
			}
		})
	}
}
