package l871

import "testing"

func TestMinRefuelStops(t *testing.T) {
	tests := []struct {
		name      string
		target    int
		startFuel int
		stations  [][]int
		want      int
	}{
		{name: "example one", target: 1, startFuel: 1, stations: [][]int{}, want: 0},
		{name: "example two", target: 100, startFuel: 1, stations: [][]int{{10, 100}}, want: -1},
		{name: "example three", target: 100, startFuel: 10, stations: [][]int{{10, 60}, {20, 30}, {30, 30}, {60, 40}}, want: 2},
		{name: "no stations insufficient fuel", target: 100, startFuel: 50, stations: [][]int{}, want: -1},
		{name: "initial fuel exceeds target", target: 10, startFuel: 20, stations: [][]int{{5, 100}}, want: 0},
		{name: "enough initial fuel with stations", target: 100, startFuel: 100, stations: [][]int{{10, 50}, {50, 50}}, want: 0},
		{name: "reach station with empty tank", target: 100, startFuel: 50, stations: [][]int{{50, 50}}, want: 1},
		{name: "reach target with empty tank", target: 100, startFuel: 25, stations: [][]int{{20, 75}}, want: 1},
		{name: "station just out of reach", target: 100, startFuel: 49, stations: [][]int{{50, 100}}, want: -1},
		{name: "gap after reachable station", target: 100, startFuel: 10, stations: [][]int{{10, 10}, {30, 100}}, want: -1},
		{name: "total fuel enough but inaccessible", target: 100, startFuel: 10, stations: [][]int{{20, 100}, {50, 100}}, want: -1},
		{name: "must use every station", target: 40, startFuel: 10, stations: [][]int{{10, 10}, {20, 10}, {30, 10}}, want: 3},
		{name: "skip smaller reachable station", target: 100, startFuel: 50, stations: [][]int{{10, 10}, {40, 50}}, want: 1},
		{name: "best station is not the farthest reachable", target: 100, startFuel: 50, stations: [][]int{{10, 50}, {40, 10}}, want: 1},
		{name: "combine previously passed stations", target: 100, startFuel: 30, stations: [][]int{{10, 40}, {20, 30}}, want: 2},
		{name: "large early refill skips later stations", target: 100, startFuel: 10, stations: [][]int{{10, 90}, {20, 1}, {30, 1}}, want: 1},
		{name: "equal station fuel amounts", target: 100, startFuel: 25, stations: [][]int{{25, 25}, {50, 25}, {75, 25}}, want: 3},
		{name: "maximum target without stops", target: 1000000000, startFuel: 1000000000, stations: [][]int{}, want: 0},
		{name: "maximum station fuel", target: 1000000000, startFuel: 1, stations: [][]int{{1, 1000000000}}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stations := make([][]int, len(tt.stations))
			for i := range tt.stations {
				stations[i] = append([]int(nil), tt.stations[i]...)
			}
			if got := minRefuelStops(tt.target, tt.startFuel, stations); got != tt.want {
				t.Errorf("minRefuelStops(%d, %d, %v) = %d, want %d", tt.target, tt.startFuel, tt.stations, got, tt.want)
			}
		})
	}
}

func TestMinRefuelStopsMaximumStations(t *testing.T) {
	stations := make([][]int, 500)
	for i := range stations {
		stations[i] = []int{i + 1, 1}
	}
	if got := minRefuelStops(501, 1, stations); got != 500 {
		t.Errorf("minRefuelStops(501, 1, 500 unit-fuel stations) = %d, want 500", got)
	}
}
