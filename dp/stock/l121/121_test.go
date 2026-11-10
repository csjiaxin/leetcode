package l121

import "testing"

func TestMaxProfit(t *testing.T) {
	tests := []struct {
		name   string
		prices []int
		want   int
	}{
		{name: "example one", prices: []int{7, 1, 5, 3, 6, 4}, want: 5},
		{name: "example two", prices: []int{7, 6, 4, 3, 1}, want: 0},
		{name: "single day", prices: []int{5}, want: 0},
		{name: "single zero price", prices: []int{0}, want: 0},
		{name: "two increasing prices", prices: []int{1, 9}, want: 8},
		{name: "two decreasing prices", prices: []int{9, 1}, want: 0},
		{name: "equal prices", prices: []int{4, 4, 4}, want: 0},
		{name: "increasing prices", prices: []int{1, 2, 3, 4, 5}, want: 4},
		{name: "lowest price after peak", prices: []int{2, 8, 1}, want: 6},
		{name: "later buying opportunity", prices: []int{3, 5, 1, 7}, want: 6},
		{name: "one transaction only", prices: []int{1, 5, 2, 8}, want: 7},
		{name: "maximum profit", prices: []int{0, 10000}, want: 10000},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfit(tt.prices); got != tt.want {
				t.Errorf("maxProfit(%v) = %d, want %d", tt.prices, got, tt.want)
			}
		})
	}
}

func TestMaxProfitLargeInput(t *testing.T) {
	prices := make([]int, 100000)
	for i := range prices {
		prices[i] = 10000
	}
	prices[0] = 0
	if got := maxProfit(prices); got != 10000 {
		t.Errorf("maxProfit(100000 prices) = %d, want 10000", got)
	}
}
