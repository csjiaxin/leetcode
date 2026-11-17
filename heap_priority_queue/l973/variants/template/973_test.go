package l973

import (
	"slices"
	"testing"
)

func squared(p []int) int64 { return int64(p[0])*int64(p[0]) + int64(p[1])*int64(p[1]) }

func TestKClosest(t *testing.T) {
	for _, tt := range []struct {
		name   string
		points [][]int
		k      int
	}{
		{"example", [][]int{{1, 3}, {-2, 2}}, 1},
		{"example two", [][]int{{3, 3}, {5, -1}, {-2, 4}}, 2},
		{"origin", [][]int{{0, 0}, {1, 1}}, 1},
		{"all points", [][]int{{1, 1}, {2, 2}, {3, 3}}, 3},
		{"negative", [][]int{{-5, -5}, {1, 1}, {2, 2}, {3, 3}}, 2},
		{"duplicates", [][]int{{1, 1}, {1, 1}, {4, 4}}, 2},
	} {
		t.Run(tt.name, func(t *testing.T) {
			input := make([][]int, len(tt.points))
			counts := map[[2]int]int{}
			distances := make([]int64, len(tt.points))
			for i, p := range tt.points {
				input[i] = slices.Clone(p)
				counts[[2]int{p[0], p[1]}]++
				distances[i] = squared(p)
			}
			slices.Sort(distances)
			cutoff := distances[tt.k-1]
			got := kClosest(input, tt.k)
			if len(got) != tt.k {
				t.Fatalf("got %v want %d points", got, tt.k)
			}
			for _, p := range got {
				if len(p) != 2 {
					t.Fatalf("invalid point %v", p)
				}
				key := [2]int{p[0], p[1]}
				if counts[key] == 0 {
					t.Fatalf("point not available in input %v", p)
				}
				counts[key]--
				if squared(p) > cutoff {
					t.Errorf("point %v exceeds distance cutoff %d", p, cutoff)
				}
			}
			for p, n := range counts {
				if n > 0 && squared(p[:]) < cutoff {
					t.Errorf("omitted closer point %v", p)
				}
			}
		})
	}
}
