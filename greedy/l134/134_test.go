package l134

import (
	"slices"
	"testing"
)

func TestCanCompleteCircuit(t *testing.T) {
	for _, tt := range []struct {
		name      string
		gas, cost []int
		want      int
	}{
		{"example one", []int{1, 2, 3, 4, 5}, []int{3, 4, 5, 1, 2}, 3},
		{"example impossible", []int{2, 3, 4}, []int{3, 4, 3}, -1},
		{"single exact", []int{5}, []int{5}, 0},
		{"single deficit", []int{4}, []int{5}, -1},
		{"single zero", []int{0}, []int{0}, 0},
		{"last start wraparound", []int{0, 0, 3}, []int{1, 1, 1}, 2},
		{"largest gas not start", []int{5, 1, 4}, []int{6, 2, 2}, 2},
		{"repeated failed starts", []int{1, 2, 3, 10}, []int{2, 3, 4, 7}, 3},
		{"total deficit", []int{5, 0, 0}, []int{1, 3, 3}, -1},
	} {
		t.Run(tt.name, func(t *testing.T) {
			if got := canCompleteCircuit(slices.Clone(tt.gas), slices.Clone(tt.cost)); got != tt.want {
				t.Errorf("canCompleteCircuit(%v, %v) = %d, want %d", tt.gas, tt.cost, got, tt.want)
			}
		})
	}
}
