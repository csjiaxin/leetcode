package l907

import (
	"slices"
	"testing"
)

func TestSumSubarrayMins(t *testing.T) {
	large := make([]int, 1000)
	for i := range large {
		large[i] = 30000
	}
	for _, tt := range []struct {
		name   string
		values []int
		want   int
	}{
		{"example", []int{3, 1, 2, 4}, 17}, {"example two", []int{11, 81, 94, 43, 87}, 444},
		{"single", []int{5}, 5}, {"equal ownership", []int{2, 2}, 6},
		{"ascending", []int{1, 2, 3}, 10}, {"descending", []int{3, 2, 1}, 10},
		{"modulo", large, int(int64(30000) * 1000 * 1001 / 2 % 1000000007)},
	} {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumSubarrayMins(slices.Clone(tt.values)); got != tt.want {
				t.Errorf("got %d want %d", got, tt.want)
			}
		})
	}
}
