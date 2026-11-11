package l2530

import "testing"

func TestMaxKelements(t *testing.T) {
	for _, tt := range []struct {
		name string
		nums []int
		k    int
		want int64
	}{
		{"example one", []int{10, 10, 10, 10, 10}, 5, 50}, {"example two", []int{1, 10, 3, 3, 3}, 3, 17}, {"round up", []int{4}, 2, 6}, {"ones", []int{1}, 5, 5}, {"large score", []int{1000000000, 1000000000, 1000000000}, 3, 3000000000},
	} {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxKelements(tt.nums, tt.k); got != tt.want {
				t.Errorf("%s: got %d want %d", t.Name(), got, tt.want)
			}
		})
	}
}
