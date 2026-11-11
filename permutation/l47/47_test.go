package l47

import (
	"fmt"
	"reflect"
	"slices"
	"testing"
)

func TestPermuteUnique(t *testing.T) {
	counts := func(rows [][]int) map[string]int {
		result := make(map[string]int)
		for _, row := range rows {
			result[fmt.Sprint(row)]++
		}
		return result
	}
	for _, tt := range []struct {
		name string
		nums []int
		want [][]int
	}{
		{"duplicates", []int{1, 1, 2}, [][]int{{1, 1, 2}, {1, 2, 1}, {2, 1, 1}}},
		{"distinct", []int{1, 2, 3}, [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}}},
		{"single", []int{0}, [][]int{{0}}},
		{"all equal", []int{2, 2, 2}, [][]int{{2, 2, 2}}},
		{"unsorted duplicates", []int{2, 1, 2}, [][]int{{1, 2, 2}, {2, 1, 2}, {2, 2, 1}}},
		{"negative and zero", []int{-1, 0, -1}, [][]int{{-1, -1, 0}, {-1, 0, -1}, {0, -1, -1}}},
	} {
		t.Run(tt.name, func(t *testing.T) {
			got := permuteUnique(slices.Clone(tt.nums))
			// Compare multiplicities as well, so duplicate results are rejected.
			if !reflect.DeepEqual(counts(got), counts(tt.want)) {
				t.Errorf("permuteUnique(%v) = %v, want %v (in any order)", tt.nums, got, tt.want)
			}
		})
	}
}
