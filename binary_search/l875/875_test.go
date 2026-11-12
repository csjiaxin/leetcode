package l875

import "testing"

func TestMinEatingSpeed(t *testing.T) {
	for _, tt := range []struct {
		name    string
		piles   []int
		h, want int
	}{
		{"example", []int{3, 6, 7, 11}, 8, 4}, {"one per pile", []int{30, 11, 23, 4, 20}, 5, 30},
		{"rounding", []int{30, 11, 23, 4, 20}, 6, 23}, {"single", []int{9}, 2, 5},
		{"speed one", []int{2, 3}, 5, 1}, {"large", []int{1000000000, 1000000000, 1000000000}, 1000000000, 4},
	} {
		t.Run(tt.name, func(t *testing.T) {
			if got := minEatingSpeed(tt.piles, tt.h); got != tt.want {
				t.Errorf("got %d want %d", got, tt.want)
			}
		})
	}
}
