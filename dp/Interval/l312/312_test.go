package l312

import "testing"

func TestMaxCoins(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"example one", []int{3, 1, 5, 8}, 167},
		{"example two", []int{1, 5}, 10},
		{"single balloon", []int{5}, 5},
		{"zero balloon", []int{0}, 0},
		{"two balloons", []int{3, 5}, 20},
		{"all ones", []int{1, 1, 1}, 3},
		{"zero between balloons", []int{3, 0, 5}, 20},
		{"maximum values", []int{100, 100, 100}, 1010100},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxCoins(tt.nums); got != tt.want {
				t.Errorf("%s: got %d, want %d", t.Name(), got, tt.want)
			}
		})
	}
}

func TestMaxCoinsMaximumLength(t *testing.T) {
	nums := make([]int, 300)
	for i := range nums {
		nums[i] = 100
	}
	if got := maxCoins(nums); got != 298010100 {
		t.Errorf("%s: got %d, want 298010100", t.Name(), got)
	}
}
