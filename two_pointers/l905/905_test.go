package l905

import (
	"slices"
	"testing"
)

func TestSortArrayByParity(t *testing.T) {
	for _, tt := range []struct {
		name string
		nums []int
	}{
		{"example", []int{3, 1, 2, 4}}, {"single zero", []int{0}}, {"all odd", []int{1, 3, 5}},
		{"all even", []int{2, 4, 6}}, {"duplicates", []int{3, 2, 3, 2, 0, 1}},
	} {
		t.Run(tt.name, func(t *testing.T) {
			got := sortArrayByParity(slices.Clone(tt.nums))
			seenOdd := false
			for _, x := range got {
				if x%2 != 0 {
					seenOdd = true
				} else if seenOdd {
					t.Errorf("even after odd: %v", got)
					break
				}
			}
			values := slices.Clone(got)
			slices.Sort(values)
			want := slices.Clone(tt.nums)
			slices.Sort(want)
			if !slices.Equal(values, want) {
				t.Errorf("elements changed: got %v want multiset %v", got, tt.nums)
			}
		})
	}
}
