package l198

import "testing"

func TestRob(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "example one",
			nums: []int{1, 2, 3, 1},
			want: 4,
		},
		{
			name: "example two",
			nums: []int{2, 7, 9, 3, 1},
			want: 12,
		},
		{
			name: "single house",
			nums: []int{5},
			want: 5,
		},
		{
			name: "single empty house",
			nums: []int{0},
			want: 0,
		},
		{
			name: "two houses prefer first",
			nums: []int{9, 2},
			want: 9,
		},
		{
			name: "two houses prefer second",
			nums: []int{2, 9},
			want: 9,
		},
		{
			name: "all zero",
			nums: []int{0, 0, 0, 0},
			want: 0,
		},
		{
			name: "equal amounts",
			nums: []int{5, 5, 5, 5, 5},
			want: 15,
		},
		{
			name: "end houses are not adjacent",
			nums: []int{2, 1, 1, 2},
			want: 4,
		},
		{
			name: "skip largest individual house",
			nums: []int{4, 5, 4},
			want: 8,
		},
		{
			name: "optimal houses need not share index parity",
			nums: []int{10, 1, 1, 10},
			want: 20,
		},
		{
			name: "increasing amounts",
			nums: []int{1, 2, 3, 4, 5, 6},
			want: 12,
		},
		{
			name: "decreasing amounts",
			nums: []int{6, 5, 4, 3, 2, 1},
			want: 12,
		},
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
		nums[i] = 400
	}
	if got := rob(nums); got != 20000 {
		t.Errorf("rob(100 houses with 400 each) = %d, want 20000", got)
	}
}
