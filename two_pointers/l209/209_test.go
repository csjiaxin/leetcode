package l209

import (
	"slices"
	"testing"
)

func TestMinSubArrayLen(t *testing.T) {
	for _, tt := range []struct {
		name   string
		target int
		nums   []int
		want   int
	}{
		{"example", 7, []int{2, 3, 1, 2, 4, 3}, 2}, {"single match", 4, []int{1, 4, 4}, 1},
		{"impossible", 11, []int{1, 1, 1, 1, 1, 1, 1, 1}, 0}, {"whole", 6, []int{1, 2, 3}, 3},
		{"shrink repeatedly", 5, []int{1, 1, 1, 5}, 1}, {"single too small", 2, []int{1}, 0},
	} {
		t.Run(tt.name, func(t *testing.T) {
			if got := minSubArrayLen(tt.target, slices.Clone(tt.nums)); got != tt.want {
				t.Errorf("got %d want %d", got, tt.want)
			}
		})
	}
}
