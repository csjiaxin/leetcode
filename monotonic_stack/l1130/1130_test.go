package l1130

import (
	"slices"
	"testing"
)

func TestMctFromLeafValues(t *testing.T) {
	for _, tt := range []struct {
		name string
		arr  []int
		want int
	}{
		{"example one", []int{6, 2, 4}, 32}, {"example two", []int{4, 11}, 44},
		{"minimum", []int{1, 1}, 1}, {"increasing", []int{1, 2, 3, 4}, 20},
		{"decreasing", []int{4, 3, 2, 1}, 20}, {"equal", []int{3, 3, 3}, 18},
		{"small middle", []int{5, 1, 4}, 24}, {"large middle", []int{1, 5, 2}, 15},
	} {
		t.Run(tt.name, func(t *testing.T) {
			if got := mctFromLeafValues(slices.Clone(tt.arr)); got != tt.want {
				t.Errorf("%s: got %d, want %d", t.Name(), got, tt.want)
			}
		})
	}
}
func TestMctFromLeafValuesMaximum(t *testing.T) {
	arr := make([]int, 40)
	for i := range arr {
		arr[i] = 15
	}
	if got := mctFromLeafValues(arr); got != 8775 {
		t.Errorf("%s: got %d, want 8775", t.Name(), got)
	}
}
