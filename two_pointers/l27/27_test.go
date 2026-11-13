package l27

import (
	"slices"
	"testing"
)

func TestRemoveElement(t *testing.T) {
	for _, tt := range []struct {
		name string
		nums []int
		val  int
		want []int
	}{
		{"example", []int{3, 2, 2, 3}, 3, []int{2, 2}},
		{"mixed", []int{0, 1, 2, 2, 3, 0, 4, 2}, 2, []int{0, 0, 1, 3, 4}},
		{"empty", nil, 0, nil}, {"all", []int{1, 1}, 1, nil}, {"none", []int{1, 2}, 3, []int{1, 2}},
	} {
		t.Run(tt.name, func(t *testing.T) {
			nums := slices.Clone(tt.nums)
			k := removeElement(nums, tt.val)
			if k < 0 || k > len(nums) {
				t.Fatalf("invalid length %d", k)
			}
			got := nums[:k]
			slices.Sort(got)
			want := slices.Clone(tt.want)
			slices.Sort(want)
			if k != len(want) || !slices.Equal(got, want) {
				t.Errorf("got %v (k=%d), want %v", got, k, want)
			}
		})
	}
}
