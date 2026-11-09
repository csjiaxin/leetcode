package l63

import "testing"

func TestUniquePathsWithObstacles(t *testing.T) {
	tests := []struct {
		name string
		grid [][]int
		want int
	}{
		{
			name: "example one",
			grid: [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}},
			want: 2,
		},
		{
			name: "example two",
			grid: [][]int{{0, 1}, {0, 0}},
			want: 1,
		},
		{
			name: "single open cell",
			grid: [][]int{{0}},
			want: 1,
		},
		{
			name: "single blocked cell",
			grid: [][]int{{1}},
			want: 0,
		},
		{
			name: "blocked start",
			grid: [][]int{{1, 0}, {0, 0}},
			want: 0,
		},
		{
			name: "blocked destination",
			grid: [][]int{{0, 0}, {0, 1}},
			want: 0,
		},
		{
			name: "single open row",
			grid: [][]int{{0, 0, 0, 0}},
			want: 1,
		},
		{
			name: "single blocked row",
			grid: [][]int{{0, 1, 0, 0}},
			want: 0,
		},
		{
			name: "single open column",
			grid: [][]int{{0}, {0}, {0}, {0}},
			want: 1,
		},
		{
			name: "single blocked column",
			grid: [][]int{{0}, {1}, {0}, {0}},
			want: 0,
		},
		{
			name: "no obstacles",
			grid: [][]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}},
			want: 6,
		},
		{
			name: "complete barrier",
			grid: [][]int{{0, 0, 0}, {1, 1, 1}, {0, 0, 0}},
			want: 0,
		},
		{
			name: "first row obstacle allows paths from below",
			grid: [][]int{{0, 1, 0}, {0, 0, 0}, {0, 0, 0}},
			want: 3,
		},
		{
			name: "first column obstacle allows paths from right",
			grid: [][]int{{0, 0, 0}, {1, 0, 0}, {0, 0, 0}},
			want: 3,
		},
		{
			name: "multiple obstacles leave one route",
			grid: [][]int{{0, 0, 0, 0}, {0, 1, 1, 0}, {0, 0, 1, 0}},
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Allow implementations that reuse the input grid for DP storage.
			grid := make([][]int, len(tt.grid))
			for i, row := range tt.grid {
				grid[i] = append([]int(nil), row...)
			}
			if got := uniquePathsWithObstacles(grid); got != tt.want {
				t.Errorf("uniquePathsWithObstacles(%v) = %d, want %d", tt.grid, got, tt.want)
			}
		})
	}
}

func TestUniquePathsWithObstaclesMaximumDimensions(t *testing.T) {
	// A blocked interior leaves exactly two routes along the outer edges.
	grid := make([][]int, 100)
	for i := range grid {
		grid[i] = make([]int, 100)
		for j := range grid[i] {
			if i > 0 && i < 99 && j > 0 && j < 99 {
				grid[i][j] = 1
			}
		}
	}
	if got := uniquePathsWithObstacles(grid); got != 2 {
		t.Errorf("uniquePathsWithObstacles(100-by-100 grid with blocked interior) = %d, want 2", got)
	}
}
