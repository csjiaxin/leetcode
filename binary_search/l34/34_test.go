package l34

import (
	"slices"
	"testing"
)

func TestSearchRange(t *testing.T) {
	for _, tt := range []struct {
		name   string
		nums   []int
		target int
		want   []int
	}{
		{"example", []int{5, 7, 7, 8, 8, 10}, 8, []int{3, 4}},
		{"missing", []int{5, 7, 7, 8, 8, 10}, 6, []int{-1, -1}},
		{"empty", nil, 0, []int{-1, -1}}, {"single", []int{1}, 1, []int{0, 0}},
		{"all equal", []int{2, 2, 2}, 2, []int{0, 2}},
		{"below", []int{2, 2}, 1, []int{-1, -1}}, {"above", []int{2, 2}, 3, []int{-1, -1}},
	} {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchRange(tt.nums, tt.target); !slices.Equal(got, tt.want) {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
