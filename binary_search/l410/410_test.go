package l410

import "testing"

func TestSplitArray(t *testing.T) {
	for _, tt := range []struct {
		name    string
		nums    []int
		k, want int
	}{
		{"example", []int{7, 2, 5, 10, 8}, 2, 18},
		{"example two", []int{1, 2, 3, 4, 5}, 2, 9},
		{"one group", []int{1, 2, 3}, 1, 6}, {"each separate", []int{1, 4, 4}, 3, 4},
		{"single", []int{9}, 1, 9}, {"zeros", []int{0, 0, 0}, 2, 0},
		{"zero groups", []int{0, 5, 0}, 3, 5},
	} {
		t.Run(tt.name, func(t *testing.T) {
			if got := splitArray(tt.nums, tt.k); got != tt.want {
				t.Errorf("got %d want %d", got, tt.want)
			}
		})
	}
}
