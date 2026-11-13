package l986

import (
	"slices"
	"testing"
)

func cloneIntervals(in [][]int) [][]int {
	out := make([][]int, len(in))
	for i := range in {
		out[i] = slices.Clone(in[i])
	}
	return out
}
func TestIntervalIntersection(t *testing.T) {
	for _, tt := range []struct {
		name       string
		a, b, want [][]int
	}{
		{"example", [][]int{{0, 2}, {5, 10}, {13, 23}, {24, 25}}, [][]int{{1, 5}, {8, 12}, {15, 24}, {25, 26}}, [][]int{{1, 2}, {5, 5}, {8, 10}, {15, 23}, {24, 24}, {25, 25}}},
		{"empty second", [][]int{{1, 3}}, nil, nil}, {"empty first", nil, [][]int{{1, 3}}, nil},
		{"endpoint", [][]int{{1, 2}}, [][]int{{2, 3}}, [][]int{{2, 2}}},
		{"contained", [][]int{{1, 10}}, [][]int{{2, 3}, {5, 7}}, [][]int{{2, 3}, {5, 7}}},
		{"disjoint", [][]int{{1, 2}}, [][]int{{3, 4}}, nil},
	} {
		t.Run(tt.name, func(t *testing.T) {
			got := intervalIntersection(cloneIntervals(tt.a), cloneIntervals(tt.b))
			if !slices.EqualFunc(got, tt.want, func(a, b []int) bool { return slices.Equal(a, b) }) {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
