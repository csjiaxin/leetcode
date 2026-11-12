package l1901

import "testing"

func TestFindPeakGrid(t *testing.T) {
	for _, tt := range []struct {
		name string
		mat  [][]int
	}{
		{"example", [][]int{{1, 4}, {3, 2}}},
		{"multiple", [][]int{{10, 20, 15}, {21, 30, 14}, {7, 16, 32}}},
		{"single", [][]int{{1}}}, {"single column", [][]int{{1}, {5}, {3}}},
		{"single row", [][]int{{1, 2, 4, 3}}}, {"left peak", [][]int{{9, 7, 5, 3}, {8, 6, 4, 2}}},
		{"right peak", [][]int{{1, 3, 5, 7}, {2, 4, 6, 8}}},
	} {
		t.Run(tt.name, func(t *testing.T) {
			p := findPeakGrid(tt.mat)
			if len(p) != 2 {
				t.Fatalf("invalid coordinate %v", p)
			}
			i, j := p[0], p[1]
			if i < 0 || i >= len(tt.mat) || j < 0 || j >= len(tt.mat[0]) {
				t.Fatalf("out of range %v", p)
			}
			for _, d := range [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
				x, y := i+d[0], j+d[1]
				if x >= 0 && x < len(tt.mat) && y >= 0 && y < len(tt.mat[0]) && tt.mat[x][y] >= tt.mat[i][j] {
					t.Errorf("%v is not a peak", p)
				}
			}
		})
	}
}
