package l287

import (
	"slices"
	"testing"
)

func TestFindDuplicate(t *testing.T) {
	for _, tt := range []struct {
		name string
		nums []int
		want int
	}{
		{"example", []int{1, 3, 4, 2, 2}, 2}, {"other", []int{3, 1, 3, 4, 2}, 3},
		{"minimum", []int{1, 1}, 1}, {"many repeats", []int{3, 3, 3, 3, 3}, 3},
		{"duplicate one", []int{1, 4, 2, 1, 3}, 1}, {"high duplicate", []int{4, 1, 2, 3, 4}, 4},
	} {
		t.Run(tt.name, func(t *testing.T) {
			nums := slices.Clone(tt.nums)
			got := findDuplicate(nums)
			if got != tt.want {
				t.Errorf("got %d want %d", got, tt.want)
			}
			if !slices.Equal(nums, tt.nums) {
				t.Error("input modified; problem requires read-only input")
			}
		})
	}
}
