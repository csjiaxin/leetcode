package l329

import "testing"

func TestLongestIncreasingPath(t *testing.T) {
	tests := []struct {
		name   string
		matrix [][]int
		want   int
	}{
		{
			name:   "example one",
			matrix: [][]int{{9, 9, 4}, {6, 6, 8}, {2, 1, 1}},
			want:   4,
		},
		{
			name:   "example two",
			matrix: [][]int{{3, 4, 5}, {3, 2, 6}, {2, 2, 1}},
			want:   4,
		},
		{name: "example three", matrix: [][]int{{1}}, want: 1},
		{name: "single zero", matrix: [][]int{{0}}, want: 1},
		{name: "maximum cell value", matrix: [][]int{{2147483647}}, want: 1},
		{name: "increasing row", matrix: [][]int{{1, 2, 3, 4}}, want: 4},
		{name: "decreasing row", matrix: [][]int{{4, 3, 2, 1}}, want: 4},
		{name: "increasing column", matrix: [][]int{{1}, {2}, {3}, {4}}, want: 4},
		{name: "decreasing column", matrix: [][]int{{4}, {3}, {2}, {1}}, want: 4},
		{name: "equal cells cannot extend path", matrix: [][]int{{7, 7}, {7, 7}}, want: 1},
		{name: "plateau splits row", matrix: [][]int{{1, 2, 2, 3}}, want: 2},
		{name: "diagonal moves forbidden", matrix: [][]int{{1, 9}, {9, 2}}, want: 2},
		{name: "cannot wrap row edges", matrix: [][]int{{2, 1, 3}}, want: 2},
		{name: "cannot wrap column edges", matrix: [][]int{{2}, {1}, {3}}, want: 2},
		{
			name:   "winding path uses all four directions",
			matrix: [][]int{{1, 2, 3}, {8, 9, 4}, {7, 6, 5}},
			want:   9,
		},
		{
			name:   "wide rectangular matrix",
			matrix: [][]int{{1, 2, 3, 4}, {8, 7, 6, 5}},
			want:   8,
		},
		{
			name:   "tall rectangular matrix",
			matrix: [][]int{{1, 8}, {2, 7}, {3, 6}, {4, 5}},
			want:   8,
		},
		{
			name:   "branches converge on shared suffix",
			matrix: [][]int{{1, 2, 3}, {2, 3, 4}, {3, 4, 5}},
			want:   5,
		},
		{
			name:   "full value range",
			matrix: [][]int{{0, 2147483647}, {1, 2147483646}},
			want:   4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			matrix := make([][]int, len(tt.matrix))
			for i, row := range tt.matrix {
				matrix[i] = append([]int(nil), row...)
			}
			if got := longestIncreasingPath(matrix); got != tt.want {
				t.Errorf("longestIncreasingPath(%v) = %d, want %d", tt.matrix, got, tt.want)
			}
		})
	}
}

func TestLongestIncreasingPathMaximumDimensions(t *testing.T) {
	const size = 200
	t.Run("all equal", func(t *testing.T) {
		matrix := make([][]int, size)
		for i := range matrix {
			matrix[i] = make([]int, size)
		}
		if got := longestIncreasingPath(matrix); got != 1 {
			t.Errorf("longestIncreasingPath(200-by-200 zeros) = %d, want 1", got)
		}
	})

	t.Run("snake visits every cell", func(t *testing.T) {
		matrix := make([][]int, size)
		for i := range matrix {
			matrix[i] = make([]int, size)
			for j := range matrix[i] {
				column := j
				if i%2 == 1 {
					column = size - 1 - j
				}
				matrix[i][column] = i*size + j
			}
		}
		if got := longestIncreasingPath(matrix); got != size*size {
			t.Errorf("longestIncreasingPath(200-by-200 snake) = %d, want %d", got, size*size)
		}
	})
}
