package l435

import "testing"

func TestEraseOverlapIntervals(t *testing.T) {
	for _, tt := range []struct {
		name      string
		intervals [][]int
		want      int
	}{
		{"example one", [][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}}, 1},
		{"duplicates", [][]int{{1, 2}, {1, 2}, {1, 2}}, 2},
		{"touching allowed", [][]int{{1, 2}, {2, 3}}, 0},
		{"single", [][]int{{1, 2}}, 0},
		{"nested", [][]int{{1, 10}, {2, 3}, {3, 4}, {4, 5}}, 1},
		{"same end", [][]int{{1, 3}, {2, 3}, {3, 4}}, 1},
		{"negative unsorted", [][]int{{0, 2}, {-3, -1}, {-1, 0}}, 0},
		{"all overlap", [][]int{{1, 5}, {2, 6}, {3, 7}}, 2},
	} {
		t.Run(tt.name, func(t *testing.T) {
			intervals := make([][]int, len(tt.intervals))
			for i, v := range tt.intervals {
				intervals[i] = append([]int(nil), v...)
			}
			if got := eraseOverlapIntervals(intervals); got != tt.want {
				t.Errorf("eraseOverlapIntervals(%v) = %d, want %d", tt.intervals, got, tt.want)
			}
		})
	}
}
