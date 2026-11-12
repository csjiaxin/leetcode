package l560

import "testing"

func TestSubarraySum(t *testing.T) {
	for _, tt := range []struct {
		name    string
		nums    []int
		k, want int
	}{
		{"example", []int{1, 1, 1}, 2, 2}, {"example two", []int{1, 2, 3}, 3, 2},
		{"negative", []int{1, -1, 1}, 1, 3}, {"zeros", []int{0, 0, 0}, 0, 6},
		{"single", []int{3}, 3, 1}, {"missing", []int{1, 2}, 4, 0},
	} {
		t.Run(tt.name, func(t *testing.T) {
			if got := subarraySum(append([]int(nil), tt.nums...), tt.k); got != tt.want {
				t.Errorf("got %d want %d", got, tt.want)
			}
		})
	}
}
