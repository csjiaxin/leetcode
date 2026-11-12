package l994

import "testing"

func TestOrangesRotting(t *testing.T) {
	for _, tt := range []struct {
		name string
		grid [][]int
		want int
	}{
		{"example", [][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}}, 4},
		{"unreachable", [][]int{{2, 1, 1}, {0, 1, 1}, {1, 0, 1}}, -1},
		{"no fresh", [][]int{{0, 2}}, 0}, {"no source", [][]int{{1}}, -1},
		{"empty cell", [][]int{{0}}, 0}, {"multiple sources", [][]int{{2, 1, 1, 1, 2}}, 2},
	} {
		t.Run(tt.name, func(t *testing.T) {
			input := make([][]int, len(tt.grid))
			for i := range input {
				input[i] = append([]int(nil), tt.grid[i]...)
			}
			if got := orangesRotting(input); got != tt.want {
				t.Errorf("got %d want %d", got, tt.want)
			}
		})
	}
}
