package l347

import (
	"slices"
	"testing"
)

func TestTopKFrequent(t *testing.T) {
	for _, tt := range []struct {
		name string
		nums []int
		k    int
		want []int
	}{
		{"example", []int{1, 1, 1, 2, 2, 3}, 2, []int{1, 2}},
		{"single", []int{1}, 1, []int{1}},
		{"frequency not value", []int{100, 1, 1, 1, 2, 2}, 1, []int{1}},
		{"negative", []int{-1, -1, 0, 0, 0, 2}, 2, []int{-1, 0}},
		{"all distinct included", []int{3, 1, 2}, 3, []int{1, 2, 3}},
		{"duplicates", []int{7, 7, 7}, 1, []int{7}},
	} {
		t.Run(tt.name, func(t *testing.T) {
			got := topKFrequent(slices.Clone(tt.nums), tt.k)
			slices.Sort(got)
			want := slices.Clone(tt.want)
			slices.Sort(want)
			if !slices.Equal(got, want) {
				t.Errorf("got %v want set %v", got, want)
			}
		})
	}
}
