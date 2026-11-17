package l322

import "testing"

func TestCoinChange(t *testing.T) {
	tests := []struct {
		name   string
		coins  []int
		amount int
		want   int
	}{
		// {name: "example one", coins: []int{1, 2, 5}, amount: 11, want: 3},
		{name: "example two", coins: []int{2}, amount: 3, want: -1},
		{name: "example three", coins: []int{1}, amount: 0, want: 0},
		{name: "zero amount without unit coin", coins: []int{2, 5}, amount: 0, want: 0},
		{name: "exact denomination", coins: []int{2, 5, 7}, amount: 7, want: 1},
		{name: "reuse a coin", coins: []int{3}, amount: 9, want: 3},
		{name: "unit coin only", coins: []int{1}, amount: 8, want: 8},
		{name: "all coins exceed amount", coins: []int{5, 7}, amount: 3, want: -1},
		{name: "unreachable despite smaller coin", coins: []int{4, 6}, amount: 7, want: -1},
		{name: "coprime coins but unreachable amount", coins: []int{3, 5}, amount: 4, want: -1},
		{name: "greedy is not optimal", coins: []int{1, 3, 4}, amount: 6, want: 2},
		{name: "greedy can get stuck", coins: []int{3, 4}, amount: 6, want: 2},
		{name: "unsorted denominations", coins: []int{5, 1, 2}, amount: 11, want: 3},
		{name: "mixed denominations", coins: []int{2, 5, 10}, amount: 17, want: 3},
		{name: "larger nongreedy example", coins: []int{186, 419, 83, 408}, amount: 6249, want: 20},
		{name: "maximum coin value", coins: []int{2147483647}, amount: 1, want: -1},
		{name: "large coin alongside usable coins", coins: []int{2147483647, 2, 3}, amount: 7, want: 3},
		{name: "maximum amount with unit coin", coins: []int{1}, amount: 10000, want: 10000},
		{name: "maximum amount with multiple coins", coins: []int{1, 2, 5}, amount: 10000, want: 2000},
		{name: "maximum amount unreachable", coins: []int{3}, amount: 10000, want: -1},
		{name: "maximum denomination count", coins: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}, amount: 24, want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			coins := append([]int(nil), tt.coins...)
			if got := coinChange(coins, tt.amount); got != tt.want {
				t.Errorf("coinChange(%v, %d) = %d, want %d", tt.coins, tt.amount, got, tt.want)
			}
		})
	}
}
