package l452

import "testing"

func TestFindMinArrowShots(t *testing.T) {
	for _, tt := range []struct {
		name   string
		points [][]int
		want   int
	}{
		{"example", [][]int{{10, 16}, {2, 8}, {1, 6}, {7, 12}}, 2},
		{"disjoint", [][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}}, 4},
		{"touching shared arrow", [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}}, 2},
		{"single", [][]int{{1, 2}}, 1},
		{"nested", [][]int{{1, 10}, {2, 9}, {3, 8}}, 1},
		{"duplicates", [][]int{{1, 2}, {1, 2}}, 1},
		{"chain not common intersection", [][]int{{1, 3}, {2, 4}, {3, 5}, {4, 6}}, 2},
		{"integer extremes", [][]int{{-2147483648, -2147483647}, {2147483646, 2147483647}}, 2},
	} {
		t.Run(tt.name, func(t *testing.T) {
			points := make([][]int, len(tt.points))
			for i, v := range tt.points {
				points[i] = append([]int(nil), v...)
			}
			if got := findMinArrowShots(points); got != tt.want {
				t.Errorf("findMinArrowShots(%v) = %d, want %d", tt.points, got, tt.want)
			}
		})
	}
}
