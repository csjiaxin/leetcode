package l123

import "testing"

func TestMaxProfit(t *testing.T) {
	tests := []struct {
		name   string
		prices []int
		want   int
	}{
		{name: "example one", prices: []int{3, 3, 5, 0, 0, 3, 1, 4}, want: 6},
		{name: "example two", prices: []int{1, 2, 3, 4, 5}, want: 4},
		{name: "example three", prices: []int{7, 6, 4, 3, 1}, want: 0},
		{name: "single day", prices: []int{5}, want: 0},
		{name: "single zero price", prices: []int{0}, want: 0},
		{name: "equal prices", prices: []int{4, 4, 4, 4}, want: 0},
		{name: "two increasing prices", prices: []int{1, 9}, want: 8},
		{name: "two decreasing prices", prices: []int{9, 1}, want: 0},
		{name: "two transactions beat one", prices: []int{1, 5, 2, 8}, want: 10},
		{name: "at most two transactions", prices: []int{1, 3, 1, 3, 1, 3}, want: 4},
		{name: "skip smaller middle transaction", prices: []int{1, 6, 2, 3, 1, 8}, want: 12},
		{name: "small dip allows second transaction", prices: []int{1, 5, 4, 10}, want: 10},
		{name: "global minimum after useful sale", prices: []int{2, 8, 0, 5}, want: 11},
		{name: "maximum prices", prices: []int{0, 100000, 0, 100000}, want: 200000},
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
		if i%2 == 1 {
			prices[i] = 100000
		}
	}
	const want = 200000
	if got := maxProfit(prices); got != want {
		t.Errorf("maxProfit(100000 alternating prices) = %d, want %d", got, want)
	}
}
