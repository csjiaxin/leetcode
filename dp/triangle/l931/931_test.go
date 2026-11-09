package l931

import "testing"

func TestMinFallingPathSum(t *testing.T) {
	tests := []struct {
		name   string
		matrix [][]int
		want   int
	}{
		{
			name:   "example one",
			matrix: [][]int{{2, 1, 3}, {6, 5, 4}, {7, 8, 9}},
			want:   13,
		},
		{
			name:   "example two",
			matrix: [][]int{{-19, 57}, {-40, -5}},
			want:   -59,
		},
		{
			name:   "single positive cell",
			matrix: [][]int{{7}},
			want:   7,
		},
		{
			name:   "single negative cell",
			matrix: [][]int{{-7}},
			want:   -7,
		},
		{
			name:   "single zero cell",
			matrix: [][]int{{0}},
			want:   0,
		},
		{
			name:   "all zero",
			matrix: [][]int{{0, 0}, {0, 0}},
			want:   0,
		},
		{
			name:   "straight down left edge",
			matrix: [][]int{{1, 9, 9}, {1, 9, 9}, {1, 9, 9}},
			want:   3,
		},
		{
			name:   "straight down right edge",
			matrix: [][]int{{9, 9, 1}, {9, 9, 1}, {9, 9, 1}},
			want:   3,
		},
		{
			name:   "diagonal left",
			matrix: [][]int{{9, 9, 1}, {9, 1, 9}, {1, 9, 9}},
			want:   3,
		},
		{
			name:   "diagonal right",
			matrix: [][]int{{1, 9, 9}, {9, 1, 9}, {9, 9, 1}},
			want:   3,
		},
		{
			name:   "cannot jump columns or wrap edges",
			matrix: [][]int{{1, 100, 100}, {100, 100, 1}, {100, 100, 1}},
			want:   102,
		},
		{
			name:   "smallest starting value is not optimal",
			matrix: [][]int{{1, 2, 3}, {100, 100, 1}, {100, 100, 1}},
			want:   4,
		},
		{
			name:   "all negative",
			matrix: [][]int{{-1, -2, -3}, {-4, -5, -6}, {-7, -8, -9}},
			want:   -18,
		},
		{
			name:   "zigzag path",
			matrix: [][]int{{1, 9, 9, 9}, {9, 1, 9, 9}, {1, 9, 9, 9}, {9, 1, 9, 9}},
			want:   4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Allow implementations that reuse the matrix for DP storage.
			matrix := make([][]int, len(tt.matrix))
			for i, row := range tt.matrix {
				matrix[i] = append([]int(nil), row...)
			}
			if got := minFallingPathSum(matrix); got != tt.want {
				t.Errorf("minFallingPathSum(%v) = %d, want %d", tt.matrix, got, tt.want)
			}
		})
	}
}

func TestMinFallingPathSumMaximumDimensions(t *testing.T) {
	for _, value := range []int{100, -100} {
		matrix := make([][]int, 100)
		for i := range matrix {
			matrix[i] = make([]int, 100)
			for j := range matrix[i] {
				matrix[i][j] = value
			}
		}
		want := 100 * value
		if got := minFallingPathSum(matrix); got != want {
			t.Errorf("minFallingPathSum(100-by-100 matrix with %d each) = %d, want %d", value, got, want)
		}
	}
}
