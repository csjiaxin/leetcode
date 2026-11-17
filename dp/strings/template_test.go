package stringsdp

import "testing"

func TestTableEmptyPrefixes(t *testing.T) {
	for _, tc := range []struct {
		name string
		m, n int
	}{
		{name: "both_empty", m: 0, n: 0},
		{name: "first_empty", m: 0, n: 4},
		{name: "second_empty", m: 3, n: 0},
	} {
		t.Run(tc.name, func(t *testing.T) {
			dp := Table(tc.m, tc.n,
				func(i, j int) int { return i + j },
				func(i, j, up, left, diagonal int) int {
					t.Fatal("empty-prefix table must not invoke step")
					return 0
				},
			)
			if len(dp) != tc.m+1 {
				t.Fatalf("rows = %d, want %d", len(dp), tc.m+1)
			}
			for i, row := range dp {
				if len(row) != tc.n+1 {
					t.Fatalf("row %d columns = %d, want %d", i, len(row), tc.n+1)
				}
				for j, got := range row {
					if got != i+j {
						t.Errorf("dp[%d][%d] = %d, want %d", i, j, got, i+j)
					}
				}
			}
		})
	}
}

func TestTablePredecessorsAreReady(t *testing.T) {
	type coordinate struct{ row, column int }
	dp := Table(3, 4,
		func(i, j int) coordinate { return coordinate{i, j} },
		func(i, j int, up, left, diagonal coordinate) coordinate {
			for _, neighbor := range []struct {
				name      string
				got, want coordinate
			}{
				{"up", up, coordinate{i - 1, j}},
				{"left", left, coordinate{i, j - 1}},
				{"diagonal", diagonal, coordinate{i - 1, j - 1}},
			} {
				if neighbor.got != neighbor.want {
					t.Fatalf("cell (%d,%d) %s = %+v, want %+v", i, j, neighbor.name, neighbor.got, neighbor.want)
				}
			}
			return coordinate{i, j}
		},
	)
	if got, want := dp[3][4], (coordinate{3, 4}); got != want {
		t.Errorf("last cell = %+v, want %+v", got, want)
	}
}
