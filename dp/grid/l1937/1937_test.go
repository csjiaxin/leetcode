package l1937

import "testing"

func TestMaxPoints(t *testing.T) {
	tests := []struct {
		name   string
		points [][]int
		want   int64
	}{
		{
			name:   "example one",
			points: [][]int{{1, 2, 3}, {1, 5, 1}, {3, 1, 1}},
			want:   9,
		},
		{
			name:   "example two",
			points: [][]int{{1, 5}, {2, 3}, {4, 2}},
			want:   11,
		},
		{name: "single zero cell", points: [][]int{{0}}, want: 0},
		{name: "single maximum cell", points: [][]int{{100000}}, want: 100000},
		{name: "single row no movement penalty", points: [][]int{{1, 8, 3, 7}}, want: 8},
		{name: "single column", points: [][]int{{2}, {0}, {5}, {3}}, want: 10},
		{name: "all zero", points: [][]int{{0, 0, 0}, {0, 0, 0}}, want: 0},
		{name: "all equal", points: [][]int{{5, 5}, {5, 5}, {5, 5}}, want: 15},
		{
			name:   "move far right and pay distance",
			points: [][]int{{10, 0, 0, 0}, {0, 0, 0, 10}},
			want:   17,
		},
		{
			name:   "move far left and pay distance",
			points: [][]int{{0, 0, 0, 10}, {10, 0, 0, 0}},
			want:   17,
		},
		{
			name:   "largest starting value is not optimal",
			points: [][]int{{10, 9, 9}, {0, 0, 100}},
			want:   109,
		},
		{
			name:   "staying beats chasing row maximum",
			points: [][]int{{10, 0, 0, 0}, {8, 0, 0, 10}},
			want:   18,
		},
		{
			name:   "movement in both directions",
			points: [][]int{{10, 0, 0}, {0, 0, 10}, {10, 0, 0}},
			want:   26,
		},
		{
			name:   "zero row still incurs movement penalties",
			points: [][]int{{10, 0, 0}, {0, 0, 0}, {0, 0, 10}},
			want:   18,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Independent rows allow in-place implementations.
			points := make([][]int, len(tt.points))
			for i, row := range tt.points {
				points[i] = append([]int(nil), row...)
			}
			if got := maxPoints(points); got != tt.want {
				t.Errorf("maxPoints(%v) = %d, want %d", tt.points, got, tt.want)
			}
		})
	}
}

func TestMaxPointsLargeInputs(t *testing.T) {
	tests := []struct {
		name       string
		rows, cols int
	}{
		{name: "maximum rows requires int64", rows: 100000, cols: 1},
		{name: "maximum columns", rows: 1, cols: 100000},
		{name: "wide grid", rows: 2, cols: 50000},
		{name: "rectangular grid", rows: 1000, cols: 100},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			const value = 100000
			points := make([][]int, tt.rows)
			for i := range points {
				points[i] = make([]int, tt.cols)
				for j := range points[i] {
					points[i][j] = value
				}
			}
			// Staying in one column collects the maximum with no penalties.
			want := int64(tt.rows) * value
			if got := maxPoints(points); got != want {
				t.Errorf("maxPoints(%d-by-%d grid of %d) = %d, want %d", tt.rows, tt.cols, value, got, want)
			}
		})
	}
}
