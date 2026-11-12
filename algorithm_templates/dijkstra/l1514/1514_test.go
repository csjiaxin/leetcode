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
		{"indirect", 3, [][]int{{0, 1}, {1, 2}, {0, 2}}, []float64{0.5, 0.5, 0.2}, 0, 2, 0.25},
		{"direct", 3, [][]int{{0, 1}, {1, 2}, {0, 2}}, []float64{0.5, 0.5, 0.3}, 0, 2, 0.3},
		{"disconnected", 3, [][]int{{0, 1}}, []float64{0.5}, 0, 2, 0},
		{"zero probability", 2, [][]int{{0, 1}}, []float64{0}, 0, 1, 0},
		{"undirected", 2, [][]int{{0, 1}}, []float64{0.7}, 1, 0, 0.7},
		{"unit cycle", 3, [][]int{{0, 1}, {1, 2}, {2, 0}}, []float64{1, 1, 1}, 0, 2, 1},
	} {
		t.Run(tt.name, func(t *testing.T) {
			edges := make([][]int, len(tt.edges))
			for i := range edges {
				edges[i] = append([]int(nil), tt.edges[i]...)
			}
			got := maxProbability(tt.n, edges, append([]float64(nil), tt.prob...), tt.start, tt.end)
			if math.IsNaN(got) || math.Abs(got-tt.want) > 1e-9 {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
