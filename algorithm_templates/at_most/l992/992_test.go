package l992

import "testing"

func TestSubarraysWithKDistinct(t *testing.T) {
	for _, tt := range []struct {
		name    string
		nums    []int
		k, want int
	}{
		{"example", []int{1, 2, 1, 2, 3}, 2, 7}, {"example two", []int{1, 2, 1, 3, 4}, 3, 3},
		{"duplicates", []int{1, 1, 1}, 1, 6}, {"single", []int{1}, 1, 1},
		{"missing", []int{1, 1}, 2, 0}, {"all distinct", []int{1, 2, 3}, 2, 2},
	} {
		t.Run(tt.name, func(t *testing.T) {
			if got := subarraysWithKDistinct(append([]int(nil), tt.nums...), tt.k); got != tt.want {
				t.Errorf("got %d want %d", got, tt.want)
			}
		})
	}
}
