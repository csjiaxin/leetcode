package l1608

import (
	"slices"
	"testing"
)

func TestSpecialArray(t *testing.T) {
	for _, tt := range []struct {
		name string
		nums []int
		want int
	}{
		{"example", []int{3, 5}, 2}, {"zeros", []int{0, 0}, -1},
		{"mixed", []int{0, 4, 3, 0, 4}, 3}, {"missing", []int{3, 6, 7, 7, 0}, -1},
		{"single positive", []int{1}, 1}, {"single zero", []int{0}, -1},
		{"duplicates", []int{2, 2, 2}, -1}, {"all qualify", []int{5, 5, 5}, 3},
	} {
		t.Run(tt.name, func(t *testing.T) {
			if got := specialArray(slices.Clone(tt.nums)); got != tt.want {
				t.Errorf("got %d want %d", got, tt.want)
			}
		})
	}
}
