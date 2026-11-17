package l1109

import (
	"slices"
	"testing"
)

func TestCorpFlightBookings(t *testing.T) {
	for _, tt := range []struct {
		name     string
		bookings [][]int
		n        int
		want     []int
	}{
		{"example", [][]int{{1, 2, 10}, {2, 3, 20}, {2, 5, 25}}, 5, []int{10, 55, 45, 25, 25}},
		{"single", [][]int{{1, 1, 5}}, 1, []int{5}},
		{"whole range", [][]int{{1, 3, 7}}, 3, []int{7, 7, 7}},
		{"last flight", [][]int{{3, 3, 4}}, 3, []int{0, 0, 4}},
		{"shared inclusive endpoint", [][]int{{1, 2, 2}, {2, 3, 3}}, 3, []int{2, 5, 3}},
	} {
		t.Run(tt.name, func(t *testing.T) {
			input := make([][]int, len(tt.bookings))
			for i := range input {
				input[i] = append([]int(nil), tt.bookings[i]...)
			}
			if got := corpFlightBookings(input, tt.n); !slices.Equal(got, tt.want) {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
