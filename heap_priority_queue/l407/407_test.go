package l407

import "testing"

func TestTrapRainWater(t *testing.T) {
	for _, tt := range []struct {
		name    string
		heights [][]int
		want    int
	}{
		{"example one", [][]int{{1, 4, 3, 1, 3, 2}, {3, 2, 1, 3, 2, 4}, {2, 3, 3, 2, 3, 1}}, 4}, {"example two", [][]int{{3, 3, 3, 3, 3}, {3, 2, 2, 2, 3}, {3, 2, 1, 2, 3}, {3, 2, 2, 2, 3}, {3, 3, 3, 3, 3}}, 10}, {"single", [][]int{{1}}, 0}, {"one row", [][]int{{3, 1, 3}}, 0}, {"basin", [][]int{{5, 5, 5}, {5, 1, 5}, {5, 5, 5}}, 4}, {"escape", [][]int{{5, 1, 5}, {5, 1, 5}, {5, 5, 5}}, 0}, {"flat", [][]int{{2, 2, 2}, {2, 2, 2}, {2, 2, 2}}, 0},
	} {
		t.Run(tt.name, func(t *testing.T) {
			if got := trapRainWater(tt.heights); got != tt.want {
				t.Errorf("%s: got %d want %d", t.Name(), got, tt.want)
			}
		})
	}
}
