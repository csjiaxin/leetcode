package l134

import "testing"

func TestCanCompleteCircuit(t *testing.T) {
	tests := []struct {
		name string
		gas  []int
		cost []int
		want int
	}{
		{name: "example one", gas: []int{1, 2, 3, 4, 5}, cost: []int{3, 4, 5, 1, 2}, want: 3},
		{name: "example two", gas: []int{2, 3, 4}, cost: []int{3, 4, 3}, want: -1},
		{name: "single station exact fuel", gas: []int{5}, cost: []int{5}, want: 0},
		{name: "single station surplus", gas: []int{6}, cost: []int{5}, want: 0},
		{name: "single station insufficient fuel", gas: []int{4}, cost: []int{5}, want: -1},
		{name: "single station zero fuel and cost", gas: []int{0}, cost: []int{0}, want: 0},
		{name: "start at first station", gas: []int{3, 0, 0}, cost: []int{1, 1, 1}, want: 0},
		{name: "start at last station", gas: []int{0, 0, 3}, cost: []int{1, 1, 1}, want: 2},
		{name: "two stations wrap around", gas: []int{0, 2}, cost: []int{1, 1}, want: 1},
		{name: "surplus with unique start", gas: []int{0, 4, 0}, cost: []int{1, 1, 1}, want: 1},
		{name: "must finish wraparound", gas: []int{0, 0, 4, 0}, cost: []int{1, 1, 1, 1}, want: 2},
		{name: "largest gas station is not the start", gas: []int{5, 1, 4}, cost: []int{6, 2, 2}, want: 2},
		{name: "multiple failed candidate starts", gas: []int{1, 2, 3, 10}, cost: []int{2, 3, 4, 7}, want: 3},
		{name: "temporary zero balance before failure", gas: []int{1, 0, 3}, cost: []int{1, 2, 1}, want: 2},
		{name: "no fuel anywhere", gas: []int{0, 0, 0}, cost: []int{0, 1, 0}, want: -1},
		{name: "local surplus but insufficient total", gas: []int{5, 0, 0}, cost: []int{1, 3, 3}, want: -1},
		{name: "maximum station values", gas: []int{0, 10000}, cost: []int{10000, 0}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gas := append([]int(nil), tt.gas...)
			cost := append([]int(nil), tt.cost...)
			if got := canCompleteCircuit(gas, cost); got != tt.want {
				t.Errorf("canCompleteCircuit(%v, %v) = %d, want %d", tt.gas, tt.cost, got, tt.want)
			}
		})
	}
}

func TestCanCompleteCircuitMaximumLength(t *testing.T) {
	const n = 100000
	t.Run("unique start at last station", func(t *testing.T) {
		gas := make([]int, n)
		cost := make([]int, n)
		for i := range gas {
			gas[i] = 1
			cost[i] = 1
		}
		gas[n-2] = 0
		gas[n-1] = 2
		if got := canCompleteCircuit(gas, cost); got != n-1 {
			t.Errorf("canCompleteCircuit(maximum length) = %d, want %d", got, n-1)
		}
	})
	t.Run("insufficient total with large values", func(t *testing.T) {
		gas := make([]int, n)
		cost := make([]int, n)
		for i := range gas {
			gas[i] = 10000
			cost[i] = 10000
		}
		gas[n-1]--
		if got := canCompleteCircuit(gas, cost); got != -1 {
			t.Errorf("canCompleteCircuit(insufficient maximum length input) = %d, want -1", got)
		}
	})
}
