package l1514

import (
	"math"
	"testing"
)

func TestMaxProbability(t *testing.T) {
	for _, tt := range []struct {
		name       string
		n          int
		edges      [][]int
		prob       []float64
		start, end int
		want       float64
	}{
		{"indirect wins", 3, [][]int{{0, 1}, {1, 2}, {0, 2}}, []float64{0.5, 0.5, 0.2}, 0, 2, 0.25}, {"direct wins", 3, [][]int{{0, 1}, {1, 2}, {0, 2}}, []float64{0.5, 0.5, 0.3}, 0, 2, 0.3}, {"disconnected", 3, [][]int{{0, 1}}, []float64{0.5}, 0, 2, 0}, {"same node", 2, nil, nil, 0, 0, 1}, {"zero probability", 2, [][]int{{0, 1}}, []float64{0}, 0, 1, 0}, {"undirected", 2, [][]int{{0, 1}}, []float64{0.7}, 1, 0, 0.7}, {"unit cycle", 3, [][]int{{0, 1}, {1, 2}, {2, 0}}, []float64{1, 1, 1}, 0, 2, 1},
	} {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProbability(tt.n, tt.edges, tt.prob, tt.start, tt.end); math.Abs(got-tt.want) > 1e-10 {
				t.Errorf("%s: got %v want %v", t.Name(), got, tt.want)
			}
		})
	}
}
