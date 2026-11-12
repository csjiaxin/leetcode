package l1248

import "testing"

func TestNumberOfSubarrays(t *testing.T) {
	for _, tt := range []struct {
		name    string
		nums    []int
		k, want int
	}{
		{"example", []int{1, 1, 2, 1, 1}, 3, 2}, {"no odds", []int{2, 4, 6}, 1, 0},
		{"even extensions", []int{2, 2, 2, 1, 2, 2, 1, 2, 2, 2}, 2, 16},
		{"single", []int{1}, 1, 1}, {"all odd", []int{1, 3, 5}, 1, 3},
		{"one odd extensions", []int{2, 1, 2}, 1, 4},
	} {
		t.Run(tt.name, func(t *testing.T) {
			if got := numberOfSubarrays(append([]int(nil), tt.nums...), tt.k); got != tt.want {
				t.Errorf("got %d want %d", got, tt.want)
			}
		})
	}
}
