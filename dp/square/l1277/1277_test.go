package l1277

import "testing"

func TestCountSquares(t *testing.T) {
	tests := []struct {
		name   string
		matrix [][]int
		want   int
	}{
		{
			name:   "example one",
			matrix: [][]int{{0, 1, 1, 1}, {1, 1, 1, 1}, {0, 1, 1, 1}},
			want:   15,
		},
		{
			name:   "example two",
			matrix: [][]int{{1, 0, 1}, {1, 1, 0}, {1, 1, 0}},
			want:   7,
		},
		{name: "nil matrix", want: 0},
		{name: "empty matrix", matrix: [][]int{}, want: 0},
		{name: "empty row", matrix: [][]int{{}}, want: 0},
		{name: "single zero", matrix: [][]int{{0}}, want: 0},
		{name: "single one", matrix: [][]int{{1}}, want: 1},
		{name: "single row", matrix: [][]int{{1, 1, 0, 1, 1}}, want: 4},
		{name: "single column", matrix: [][]int{{1}, {1}, {0}, {1}}, want: 3},
		{name: "all zero", matrix: [][]int{{0, 0}, {0, 0}}, want: 0},
		{name: "two by two ones", matrix: [][]int{{1, 1}, {1, 1}}, want: 5},
		{name: "three by three ones", matrix: [][]int{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}}, want: 14},
		{name: "wide rectangle", matrix: [][]int{{1, 1, 1, 1}, {1, 1, 1, 1}}, want: 11},
		{name: "tall rectangle", matrix: [][]int{{1, 1}, {1, 1}, {1, 1}, {1, 1}}, want: 11},
		{name: "center zero", matrix: [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}, want: 8},
		{name: "diagonal neighbor limits square", matrix: [][]int{{0, 1, 1}, {1, 1, 1}, {1, 1, 1}}, want: 11},
		{name: "isolated ones", matrix: [][]int{{1, 0, 1}, {0, 1, 0}, {1, 0, 1}}, want: 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			matrix := make([][]int, len(tt.matrix))
			for i, row := range tt.matrix {
				matrix[i] = append([]int(nil), row...)
			}
			if got := countSquares(matrix); got != tt.want {
				t.Errorf("countSquares(%v) = %d, want %d", tt.matrix, got, tt.want)
			}
			for i, row := range matrix {
				for j, value := range row {
					if value != tt.matrix[i][j] {
						t.Fatalf("matrix[%d][%d] was modified", i, j)
					}
				}
			}
		})
	}
}

func TestCountSquaresMaximumDimensions(t *testing.T) {
	const size = 300
	matrix := make([][]int, size)
	for i := range matrix {
		matrix[i] = make([]int, size)
		for j := range matrix[i] {
			matrix[i][j] = 1
		}
	}
	// There are (size-side+1)^2 squares of each side length.
	want := size * (size + 1) * (2*size + 1) / 6
	if got := countSquares(matrix); got != want {
		t.Errorf("countSquares(%d-by-%d ones) = %d, want %d", size, size, got, want)
	}
}
