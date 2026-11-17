package l46

import (
	"fmt"
	"reflect"
	"slices"
	"testing"
)

func TestPermute(t *testing.T) {
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
		{"example", []int{1, 2, 3}, [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}}},
		{"pair", []int{0, 1}, [][]int{{0, 1}, {1, 0}}},
		{"single", []int{1}, [][]int{{1}}},
		{"negative and zero", []int{-1, 0}, [][]int{{-1, 0}, {0, -1}}},
		{"unsorted", []int{3, 1}, [][]int{{3, 1}, {1, 3}}},
	} {
		t.Run(tt.name, func(t *testing.T) {
			got := permute(slices.Clone(tt.nums))
			// Count multiplicities to reject duplicate answers as well.
			if !reflect.DeepEqual(counts(got), counts(tt.want)) {
				t.Errorf("permute(%v) = %v, want %v (in any order)", tt.nums, got, tt.want)
			}
		})
	}
}
