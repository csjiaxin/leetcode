package l1289

import "testing"

func TestMinFallingPathSum(t *testing.T) {
	tests := []struct {
		name string
		grid [][]int
		want int
	}{
		{
			name: "example one",
			grid: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			want: 13,
		},
		{
			name: "example two",
			grid: [][]int{{7}},
			want: 7,
		},
		{name: "single negative cell", grid: [][]int{{-99}}, want: -99},
		{name: "single zero cell", grid: [][]int{{0}}, want: 0},
		{
			name: "cannot reuse column in consecutive rows",
			grid: [][]int{{1, 10}, {2, 20}},
			want: 12,
		},
		{
			name: "can jump across columns and return later",
			grid: [][]int{{1, 100, 100}, {100, 100, 1}, {1, 100, 100}},
			want: 3,
		},
		{
			name: "tied minimum values in different columns",
			grid: [][]int{{1, 1, 9}, {1, 1, 9}, {1, 1, 9}},
			want: 3,
		},
		{
			name: "all equal",
			grid: [][]int{{5, 5, 5}, {5, 5, 5}, {5, 5, 5}},
			want: 15,
		},
		{
			name: "all zero",
			grid: [][]int{{0, 0}, {0, 0}},
			want: 0,
		},
		{
			name: "all negative",
			grid: [][]int{{-1, -2}, {-3, -4}},
			want: -5,
		},
		{
			name: "smallest starting cell is not optimal",
			grid: [][]int{{1, 2}, {-99, 99}},
			want: -97,
		},
		{
			name: "alternating negative columns",
			grid: [][]int{{-9, 5, 5}, {5, -9, 5}, {-9, 5, 5}},
			want: -27,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Use independent rows to allow in-place DP implementations.
			grid := make([][]int, len(tt.grid))
			for i, row := range tt.grid {
				grid[i] = append([]int(nil), row...)
			}
			if got := minFallingPathSum(grid); got != tt.want {
				t.Errorf("minFallingPathSum(%v) = %d, want %d", tt.grid, got, tt.want)
			}
		})
	}
}

func TestMinFallingPathSumMaximumDimensions(t *testing.T) {
	const size = 200
	for _, value := range []int{99, -99} {
		name := "positive"
		if value < 0 {
			name = "negative"
		}
		t.Run(name, func(t *testing.T) {
			grid := make([][]int, size)
			for i := range grid {
				grid[i] = make([]int, size)
				for j := range grid[i] {
					grid[i][j] = value
				}
			}
			want := size * value
			if got := minFallingPathSum(grid); got != want {
				t.Errorf("minFallingPathSum(%d-by-%d grid of %d) = %d, want %d", size, size, value, got, want)
			}
		})
	}
}
