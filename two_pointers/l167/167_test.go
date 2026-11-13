package l167

import (
	"slices"
	"testing"
)

func TestTwoSum(t *testing.T) {
	for _, tt := range []struct {
		name    string
		numbers []int
		target  int
		want    []int
	}{
		{"example", []int{2, 7, 11, 15}, 9, []int{1, 2}},
		{"endpoints", []int{2, 3, 4}, 6, []int{1, 3}},
		{"negative", []int{-1, 0}, -1, []int{1, 2}},
		{"duplicates", []int{1, 1, 3}, 2, []int{1, 2}},
		{"right pair", []int{1, 2, 4, 8}, 12, []int{3, 4}},
	} {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoSum(tt.numbers, tt.target); !slices.Equal(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
