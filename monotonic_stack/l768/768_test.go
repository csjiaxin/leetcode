package l768

import (
	"slices"
	"testing"
)

func TestMaxChunksToSorted(t *testing.T) {
	for _, tt := range []struct {
		name string
		arr  []int
		want int
	}{
		{"example one", []int{5, 4, 3, 2, 1}, 1}, {"example two", []int{2, 1, 3, 4, 4}, 4},
		{"single", []int{1}, 1}, {"increasing", []int{1, 2, 3, 4}, 4}, {"equal", []int{2, 2, 2}, 3},
		{"merge several blocks", []int{1, 3, 5, 4, 2, 6}, 3}, {"preserve merged maximum", []int{2, 0, 1, 3}, 2},
		{"duplicates across blocks", []int{1, 0, 1, 3, 2}, 3}, {"zeros", []int{0, 0}, 2},
	} {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxChunksToSorted(slices.Clone(tt.arr)); got != tt.want {
				t.Errorf("%s: got %d, want %d", t.Name(), got, tt.want)
			}
		})
	}
}
func TestMaxChunksToSortedMaximum(t *testing.T) {
	arr := make([]int, 2000)
	for i := range arr {
		arr[i] = 100000000
	}
	if got := maxChunksToSorted(arr); got != len(arr) {
		t.Errorf("%s: got %d, want %d", t.Name(), got, len(arr))
	}
}
