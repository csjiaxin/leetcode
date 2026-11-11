package l857

import (
	"math"
	"testing"
)

func TestMincostToHireWorkers(t *testing.T) {
	for _, tt := range []struct {
		name          string
		quality, wage []int
		k             int
		want          float64
	}{
		{"example one", []int{10, 20, 5}, []int{70, 50, 30}, 2, 105}, {"example two", []int{3, 1, 10, 10, 1}, []int{4, 8, 2, 2, 7}, 3, 30.666666666666668}, {"one worker", []int{10, 2}, []int{9, 8}, 1, 8}, {"equal ratios", []int{1, 2, 3}, []int{2, 4, 6}, 2, 6}, {"all workers", []int{10, 20}, []int{70, 50}, 2, 210},
	} {
		t.Run(tt.name, func(t *testing.T) {
			if got := mincostToHireWorkers(tt.quality, tt.wage, tt.k); math.Abs(got-tt.want) > 1e-7*math.Max(1, math.Abs(tt.want)) {
				t.Errorf("%s: got %.10f want %.10f", t.Name(), got, tt.want)
			}
		})
	}
}
