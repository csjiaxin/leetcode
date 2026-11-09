package l741

import "testing"

func TestCherryPickup(t *testing.T) {
	tests := []struct {
		name string
		grid [][]int
		want int
	}{
		{
			name: "example one",
			grid: [][]int{{0, 1, -1}, {1, 0, -1}, {1, 1, 1}},
			want: 5,
		},
		{
			name: "example two no complete path",
			grid: [][]int{{1, 1, -1}, {1, -1, 1}, {-1, 1, 1}},
			want: 0,
		},
		{name: "single empty cell", grid: [][]int{{0}}, want: 0},
		{name: "single cherry counted once", grid: [][]int{{1}}, want: 1},
		{name: "no cherries", grid: [][]int{{0, 0}, {0, 0}}, want: 0},
		{
			name: "different outbound and return paths",
			grid: [][]int{{1, 1}, {1, 1}},
			want: 4,
		},
		{
			name: "shared endpoints counted once",
			grid: [][]int{{1, 0}, {0, 1}},
			want: 2,
		},
		{
			name: "single route cherries counted once",
			grid: [][]int{{1, 1, -1}, {-1, 1, -1}, {-1, 1, 1}},
			want: 5,
		},
		{
			name: "routes share interior bottleneck",
			grid: [][]int{{1, 1, -1}, {1, 1, 1}, {-1, 1, 1}},
			want: 7,
		},
		{
			name: "cherry in dead end cannot be collected",
			grid: [][]int{{0, 1, -1}, {1, -1, -1}, {0, 0, 0}},
			want: 1,
		},
		{
			name: "cherries reachable from start but destination blocked off",
			grid: [][]int{{1, 1, 1}, {-1, -1, -1}, {1, 1, 1}},
			want: 0,
		},
		{
			name: "all cherries three by three",
			grid: [][]int{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}},
			want: 8,
		},
		{
			name: "all cherries four by four",
			grid: [][]int{{1, 1, 1, 1}, {1, 1, 1, 1}, {1, 1, 1, 1}, {1, 1, 1, 1}},
			want: 12,
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
	const size = 50
	tests := []struct {
		name  string
		value int
		want  int
	}{
		{name: "all empty", value: 0, want: 0},
		// Each path visits 2*size-1 cells. Two paths can share only their
		// endpoints, collecting 2*(2*size-1)-2 distinct cherries.
		{name: "all cherries", value: 1, want: 4*size - 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			grid := make([][]int, size)
			for i := range grid {
				grid[i] = make([]int, size)
				for j := range grid[i] {
					grid[i][j] = tt.value
				}
			}
			if got := cherryPickup(grid); got != tt.want {
				t.Errorf("cherryPickup(%d-by-%d grid of %d) = %d, want %d", size, size, tt.value, got, tt.want)
			}
		})
	}
}
