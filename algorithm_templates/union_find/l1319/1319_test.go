package l1319

import "testing"

func TestMakeConnected(t *testing.T) {
	for _, tt := range []struct {
		name  string
		n     int
		edges [][]int
		want  int
	}{
		{"example", 4, [][]int{{0, 1}, {0, 2}, {1, 2}}, 1},
		{"two moves", 6, [][]int{{0, 1}, {0, 2}, {0, 3}, {1, 2}, {1, 3}}, 2},
		{"insufficient cables", 6, [][]int{{0, 1}, {0, 2}, {0, 3}, {1, 2}}, -1},
		{"single", 1, nil, 0}, {"two disconnected", 2, nil, -1},
		{"connected", 3, [][]int{{0, 1}, {1, 2}}, 0},
	} {
		t.Run(tt.name, func(t *testing.T) {
			input := make([][]int, len(tt.edges))
			for i := range input {
				input[i] = append([]int(nil), tt.edges[i]...)
			}
			if got := makeConnected(tt.n, input); got != tt.want {
				t.Errorf("got %d want %d", got, tt.want)
			}
		})
	}
}
