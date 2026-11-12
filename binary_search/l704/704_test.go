package l704

import "testing"

func TestSearch(t *testing.T) {
	for _, tt := range []struct {
		name         string
		nums         []int
		target, want int
	}{
		{"example", []int{-1, 0, 3, 5, 9, 12}, 9, 4},
		{"missing", []int{-1, 0, 3, 5, 9, 12}, 2, -1},
		{"empty", nil, 1, -1}, {"single", []int{5}, 5, 0},
		{"below", []int{1, 3}, 0, -1}, {"above", []int{1, 3}, 4, -1},
		{"first", []int{1, 3}, 1, 0}, {"last", []int{1, 3}, 3, 1},
	} {
		t.Run(tt.name, func(t *testing.T) {
			if got := search(tt.nums, tt.target); got != tt.want {
				t.Errorf("got %d want %d", got, tt.want)
			}
		})
	}
}
