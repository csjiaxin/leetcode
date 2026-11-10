package l714

import "testing"

func TestMaxProfit(t *testing.T) {
	tests := []struct {
		name   string
		prices []int
		fee    int
		want   int
	}{
		{name: "example one", prices: []int{1, 3, 2, 8, 4, 9}, fee: 2, want: 8},
		{name: "example two", prices: []int{1, 3, 7, 5, 10, 3}, fee: 3, want: 6},
		{name: "single day", prices: []int{5}, fee: 2, want: 0},
		{name: "equal prices", prices: []int{4, 4, 4}, fee: 1, want: 0},
		{name: "decreasing prices", prices: []int{9, 7, 4, 1}, fee: 2, want: 0},
		{name: "zero fee", prices: []int{1, 5, 2, 8}, fee: 0, want: 10},
		{name: "gain equals fee", prices: []int{1, 4}, fee: 3, want: 0},
		{name: "fee exceeds gain", prices: []int{1, 4}, fee: 5, want: 0},
		{name: "charge fee once per transaction", prices: []int{1, 10}, fee: 2, want: 7},
		{name: "hold through small dip", prices: []int{1, 5, 4, 10}, fee: 2, want: 7},
		{name: "two transactions beat holding", prices: []int{1, 8, 2, 10}, fee: 2, want: 11},
		{name: "fee erases repeated gains", prices: []int{1, 3, 1, 3, 1, 3}, fee: 2, want: 0},
		{name: "maximum price difference", prices: []int{1, 49999}, fee: 0, want: 49998},
		{name: "maximum fee", prices: []int{1, 49999}, fee: 49999, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfit(tt.prices, tt.fee); got != tt.want {
				t.Errorf("maxProfit(%v, %d) = %d, want %d", tt.prices, tt.fee, got, tt.want)
			}
		})
	}
}

func TestMaxProfitLargeInput(t *testing.T) {
	prices := make([]int, 50000)
	for i := range prices {
		prices[i] = 1
		if i%2 == 1 {
			prices[i] = 49999
		}
	}
	const fee = 2
	const want = 1249900000
	if got := maxProfit(prices, fee); got != want {
		t.Errorf("maxProfit(50000 alternating prices, %d) = %d, want %d", fee, got, want)
	}
}
