package l456

import (
	"slices"
	"testing"
)

func TestFind132pattern(t *testing.T) {
	for _, tt := range []struct {
		name string
		nums []int
		want bool
	}{
		{"increasing", []int{1, 2, 3, 4}, false}, {"example", []int{3, 1, 4, 2}, true},
		{"negative", []int{-1, 3, 2, 0}, true}, {"single", []int{1}, false},
		{"two", []int{1, 2}, false}, {"decreasing", []int{4, 3, 2, 1}, false},
		{"equal values", []int{1, 1, 1}, false}, {"equal two and three", []int{1, 2, 2}, false},
		{"equal one and two", []int{1, 2, 1}, false}, {"multiple pops", []int{3, 5, 0, 3, 4}, true},
		{"no valid ordering", []int{1, 0, 1, -4, -3}, false},
	} {
		t.Run(tt.name, func(t *testing.T) {
			if got := find132pattern(slices.Clone(tt.nums)); got != tt.want {
				t.Errorf("%s: got %t, want %t", t.Name(), got, tt.want)
			}
		})
	}
}
func TestFind132patternMaximum(t *testing.T) {
	nums := make([]int, 200000)
	for i := range nums {
		nums[i] = i
	}
	if got := find132pattern(nums); got {
		t.Errorf("%s: got true for increasing array", t.Name())
	}
}
