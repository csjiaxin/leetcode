package l1760

import "testing"

func TestMinimumSize(t *testing.T) {
	for _, tt := range []struct {
		name             string
		nums             []int
		operations, want int
	}{
		{"example", []int{9}, 2, 3}, {"multiple", []int{2, 4, 8, 2}, 4, 2},
		{"no operations", []int{2, 7, 3}, 0, 7}, {"one", []int{1}, 0, 1},
		{"enough operations", []int{2, 3}, 3, 1}, {"rounding", []int{10}, 2, 4},
		{"large", []int{1000000000}, 999999999, 1},
	} {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumSize(tt.nums, tt.operations); got != tt.want {
				t.Errorf("got %d want %d", got, tt.want)
			}
		})
	}
}
