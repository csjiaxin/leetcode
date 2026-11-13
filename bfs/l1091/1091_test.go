package l1091

import "testing"

func TestShortestPathBinaryMatrix(t *testing.T) {
	for _, tt := range []struct {
		name string
		grid [][]int
		want int
	}{
		{"diagonal", [][]int{{0, 1}, {1, 0}}, 2},
		{"example", [][]int{{0, 0, 0}, {1, 1, 0}, {1, 1, 0}}, 4},
		{"single open", [][]int{{0}}, 1}, {"blocked start", [][]int{{1}}, -1},
		{"blocked end", [][]int{{0, 0}, {0, 1}}, -1},
		{"open grid", [][]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}}, 3},
		{"disconnected", [][]int{{0, 1, 1}, {1, 1, 1}, {1, 1, 0}}, -1},
	} {
		t.Run(tt.name, func(t *testing.T) {
			input := make([][]int, len(tt.grid))
			for i := range input {
				input[i] = append([]int(nil), tt.grid[i]...)
			}
			if got := shortestPathBinaryMatrix(input); got != tt.want {
				t.Errorf("got %d want %d", got, tt.want)
			}
		})
	}
}
