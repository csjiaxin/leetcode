package l188

import "testing"

func TestMaxProfit(t *testing.T) {
	tests := []struct {
		name   string
		k      int
		prices []int
		want   int
	}{
		{name: "example one", k: 2, prices: []int{2, 4, 1}, want: 2},
		{name: "example two", k: 2, prices: []int{3, 2, 6, 5, 0, 3}, want: 7},
		{name: "single day", k: 2, prices: []int{5}, want: 0},
		{name: "single zero price", k: 1, prices: []int{0}, want: 0},
		{name: "equal prices", k: 3, prices: []int{4, 4, 4, 4}, want: 0},
		{name: "decreasing prices", k: 2, prices: []int{7, 6, 4, 3, 1}, want: 0},
		{name: "increasing prices", k: 3, prices: []int{1, 2, 3, 4, 5}, want: 4},
		{name: "one transaction only", k: 1, prices: []int{1, 5, 2, 8}, want: 7},
		{name: "two transactions", k: 2, prices: []int{1, 5, 2, 8}, want: 10},
		{name: "limit excludes third transaction", k: 2, prices: []int{1, 3, 1, 3, 1, 3}, want: 4},
		{name: "three transactions", k: 3, prices: []int{1, 3, 1, 3, 1, 3}, want: 6},
		{name: "skip smaller middle transaction", k: 2, prices: []int{1, 6, 2, 3, 1, 8}, want: 12},
		{name: "more transactions than possible", k: 100, prices: []int{1, 5, 2, 8}, want: 10},
		{name: "two days maximum prices", k: 100, prices: []int{0, 1000}, want: 1000},
		{name: "one transaction spans small dips", k: 1, prices: []int{1, 5, 4, 10}, want: 9},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfit(tt.k, tt.prices); got != tt.want {
				t.Errorf("maxProfit(%d, %v) = %d, want %d", tt.k, tt.prices, got, tt.want)
			}
		})
	}
}

func TestMaxProfitLargeInput(t *testing.T) {
	prices := make([]int, 1000)
	for i := range prices {
		if i%2 == 1 {
			prices[i] = 1000
		}
	}
	const k = 100
	const want = 100000
	if got := maxProfit(k, prices); got != want {
		t.Errorf("maxProfit(%d, 1000 alternating prices) = %d, want %d", k, got, want)
	}
}
