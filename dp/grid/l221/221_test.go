package l221

import "testing"

func TestMaximalSquare(t *testing.T) {
	tests := []struct {
		name string
		rows []string
		want int
	}{
		{name: "example one", rows: []string{"10100", "10111", "11111", "10010"}, want: 4},
		{name: "example two", rows: []string{"01", "10"}, want: 1},
		{name: "example three", rows: []string{"0"}, want: 0},
		{name: "single one", rows: []string{"1"}, want: 1},
		{name: "empty matrix", want: 0},
		{name: "empty row", rows: []string{""}, want: 0},
		{name: "single row", rows: []string{"11111"}, want: 1},
		{name: "single column", rows: []string{"1", "1", "0", "1"}, want: 1},
		{name: "all zero", rows: []string{"000", "000"}, want: 0},
		{name: "all one square", rows: []string{"111", "111", "111"}, want: 9},
		{name: "wide rectangle", rows: []string{"11111", "11111"}, want: 4},
		{name: "tall rectangle", rows: []string{"11", "11", "11", "11"}, want: 4},
		{name: "zero in center", rows: []string{"111", "101", "111"}, want: 1},
		{name: "diagonal neighbor limits square", rows: []string{"011", "111", "111"}, want: 4},
		{name: "square at bottom right", rows: []string{"0000", "0111", "0111", "0111"}, want: 9},
		{name: "largest square before last cell", rows: []string{"1110", "1110", "1110", "0000"}, want: 9},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			matrix := make([][]byte, len(tt.rows))
			for i, row := range tt.rows {
				matrix[i] = []byte(row)
			}
			if got := maximalSquare(matrix); got != tt.want {
				t.Errorf("maximalSquare(%q) = %d, want %d", tt.rows, got, tt.want)
			}
			for i, row := range matrix {
				if string(row) != tt.rows[i] {
					t.Errorf("matrix row %d was modified: got %q, want %q", i, row, tt.rows[i])
				}
			}
		})
	}
}

func TestMaximalSquareMaximumDimensions(t *testing.T) {
	matrix := make([][]byte, 300)
	for i := range matrix {
		matrix[i] = make([]byte, 300)
		for j := range matrix[i] {
			matrix[i][j] = '1'
		}
	}
	if got := maximalSquare(matrix); got != 300*300 {
		t.Errorf("maximalSquare(300-by-300 ones) = %d, want %d", got, 300*300)
	}
}
