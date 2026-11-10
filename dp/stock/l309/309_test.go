package l309

import "testing"

func TestMaxProfit(t *testing.T) {
	tests := []struct {
		name   string
		prices []int
		want   int
	}{
		{name: "example one", prices: []int{1, 2, 3, 0, 2}, want: 3},
		{name: "example two", prices: []int{1}, want: 0},
		{name: "single zero price", prices: []int{0}, want: 0},
		{name: "two increasing prices", prices: []int{1, 9}, want: 8},
		{name: "two decreasing prices", prices: []int{9, 1}, want: 0},
		{name: "equal prices", prices: []int{4, 4, 4, 4}, want: 0},
		{name: "increasing prices", prices: []int{1, 2, 3, 4, 5}, want: 4},
		{name: "decreasing prices", prices: []int{7, 6, 4, 3, 1}, want: 0},
		{name: "cannot buy immediately after sale", prices: []int{1, 5, 2, 8}, want: 7},
		{name: "buy after exactly one cooldown day", prices: []int{1, 5, 9, 1, 6}, want: 9},
		{name: "skip early trade for lower buy price", prices: []int{1, 2, 0, 4}, want: 4},
		{name: "repeated opportunities limited by cooldown", prices: []int{1, 3, 1, 3, 1, 3}, want: 4},
		{name: "last day sale", prices: []int{3, 1, 4}, want: 3},
		{name: "maximum price difference", prices: []int{0, 1000}, want: 1000},
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
	prices := make([]int, 5000)
	// Each block allows buying, selling, then waiting for one day.
	for i := range prices {
		if i%3 == 1 {
			prices[i] = 1000
		}
	}
	const want = 1667000
	if got := maxProfit(prices); got != want {
		t.Errorf("maxProfit(5000 prices with cooldown opportunities) = %d, want %d", got, want)
	}
}
