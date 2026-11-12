package l1011

import "testing"

func TestShipWithinDays(t *testing.T) {
	for _, tt := range []struct {
		name       string
		weights    []int
		days, want int
	}{
		{"example", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5, 15},
		{"example two", []int{3, 2, 2, 4, 1, 4}, 3, 6}, {"example three", []int{1, 2, 3, 1, 1}, 4, 3},
		{"one day", []int{1, 2, 3}, 1, 6}, {"one per day", []int{1, 2, 3}, 3, 3},
		{"single", []int{7}, 1, 7}, {"spare days", []int{2, 2}, 5, 2},
	} {
		t.Run(tt.name, func(t *testing.T) {
			if got := shipWithinDays(tt.weights, tt.days); got != tt.want {
				t.Errorf("got %d want %d", got, tt.want)
			}
		})
	}
}
