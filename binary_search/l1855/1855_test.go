package l1855

import "testing"

func TestMaxDistance(t *testing.T) {
	for _, tt := range []struct {
		name string
		a, b []int
		want int
	}{
		{"example", []int{55, 30, 5, 4, 2}, []int{100, 20, 10, 10, 5}, 2},
		{"duplicates", []int{2, 2, 2}, []int{10, 10, 1}, 1},
		{"example three", []int{30, 29, 19, 5}, []int{25, 25, 25, 25, 25}, 2},
		{"none", []int{5, 4}, []int{3, 2}, 0}, {"single equal", []int{2}, []int{2}, 0},
		{"all equal", []int{2, 2}, []int{2, 2, 2, 2}, 3}, {"short second", []int{5, 4, 3}, []int{5}, 0},
	} {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxDistance(tt.a, tt.b); got != tt.want {
				t.Errorf("got %d want %d", got, tt.want)
			}
		})
	}
}
