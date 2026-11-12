package l2089

import (
	"slices"
	"testing"
)

func TestTargetIndices(t *testing.T) {
	for _, tt := range []struct {
		name   string
		nums   []int
		target int
		want   []int
	}{
		{"example", []int{1, 2, 5, 2, 3}, 2, []int{1, 2}}, {"single", []int{1}, 1, []int{0}},
		{"missing", []int{1, 2, 3}, 4, []int{}}, {"below", []int{2, 3}, 1, []int{}},
		{"all equal", []int{2, 2, 2}, 2, []int{0, 1, 2}}, {"empty", nil, 1, []int{}},
	} {
		t.Run(tt.name, func(t *testing.T) {
			if got := targetIndices(slices.Clone(tt.nums), tt.target); !slices.Equal(got, tt.want) {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
