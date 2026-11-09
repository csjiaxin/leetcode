package l46

import (
	"fmt"
	"reflect"
	"testing"
)

func TestPermute(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want [][]int
	}{
		{
			name: "example one",
			nums: []int{1, 2, 3},
			want: [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}},
		},
		{
			name: "example two",
			nums: []int{0, 1},
			want: [][]int{{0, 1}, {1, 0}},
		},
		{
			name: "example three",
			nums: []int{1},
			want: [][]int{{1}},
		},
		{
			name: "negative and zero values",
			nums: []int{-1, 0, 2},
			want: [][]int{{-1, 0, 2}, {-1, 2, 0}, {0, -1, 2}, {0, 2, -1}, {2, -1, 0}, {2, 0, -1}},
		},
		{
			name: "unsorted input",
			nums: []int{3, 1, 2},
			want: [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := permute(tt.nums)
			// LeetCode accepts permutations in any order; compare multiplicities.
			counts := func(permutations [][]int) map[string]int {
				result := make(map[string]int)
				for _, permutation := range permutations {
					result[fmt.Sprint(permutation)]++
				}
				return result
			}
			if !reflect.DeepEqual(counts(got), counts(tt.want)) {
				t.Errorf("permute(%v) = %v, want %v (in any order)", tt.nums, got, tt.want)
			}
		})
	}
}

func TestPermuteIndependentResults(t *testing.T) {
	got := permute([]int{1, 2, 3})
	if len(got) != 6 {
		t.Fatalf("got %d permutations, want 6", len(got))
	}

	before := make([][]int, len(got))
	for i, permutation := range got {
		if len(permutation) != 3 {
			t.Fatalf("permutation %d has length %d, want 3", i, len(permutation))
		}
		before[i] = append([]int(nil), permutation...)
	}

	got[0][0] = 99
	for i := 1; i < len(got); i++ {
		if !reflect.DeepEqual(got[i], before[i]) {
			t.Errorf("mutating permutation 0 changed permutation %d: got %v, want %v", i, got[i], before[i])
		}
	}
}
