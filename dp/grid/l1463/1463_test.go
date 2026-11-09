package l1463

import "testing"

func TestCherryPickup(t *testing.T) {
	tests := []struct {
		name string
		grid [][]int
		want int
	}{
		{
			name: "example one",
			grid: [][]int{{3, 1, 1}, {2, 5, 1}, {1, 5, 5}, {2, 1, 1}},
			want: 24,
		},
		{
			name: "example two",
			grid: [][]int{
				{1, 0, 0, 0, 0, 0, 1},
				{2, 0, 0, 0, 0, 3, 0},
				{2, 0, 9, 0, 0, 0, 0},
				{0, 3, 0, 5, 4, 0, 0},
				{1, 0, 2, 3, 0, 0, 6},
			},
			want: 28,
		},
		{name: "minimum grid", grid: [][]int{{1, 2}, {3, 4}}, want: 10},
		{name: "all zero", grid: [][]int{{0, 0}, {0, 0}}, want: 0},
		{
			name: "shared cell counted once",
			grid: [][]int{{1, 1}, {100, 0}},
			want: 102,
		},
		{
			name: "robots should not both take the same cell",
			grid: [][]int{{0, 0, 0}, {0, 10, 0}, {0, 0, 0}},
			want: 10,
		},
		{
			name: "cannot move more than one column per row",
			grid: [][]int{{0, 0, 0, 0, 0}, {0, 0, 100, 0, 0}},
			want: 0,
		},
		{
			name: "can reach center after enough rows",
			grid: [][]int{{0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 100, 0, 0}},
			want: 100,
		},
		{
			name: "staying in the same column is allowed",
			grid: [][]int{{1, 0, 0, 1}, {50, 0, 0, 50}, {50, 0, 0, 50}},
			want: 202,
		},
		{
			name: "greedy first move is not optimal",
			// Robot 1 must skip the 5 to reach column 2 in time.
			grid: [][]int{{0, 0, 0, 0, 0}, {5, 0, 0, 0, 0}, {0, 0, 100, 100, 0}},
			want: 200,
		},
		{
			name: "robots cross over",
			grid: [][]int{{0, 0, 0}, {0, 0, 0}, {50, 0, 60}, {0, 0, 0}},
			want: 110,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Independent rows allow implementations that modify the grid.
			grid := make([][]int, len(tt.grid))
			for i, row := range tt.grid {
				grid[i] = append([]int(nil), row...)
			}
			if got := cherryPickup(grid); got != tt.want {
				t.Errorf("cherryPickup(%v) = %d, want %d", tt.grid, got, tt.want)
			}
		})
	}
}

func TestCherryPickupMaximumDimensions(t *testing.T) {
	tests := []struct {
		name       string
		rows, cols int
		value      int
		want       int
	}{
		{name: "70 by 70 zeros", rows: 70, cols: 70, value: 0, want: 0},
		// Two robots can always stay in distinct cells, collecting two per row.
		{name: "70 by 70 full", rows: 70, cols: 70, value: 100, want: 70 * 200},
		{name: "70 by 2 full", rows: 70, cols: 2, value: 100, want: 70 * 200},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			grid := make([][]int, tt.rows)
			for i := range grid {
				grid[i] = make([]int, tt.cols)
				for j := range grid[i] {
					grid[i][j] = tt.value
				}
			}
			if got := cherryPickup(grid); got != tt.want {
				t.Errorf("cherryPickup(%d-by-%d grid of %d) = %d, want %d", tt.rows, tt.cols, tt.value, got, tt.want)
			}
		})
	}
}
