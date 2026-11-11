package l373

import (
	"slices"
	"testing"
)

func TestKSmallestPairs(t *testing.T) {
	for _, tt := range []struct {
		name string
		a, b []int
		k    int
		want [][]int
	}{
		{"example", []int{1, 7, 11}, []int{2, 4, 6}, 3, [][]int{{1, 2}, {1, 4}, {1, 6}}}, {"duplicates", []int{1, 1, 2}, []int{1, 2, 3}, 2, [][]int{{1, 1}, {1, 1}}}, {"k exceeds product", []int{1, 2}, []int{3}, 5, [][]int{{1, 3}, {2, 3}}}, {"negative", []int{-3, -1}, []int{-2, 4}, 3, [][]int{{-3, -2}, {-1, -2}, {-3, 4}}},
	} {
		t.Run(tt.name, func(t *testing.T) {
			got := kSmallestPairs(tt.a, tt.b, tt.k)
			if len(got) != len(tt.want) {
				t.Fatalf("%s: length %d want %d", t.Name(), len(got), len(tt.want))
			}
			for i := range got {
				if !slices.Equal(got[i], tt.want[i]) {
					t.Errorf("%s: pair %d got %v want %v", t.Name(), i, got[i], tt.want[i])
				}
			}
		})
	}
}
