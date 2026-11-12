package l1094

import "testing"

func TestCarPooling(t *testing.T) {
	for _, tt := range []struct {
		name     string
		trips    [][]int
		capacity int
		want     bool
	}{
		{"overlap", [][]int{{2, 1, 5}, {3, 3, 7}}, 4, false},
		{"exact capacity", [][]int{{2, 1, 5}, {3, 3, 7}}, 5, true},
		{"shared endpoint", [][]int{{2, 1, 5}, {3, 5, 7}}, 3, true},
		{"unsorted", [][]int{{3, 5, 7}, {2, 1, 5}}, 3, true},
		{"single overload", [][]int{{5, 0, 1}}, 4, false},
		{"simultaneous starts", [][]int{{2, 0, 2}, {2, 0, 3}}, 3, false},
	} {
		t.Run(tt.name, func(t *testing.T) {
			input := make([][]int, len(tt.trips))
			for i := range input {
				input[i] = append([]int(nil), tt.trips[i]...)
			}
			if got := carPooling(input, tt.capacity); got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
