package l122

import "testing"

func TestMaxProfit(t *testing.T) {
	tests := []struct {
		name   string
		prices []int
		want   int
	}{
		{name: "example one", prices: []int{7, 1, 5, 3, 6, 4}, want: 7},
		{name: "example two", prices: []int{1, 2, 3, 4, 5}, want: 4},
		{name: "example three", prices: []int{7, 6, 4, 3, 1}, want: 0},
		{name: "single day", prices: []int{5}, want: 0},
		{name: "single zero price", prices: []int{0}, want: 0},
		{name: "equal prices", prices: []int{4, 4, 4, 4}, want: 0},
		{name: "two increasing prices", prices: []int{1, 9}, want: 8},
		{name: "two decreasing prices", prices: []int{9, 1}, want: 0},
		{name: "multiple transactions", prices: []int{1, 5, 2, 8}, want: 10},
		{name: "alternating prices", prices: []int{1, 3, 1, 3, 1, 3}, want: 6},
		{name: "plateaus", prices: []int{2, 2, 5, 5, 1, 1, 4, 4}, want: 6},
		{name: "zero price between peaks", prices: []int{0, 10000, 0, 10000}, want: 20000},
		{name: "ignore initial high price", prices: []int{10, 1, 4, 2, 7}, want: 8},
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
	prices := make([]int, 30000)
	for i := range prices {
		if i%2 == 1 {
			prices[i] = 10000
		}
	}
	const want = 150000000
	if got := maxProfit(prices); got != want {
		t.Errorf("maxProfit(30000 alternating prices) = %d, want %d", got, want)
	}
}
