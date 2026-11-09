package l39

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

func TestCombinationSum(t *testing.T) {
	tests := []struct {
		name       string
		candidates []int
		target     int
		want       [][]int
	}{
		{
			name:       "example one",
			candidates: []int{2, 3, 6, 7},
			target:     7,
			want:       [][]int{{2, 2, 3}, {7}},
		},
		{
			name:       "example two",
			candidates: []int{2, 3, 5},
			target:     8,
			want:       [][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}},
		},
		{
			name:       "example three",
			candidates: []int{2},
			target:     1,
			want:       [][]int{},
		},
		{
			name:       "reuse single candidate",
			candidates: []int{3},
			target:     9,
			want:       [][]int{{3, 3, 3}},
		},
		{
			name:       "exact match",
			candidates: []int{7},
			target:     7,
			want:       [][]int{{7}},
		},
		{
			name:       "unreachable target",
			candidates: []int{4, 6},
			target:     7,
			want:       [][]int{},
		},
		{
			name:       "all candidates exceed target",
			candidates: []int{5, 8},
			target:     3,
			want:       [][]int{},
		},
		{
			name:       "unsorted candidates",
			candidates: []int{7, 3, 2, 6},
			target:     7,
			want:       [][]int{{2, 2, 3}, {7}},
		},
	}

	// Ignore ordering within and between combinations, but preserve duplicate
	// counts so repeated combinations cause a test failure.
	counts := func(combinations [][]int) map[string]int {
		result := make(map[string]int)
		for _, combination := range combinations {
			canonical := append([]int(nil), combination...)
			sort.Ints(canonical)
			result[fmt.Sprint(canonical)]++
		}
		return result
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := append([]int(nil), tt.candidates...)
			got := combinationSum(input, tt.target)
			if !reflect.DeepEqual(counts(got), counts(tt.want)) {
				t.Errorf("combinationSum(%v, %d) = %v, want %v (in any order)", tt.candidates, tt.target, got, tt.want)
			}
		})
	}
}

func TestCombinationSumIndependentResults(t *testing.T) {
	got := combinationSum([]int{2, 3, 5}, 8)
	if len(got) != 3 {
		t.Fatalf("got %d combinations, want 3", len(got))
	}

	before := make([][]int, len(got))
	for i, combination := range got {
		if len(combination) == 0 {
			t.Fatalf("combination %d is empty for a positive target", i)
		}
		before[i] = append([]int(nil), combination...)
	}

	got[0][0] = 99
	for i := 1; i < len(got); i++ {
		if !reflect.DeepEqual(got[i], before[i]) {
			t.Errorf("mutating combination 0 changed combination %d: got %v, want %v", i, got[i], before[i])
		}
	}
}
