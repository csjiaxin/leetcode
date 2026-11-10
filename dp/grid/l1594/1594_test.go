package l1594

import "testing"

func TestMaxProductPath(t *testing.T) {
	tests := []struct {
		name string
		grid [][]int
		want int
	}{
		{
			name: "example one all paths negative",
			grid: [][]int{{-1, -2, -3}, {-2, -3, -3}, {-3, -3, -2}},
			want: -1,
		},
		{
			name: "example two",
			grid: [][]int{{1, -2, 1}, {1, -2, 1}, {3, -4, 1}},
			want: 8,
		},
		{
			name: "example three zero path",
			grid: [][]int{{1, 3}, {0, -4}},
			want: 0,
		},
		{name: "single positive", grid: [][]int{{4}}, want: 4},
		{name: "single negative", grid: [][]int{{-4}}, want: -1},
		{name: "single zero", grid: [][]int{{0}}, want: 0},
		{name: "single row even negatives", grid: [][]int{{-1, 2, -3, 4}}, want: 24},
		{name: "single row odd negatives", grid: [][]int{{1, -2, 3}}, want: -1},
		{name: "single column", grid: [][]int{{-2}, {3}, {-4}}, want: 24},
		{name: "zero in only path", grid: [][]int{{-4, 0, 4}}, want: 0},
		{name: "all positive", grid: [][]int{{1, 2}, {3, 4}}, want: 12},
		{name: "even negative count", grid: [][]int{{-1, -2}, {-3, 4}}, want: 12},
		{name: "zero start", grid: [][]int{{0, 4}, {4, 4}}, want: 0},
		{name: "zero destination", grid: [][]int{{4, -4}, {4, 0}}, want: 0},
		{
			name: "positive path beats zero",
			grid: [][]int{{1, 0}, {2, 3}},
			want: 6,
		},
		{
			name: "negative destination needs minimum predecessor",
			grid: [][]int{{1, -2}, {3, -4}},
			want: 8,
		},
		{
			name: "retain negative intermediate product",
			grid: [][]int{{1, -2, 1}, {3, 4, -4}},
			want: 32,
		},
		{
			name: "no diagonal shortcut",
			grid: [][]int{{1, -1}, {-1, 1}},
			want: -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Independent rows allow in-place implementations.
			grid := make([][]int, len(tt.grid))
			for i, row := range tt.grid {
				grid[i] = append([]int(nil), row...)
			}
			if got := maxProductPath(grid); got != tt.want {
				t.Errorf("maxProductPath(%v) = %d, want %d", tt.grid, got, tt.want)
			}
		})
	}
}

func TestMaxProductPathMaximumDimensions(t *testing.T) {
	const size = 15
	const mod = 1000000007
	tests := []struct {
		name  string
		value int
		want  int
	}{
		{name: "all zeros", value: 0, want: 0},
		{name: "all ones", value: 1, want: 1},
		// Every path visits 29 cells, so its product is negative.
		{name: "all negative fours", value: -4, want: -1},
		// 4^29 = 2^58: exceeds 32-bit range but fits in int64.
		{name: "all fours require modulo", value: 4, want: int((int64(1) << 58) % mod)},
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
			if got := maxProductPath(grid); got != tt.want {
				t.Errorf("maxProductPath(%d-by-%d grid of %d) = %d, want %d", size, size, tt.value, got, tt.want)
			}
		})
	}
}
