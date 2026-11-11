package l215

import "testing"

func TestFindKthLargest(t *testing.T) {
	for _, tt := range []struct {
		name    string
		nums    []int
		k, want int
	}{
		{"example one", []int{3, 2, 1, 5, 6, 4}, 2, 5}, {"duplicates", []int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4, 4}, {"single", []int{-1}, 1, -1}, {"minimum", []int{3, -2, 1}, 3, -2}, {"maximum", []int{-5, 0, 4}, 1, 4}, {"all equal", []int{2, 2, 2}, 2, 2},
	} {
		t.Run(tt.name, func(t *testing.T) {
			if got := findKthLargest(tt.nums, tt.k); got != tt.want {
				t.Errorf("%s: got %d want %d", t.Name(), got, tt.want)
			}
		})
	}
}
