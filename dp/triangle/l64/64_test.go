package l64

import "testing"

func TestMinPathSum(t *testing.T) {
	tests := []struct {
		name string
		grid [][]int
		want int
	}{
		{
			name: "example one",
			grid: [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}},
			want: 7,
		},
		{
			name: "example two",
			grid: [][]int{{1, 2, 3}, {4, 5, 6}},
			want: 12,
		},
		{
			name: "single cell",
			grid: [][]int{{7}},
			want: 7,
		},
		{
			name: "single zero cell",
			grid: [][]int{{0}},
			want: 0,
		},
		{
			name: "single row",
			grid: [][]int{{1, 2, 3, 4}},
			want: 10,
		},
		{
			name: "single column",
			grid: [][]int{{1}, {2}, {3}, {4}},
			want: 10,
		},
		{
			name: "all zero",
			grid: [][]int{{0, 0, 0}, {0, 0, 0}},
			want: 0,
		},
		{
			name: "prefer right first",
			grid: [][]int{{1, 1}, {9, 1}},
			want: 3,
		},
		{
			name: "prefer down first",
			grid: [][]int{{1, 9}, {1, 1}},
			want: 3,
		},
		{
			name: "smaller next cell is not always optimal",
			grid: [][]int{{1, 1, 100}, {2, 100, 100}, {1, 1, 1}},
			want: 6,
		},
		{
			name: "zero cost route",
			grid: [][]int{{0, 0, 9}, {9, 0, 9}, {9, 0, 0}},
			want: 0,
		},
		{
			name: "equal cost routes",
			grid: [][]int{{2, 2, 2}, {2, 2, 2}, {2, 2, 2}},
			want: 10,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Use independent rows so in-place DP implementations are supported.
			grid := make([][]int, len(tt.grid))
			for i, row := range tt.grid {
				grid[i] = append([]int(nil), row...)
			}
			if got := minPathSum(grid); got != tt.want {
				t.Errorf("minPathSum(%v) = %d, want %d", tt.grid, got, tt.want)
			}
		})
	}
}

func TestMinPathSumMaximumInput(t *testing.T) {
	grid := make([][]int, 200)
	for i := range grid {
		grid[i] = make([]int, 200)
		for j := range grid[i] {
			grid[i][j] = 200
		}
	}
	// Every path visits 200 + 200 - 1 cells.
	if got := minPathSum(grid); got != 79800 {
		t.Errorf("minPathSum(200-by-200 grid with 200 each) = %d, want 79800", got)
	}
}
