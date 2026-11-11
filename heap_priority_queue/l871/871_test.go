package l871

import "testing"

func TestMinRefuelStops(t *testing.T) {
	for _, tt := range []struct {
		name         string
		target, fuel int
		stations     [][]int
		want         int
	}{
		{"no stops", 1, 1, nil, 0}, {"unreachable", 100, 1, [][]int{{10, 100}}, -1}, {"example", 100, 10, [][]int{{10, 60}, {20, 30}, {30, 30}, {60, 40}}, 2}, {"exact arrival", 100, 50, [][]int{{50, 50}}, 1}, {"gap", 100, 10, [][]int{{10, 10}, {30, 100}}, -1}, {"combine passed stations", 100, 30, [][]int{{10, 40}, {20, 30}}, 2}, {"largest reachable", 100, 50, [][]int{{10, 50}, {40, 10}}, 1}, {"large fuel", 1000000000, 1, [][]int{{1, 1000000000}}, 1},
	} {
		t.Run(tt.name, func(t *testing.T) {
			if got := minRefuelStops(tt.target, tt.fuel, tt.stations); got != tt.want {
				t.Errorf("%s: got %d want %d", t.Name(), got, tt.want)
			}
		})
	}
}
