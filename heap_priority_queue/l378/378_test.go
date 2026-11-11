package l378

import "testing"

func TestKthSmallest(t *testing.T) {
	for _, tt := range []struct {
		name    string
		matrix  [][]int
		k, want int
	}{
		{"example", [][]int{{1, 5, 9}, {10, 11, 13}, {12, 13, 15}}, 8, 13}, {"single", [][]int{{-5}}, 1, -5}, {"duplicates", [][]int{{1, 1}, {1, 2}}, 3, 1}, {"smallest", [][]int{{-5, -4}, {-3, 0}}, 1, -5}, {"largest", [][]int{{1, 2}, {3, 4}}, 4, 4},
	} {
		t.Run(tt.name, func(t *testing.T) {
			if got := kthSmallest(tt.matrix, tt.k); got != tt.want {
				t.Errorf("%s: got %d want %d", t.Name(), got, tt.want)
			}
		})
	}
}
