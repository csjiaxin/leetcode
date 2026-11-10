package l1014

import "testing"

func TestMaxScoreSightseeingPair(t *testing.T) {
	tests := []struct {
		name   string
		values []int
		want   int
	}{
		{name: "example one", values: []int{8, 1, 5, 2, 6}, want: 11},
		{name: "example two", values: []int{1, 2}, want: 2},
		{name: "minimum values", values: []int{1, 1}, want: 1},
		{name: "equal values", values: []int{5, 5, 5, 5}, want: 9},
		{name: "increasing", values: []int{1, 2, 3, 4, 5}, want: 8},
		{name: "decreasing", values: []int{5, 4, 3, 2, 1}, want: 8},
		{name: "distant best pair", values: []int{10, 1, 1, 1, 10}, want: 16},
		{name: "best pair at end", values: []int{1, 1, 10, 10}, want: 19},
		{name: "cannot reuse same index", values: []int{1, 1000}, want: 1000},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxScoreSightseeingPair(tt.values); got != tt.want {
				t.Errorf("maxScoreSightseeingPair(%v) = %d, want %d", tt.values, got, tt.want)
			}
		})
	}
}

func TestMaxScoreSightseeingPairLargeInput(t *testing.T) {
	values := make([]int, 50000)
	for i := range values {
		values[i] = 1000
	}
	// Adjacent indices have the smallest possible penalty.
	if got := maxScoreSightseeingPair(values); got != 1999 {
		t.Errorf("maxScoreSightseeingPair(50000 values of 1000) = %d, want 1999", got)
	}
}
