package l743

import "testing"

func TestNetworkDelayTime(t *testing.T) {
	for _, tt := range []struct {
		name       string
		times      [][]int
		n, k, want int
	}{
		{"example", [][]int{{2, 1, 1}, {2, 3, 1}, {3, 4, 1}}, 4, 2, 2}, {"two nodes", [][]int{{1, 2, 1}}, 2, 1, 1}, {"wrong direction", [][]int{{1, 2, 1}}, 2, 2, -1}, {"single", nil, 1, 1, 0}, {"stale shorter route", [][]int{{1, 2, 10}, {1, 3, 1}, {3, 2, 1}, {2, 4, 1}}, 4, 1, 3}, {"zero weight", [][]int{{1, 2, 0}, {2, 3, 0}, {3, 1, 0}}, 3, 1, 0},
	} {
		t.Run(tt.name, func(t *testing.T) {
			if got := networkDelayTime(tt.times, tt.n, tt.k); got != tt.want {
				t.Errorf("%s: got %d want %d", t.Name(), got, tt.want)
			}
		})
	}
}
