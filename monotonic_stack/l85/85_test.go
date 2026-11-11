package l85

import "testing"

func TestMaximalRectangle(t *testing.T) {
	for _, tt := range []struct {
		name string
		rows []string
		want int
	}{
		{"empty matrix", nil, 0}, {"empty row", []string{""}, 0},
		{"example", []string{"10100", "10111", "11111", "10010"}, 6},
		{"single zero", []string{"0"}, 0}, {"single one", []string{"1"}, 1},
		{"one row", []string{"11011"}, 2}, {"one column", []string{"1", "1", "0", "1"}, 2},
		{"all ones", []string{"111", "111"}, 6}, {"height reset", []string{"111", "101", "111"}, 3},
		{"all zeros", []string{"00", "00"}, 0},
	} {
		t.Run(tt.name, func(t *testing.T) {
			matrix := make([][]byte, len(tt.rows))
			for i, row := range tt.rows {
				matrix[i] = []byte(row)
			}
			if got := maximalRectangle(matrix); got != tt.want {
				t.Errorf("%s: got %d, want %d", t.Name(), got, tt.want)
			}
		})
	}
}
func TestMaximalRectangleMaximum(t *testing.T) {
	matrix := make([][]byte, 200)
	for r := range matrix {
		matrix[r] = make([]byte, 200)
		for c := range matrix[r] {
			matrix[r][c] = '1'
		}
	}
	if got := maximalRectangle(matrix); got != 40000 {
		t.Errorf("%s: got %d, want 40000", t.Name(), got)
	}
}
