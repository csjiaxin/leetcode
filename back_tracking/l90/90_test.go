package l90

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

func TestSubsetsWithDup(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want [][]int
	}{
		{
			name: "example one",
			nums: []int{1, 2, 2},
			want: [][]int{{}, {1}, {2}, {1, 2}, {2, 2}, {1, 2, 2}},
		},
		{
			name: "example two",
			nums: []int{0},
			want: [][]int{{}, {0}},
		},
		{
			name: "all equal",
			nums: []int{2, 2, 2},
			want: [][]int{{}, {2}, {2, 2}, {2, 2, 2}},
		},
		{
			name: "distinct elements",
			nums: []int{1, 2, 3},
			want: [][]int{{}, {1}, {2}, {3}, {1, 2}, {1, 3}, {2, 3}, {1, 2, 3}},
		},
		{
			name: "unsorted duplicates",
			nums: []int{2, 1, 2},
			want: [][]int{{}, {1}, {2}, {1, 2}, {2, 2}, {1, 2, 2}},
		},
		{
			name: "negative values and zero",
			nums: []int{-1, 0, -1},
			want: [][]int{{}, {-1}, {-1, -1}, {0}, {-1, 0}, {-1, -1, 0}},
		},
		{
			name: "multiple duplicate groups",
			nums: []int{1, 1, 2, 2},
			want: [][]int{{}, {1}, {1, 1}, {2}, {1, 2}, {1, 1, 2}, {2, 2}, {1, 2, 2}, {1, 1, 2, 2}},
		},
	}

	// Ignore ordering within and between subsets, but retain multiplicities
	// so duplicate subsets still cause a test failure.
	counts := func(subsets [][]int) map[string]int {
		result := make(map[string]int)
		for _, subset := range subsets {
			canonical := append([]int(nil), subset...)
			sort.Ints(canonical)
			result[fmt.Sprint(canonical)]++
		}
		return result
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := append([]int(nil), tt.nums...)
			got := subsetsWithDup(input)
			if !reflect.DeepEqual(counts(got), counts(tt.want)) {
				t.Errorf("subsetsWithDup(%v) = %v, want %v (in any order)", tt.nums, got, tt.want)
			}
		})
	}
}

func TestSubsetsWithDupIndependentResults(t *testing.T) {
	got := subsetsWithDup([]int{1, 2, 2})
	if len(got) != 6 {
		t.Fatalf("got %d subsets, want 6", len(got))
	}

	before := make([][]int, len(got))
	for i, subset := range got {
		before[i] = append([]int{}, subset...)
	}

	// Mutate each nonempty subset in turn to detect shared backing storage.
	for i, subset := range got {
		if len(subset) == 0 {
			continue
		}
		original := subset[0]
		subset[0] = 99
		for j := range got {
			if j != i && !reflect.DeepEqual(append([]int{}, got[j]...), before[j]) {
				t.Errorf("mutating subset %d changed subset %d: got %v, want %v", i, j, got[j], before[j])
			}
		}
		subset[0] = original
	}
}
