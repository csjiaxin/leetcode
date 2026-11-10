package l2684

import "testing"

func TestMaxMoves(t *testing.T) {
	tests := []struct {
		name string
		grid [][]int
		want int
	}{
		{
			name: "example one",
			grid: [][]int{{2, 4, 3, 5}, {5, 4, 9, 3}, {3, 4, 2, 11}, {10, 9, 13, 15}},
			want: 3,
		},
		{
			name: "example two",
			grid: [][]int{{3, 2, 4}, {2, 1, 9}, {1, 1, 7}},
			want: 0,
		},
		{name: "minimum grid one move", grid: [][]int{{1, 2}, {1, 2}}, want: 1},
		{name: "equal values cannot extend path", grid: [][]int{{5, 5}, {5, 5}}, want: 0},
		{
			name: "start from any row",
			grid: [][]int{{100, 2, 3, 4}, {100, 2, 3, 4}, {1, 2, 3, 4}},
			want: 3,
		},
		{
			name: "diagonal up moves",
			grid: [][]int{{100, 100, 100, 4}, {100, 100, 3, 1}, {100, 2, 1, 1}, {1, 1, 1, 1}},
			want: 3,
		},
		{
			name: "diagonal down moves",
			grid: [][]int{{1, 1, 1, 1}, {100, 2, 1, 1}, {100, 100, 3, 1}, {100, 100, 100, 4}},
			want: 3,
		},
		{
			name: "cannot jump rows",
			grid: [][]int{{1, 1, 1}, {1, 1, 1}, {1, 2, 1}},
			want: 1,
		},
		{
			name: "cannot skip columns",
			grid: [][]int{{1, 1, 3}, {1, 1, 3}},
			want: 0,
		},
		{
			name: "cannot start in a later column",
			grid: [][]int{{9, 1, 2, 3}, {9, 1, 2, 3}},
			want: 0,
		},
		{
			name: "path stops before last column",
			grid: [][]int{{1, 2, 3, 1}, {1, 2, 3, 1}},
			want: 2,
		},
		{
			name: "larger next value can lead to dead end",
			grid: [][]int{{1, 100, 1, 1}, {1, 2, 3, 4}},
			want: 3,
		},
		{
			name: "cannot wrap top and bottom rows",
			grid: [][]int{{1, 1, 1}, {100, 1, 1}, {100, 2, 3}},
			want: 0,
		},
		{
			name: "maximum cell value",
			grid: [][]int{{1, 1000000}, {1000000, 1000000}},
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Independent rows allow in-place implementations.
			grid := make([][]int, len(tt.grid))
			for i, row := range tt.grid {
				grid[i] = append([]int(nil), row...)
			}
			if got := maxMoves(grid); got != tt.want {
				t.Errorf("maxMoves(%v) = %d, want %d", tt.grid, got, tt.want)
			}
		})
	}
}

func TestMaxMovesLargeGrids(t *testing.T) {
	tests := []struct {
		name       string
		rows, cols int
		increasing bool
		want       int
	}{
		{name: "maximum columns", rows: 100, cols: 1000, increasing: true, want: 999},
		{name: "maximum rows", rows: 1000, cols: 100, increasing: true, want: 99},
		{name: "large equal grid", rows: 100, cols: 1000, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			grid := make([][]int, tt.rows)
			for i := range grid {
				grid[i] = make([]int, tt.cols)
				for j := range grid[i] {
					grid[i][j] = 1
					if tt.increasing {
						grid[i][j] = j + 1
					}
				}
			}
			if got := maxMoves(grid); got != tt.want {
				t.Errorf("maxMoves(%d-by-%d grid) = %d, want %d", tt.rows, tt.cols, got, tt.want)
			}
		})
	}
}
