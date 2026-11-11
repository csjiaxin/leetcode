package l907

import (
	"slices"
	"testing"
)

func TestSumSubarrayMins(t *testing.T) {
	for _, tt := range []struct {
		name string
		arr  []int
		want int
	}{
		{"example one", []int{3, 1, 2, 4}, 17}, {"example two", []int{11, 81, 94, 43, 3}, 444},
		{"single", []int{5}, 5}, {"equal pair", []int{2, 2}, 6}, {"equal triple", []int{2, 2, 2}, 12},
		{"increasing", []int{1, 2, 3}, 10}, {"decreasing", []int{3, 2, 1}, 10}, {"separated equal minima", []int{1, 3, 1}, 8},
	} {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumSubarrayMins(slices.Clone(tt.arr)); got != tt.want {
				t.Errorf("%s: got %d, want %d", t.Name(), got, tt.want)
			}
		})
	}
}
func TestSumSubarrayMinsMaximum(t *testing.T) {
	const n = 30000
	arr := make([]int, n)
	for i := range arr {
		arr[i] = 30000
	}
	want := int(int64(n) * int64(n+1) / 2 * 30000 % 1000000007)
	if got := sumSubarrayMins(arr); got != want {
		t.Errorf("%s: got %d, want %d", t.Name(), got, want)
	}
}
