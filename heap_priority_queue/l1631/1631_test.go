package l1631

import "testing"

func TestMinimumEffortPath(t *testing.T) {
	for _, tt := range []struct {
		name    string
		heights [][]int
		want    int
	}{
		{"example one", [][]int{{1, 2, 2}, {3, 8, 2}, {5, 3, 5}}, 2}, {"example two", [][]int{{1, 2, 3}, {3, 8, 4}, {5, 3, 5}}, 1}, {"single", [][]int{{100}}, 0}, {"flat", [][]int{{2, 2}, {2, 2}}, 0}, {"one row", [][]int{{1, 10, 6}}, 9}, {"one column", [][]int{{1}, {2}, {4}}, 2}, {"avoid high center", [][]int{{1, 1, 1}, {1, 100, 1}, {1, 1, 1}}, 0},
	} {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumEffortPath(tt.heights); got != tt.want {
				t.Errorf("%s: got %d want %d", t.Name(), got, tt.want)
			}
		})
	}
}
