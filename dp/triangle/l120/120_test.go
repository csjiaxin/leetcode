package l120

import "testing"

func TestMinimumTotal(t *testing.T) {
	tests := []struct {
		name     string
		triangle [][]int
		want     int
	}{
		{
			name:     "example one",
			triangle: [][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}},
			want:     11,
		},
		{
			name:     "example two",
			triangle: [][]int{{-10}},
			want:     -10,
		},
		{
			name:     "single positive value",
			triangle: [][]int{{7}},
			want:     7,
		},
		{
			name:     "single zero",
			triangle: [][]int{{0}},
			want:     0,
		},
		{
			name:     "two rows prefer left",
			triangle: [][]int{{1}, {2, 5}},
			want:     3,
		},
		{
			name:     "two rows prefer right",
			triangle: [][]int{{1}, {5, 2}},
			want:     3,
		},
		{
			name:     "all negative",
			triangle: [][]int{{-1}, {-2, -3}, {-4, -5, -6}},
			want:     -10,
		},
		{
			name:     "mixed signs",
			triangle: [][]int{{-1}, {2, 3}, {1, -1, -3}},
			want:     -1,
		},
		{
			name:     "all zero",
			triangle: [][]int{{0}, {0, 0}, {0, 0, 0}},
			want:     0,
		},
		{
			name:     "left edge path",
			triangle: [][]int{{1}, {1, 9}, {1, 9, 9}, {1, 9, 9, 9}},
			want:     4,
		},
		{
			name:     "right edge path",
			triangle: [][]int{{1}, {9, 1}, {9, 9, 1}, {9, 9, 9, 1}},
			want:     4,
		},
		{
			name:     "greedy next value is not optimal",
			triangle: [][]int{{1}, {1, 2}, {100, 100, 1}},
			want:     4,
		},
		{
			name:     "cannot jump to nonadjacent minimum",
			triangle: [][]int{{1}, {1, 100}, {100, 100, 1}},
			want:     102,
		},
		{
			name:     "equal cost paths",
			triangle: [][]int{{2}, {2, 2}, {2, 2, 2}, {2, 2, 2, 2}},
			want:     8,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Allow implementations that reuse the triangle for DP storage.
			triangle := make([][]int, len(tt.triangle))
			for i, row := range tt.triangle {
				triangle[i] = append([]int(nil), row...)
			}
			if got := minimumTotal(triangle); got != tt.want {
				t.Errorf("minimumTotal(%v) = %d, want %d", tt.triangle, got, tt.want)
			}
		})
	}
}

func TestMinimumTotalMaximumRows(t *testing.T) {
	for _, value := range []int{10000, -10000} {
		triangle := make([][]int, 200)
		for i := range triangle {
			triangle[i] = make([]int, i+1)
			for j := range triangle[i] {
				triangle[i][j] = value
			}
		}
		want := 200 * value
		if got := minimumTotal(triangle); got != want {
			t.Errorf("minimumTotal(200 rows with %d each) = %d, want %d", value, got, want)
		}
	}
}
