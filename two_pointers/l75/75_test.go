package l75

import (
	"slices"
	"testing"
)

func TestSortColors(t *testing.T) {
	for _, tt := range []struct {
		name       string
		nums, want []int
	}{
		{"example", []int{2, 0, 2, 1, 1, 0}, []int{0, 0, 1, 1, 2, 2}},
		{"recheck swapped", []int{2, 0, 1}, []int{0, 1, 2}}, {"single", []int{0}, []int{0}},
		{"equal", []int{2, 2, 2}, []int{2, 2, 2}}, {"sorted", []int{0, 1, 2}, []int{0, 1, 2}},
		{"reverse", []int{2, 2, 1, 1, 0, 0}, []int{0, 0, 1, 1, 2, 2}},
	} {
		t.Run(tt.name, func(t *testing.T) {
			got := slices.Clone(tt.nums)
			sortColors(got)
			if !slices.Equal(got, tt.want) {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
