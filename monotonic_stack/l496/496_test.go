package l496

import (
	"slices"
	"testing"
)

func TestNextGreaterElement(t *testing.T) {
	tests := []struct {
		name              string
		query, nums, want []int
	}{
		{"example one", []int{4, 1, 2}, []int{1, 3, 4, 2}, []int{-1, 3, -1}},
		{"example two", []int{2, 4}, []int{1, 2, 3, 4}, []int{3, -1}},
		{"single", []int{0}, []int{0}, []int{-1}},
		{"decreasing", []int{3, 2, 1}, []int{3, 2, 1}, []int{-1, -1, -1}},
		{"nearest not largest", []int{2, 1}, []int{2, 3, 1, 4}, []int{3, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nextGreaterElement(slices.Clone(tt.query), slices.Clone(tt.nums)); !slices.Equal(got, tt.want) {
				t.Errorf("%s: got %v, want %v", t.Name(), got, tt.want)
			}
		})
	}
}
