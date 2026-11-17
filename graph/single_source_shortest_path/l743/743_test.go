package l743

import "testing"

func TestNetworkDelayTime(t *testing.T) {
	for _, tt := range []struct {
		name       string
		times      [][]int
		n, k, want int
	}{
		{"example", [][]int{{2, 1, 1}, {2, 3, 1}, {3, 4, 1}}, 4, 2, 2},
		{"one edge", [][]int{{1, 2, 1}}, 2, 1, 1}, {"directed", [][]int{{1, 2, 1}}, 2, 2, -1},
		{"stale distance", [][]int{{1, 2, 10}, {1, 3, 1}, {3, 2, 1}}, 3, 1, 2},
		{"zero weight", [][]int{{1, 2, 0}, {2, 3, 0}}, 3, 1, 0},
		{"equal paths", [][]int{{1, 2, 1}, {1, 3, 1}, {2, 4, 1}, {3, 4, 1}}, 4, 1, 2},
	} {
		t.Run(tt.name, func(t *testing.T) {
			input := make([][]int, len(tt.times))
			for i := range input {
				input[i] = append([]int(nil), tt.times[i]...)
			}
			if got := networkDelayTime(input, tt.n, tt.k); got != tt.want {
				t.Errorf("got %d want %d", got, tt.want)
			}
		})
	}
}
