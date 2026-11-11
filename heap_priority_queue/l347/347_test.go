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
		{"example", []int{1, 1, 1, 2, 2, 3}, 2, []int{1, 2}}, {"single", []int{1}, 1, []int{1}}, {"negative", []int{-1, -1, 0, 2, 2, 2}, 2, []int{-1, 2}}, {"all unique included", []int{3, 2, 1}, 3, []int{1, 2, 3}},
	} {
		t.Run(tt.name, func(t *testing.T) {
			got := topKFrequent(tt.nums, tt.k)
			slices.Sort(got)
			slices.Sort(tt.want)
			if !slices.Equal(got, tt.want) {
				t.Errorf("%s: got %v want %v", t.Name(), got, tt.want)
			}
		})
	}
}
