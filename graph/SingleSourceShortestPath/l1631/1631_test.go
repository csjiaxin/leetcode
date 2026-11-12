package l1631

import "testing"

func TestMinimumEffortPath(t *testing.T) {
	for _, tt := range []struct {
		name    string
		heights [][]int
		want    int
	}{
		{"example", [][]int{{1, 2, 2}, {3, 8, 2}, {5, 3, 5}}, 2},
		{"example two", [][]int{{1, 2, 3}, {3, 8, 4}, {5, 3, 5}}, 1},
		{"single", [][]int{{7}}, 0}, {"equal", [][]int{{3, 3}, {3, 3}}, 0},
		{"row", [][]int{{1, 10, 6}}, 9}, {"avoid cliff", [][]int{{1, 100}, {2, 3}}, 1},
	} {
		t.Run(tt.name, func(t *testing.T) {
			input := make([][]int, len(tt.heights))
			for i := range input {
				input[i] = append([]int(nil), tt.heights[i]...)
			}
			if got := minimumEffortPath(input); got != tt.want {
				t.Errorf("got %d want %d", got, tt.want)
			}
		})
	}
}
