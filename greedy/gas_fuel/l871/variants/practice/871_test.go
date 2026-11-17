package l871

import "testing"

func TestMinRefuelStops(t *testing.T) {
	for _, tt := range []struct {
		name         string
		target, fuel int
		stations     [][]int
		want         int
	}{
		{"already enough", 1, 1, nil, 0},
		{"first station unreachable", 100, 1, [][]int{{10, 100}}, -1},
		{"example", 100, 10, [][]int{{10, 60}, {20, 30}, {30, 30}, {60, 40}}, 2},
		{"no stations insufficient", 100, 50, nil, -1},
		{"exact reach station", 100, 50, [][]int{{50, 50}}, 1},
		{"gap after refill", 100, 10, [][]int{{10, 10}, {30, 100}}, -1},
		{"largest not farthest", 100, 50, [][]int{{10, 50}, {40, 10}}, 1},
		{"combine earlier stations", 100, 30, [][]int{{10, 40}, {20, 30}}, 2},
		{"use all stations", 40, 10, [][]int{{10, 10}, {20, 10}, {30, 10}}, 3},
		{"large fuel", 1000000000, 1, [][]int{{1, 1000000000}}, 1},
	} {
		t.Run(tt.name, func(t *testing.T) {
			stations := make([][]int, len(tt.stations))
			for i, v := range tt.stations {
				stations[i] = append([]int(nil), v...)
			}
			if got := minRefuelStops(tt.target, tt.fuel, stations); got != tt.want {
				t.Errorf("minRefuelStops(%d, %d, %v) = %d, want %d", tt.target, tt.fuel, tt.stations, got, tt.want)
			}
		})
	}
}
