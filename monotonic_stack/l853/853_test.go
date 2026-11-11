package l853

import (
	"slices"
	"testing"
)

func TestCarFleet(t *testing.T) {
	for _, tt := range []struct {
		name            string
		target          int
		position, speed []int
		want            int
	}{
		{"example one", 12, []int{10, 8, 0, 5, 3}, []int{2, 4, 1, 1, 3}, 3},
		{"single", 10, []int{3}, []int{3}, 1},
		{"all merge", 100, []int{0, 2, 4}, []int{4, 2, 1}, 1},
		{"meet exactly at target", 10, []int{0, 5}, []int{2, 1}, 1},
		{"no merge", 10, []int{0, 5}, []int{1, 1}, 2},
		{"unsorted", 10, []int{5, 0, 8}, []int{1, 1, 1}, 3},
		{"fractional time not integer", 10, []int{6, 5}, []int{3, 3}, 2},
		{"merge chain", 10, []int{0, 4, 6}, []int{5, 3, 2}, 1},
	} {
		t.Run(tt.name, func(t *testing.T) {
			if got := carFleet(tt.target, slices.Clone(tt.position), slices.Clone(tt.speed)); got != tt.want {
				t.Errorf("%s: got %d, want %d", t.Name(), got, tt.want)
			}
		})
	}
}
