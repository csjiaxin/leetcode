package l518

import "testing"

func TestChange(t *testing.T) {
	tests := []struct {
		name   string
		amount int
		coins  []int
		want   int
	}{
		{name: "example one", amount: 5, coins: []int{1, 2, 5}, want: 4},
		{name: "example two", amount: 3, coins: []int{2}, want: 0},
		{name: "example three", amount: 10, coins: []int{10}, want: 1},
		{name: "zero amount", amount: 0, coins: []int{1, 2, 5}, want: 1},
		{name: "zero amount without unit coin", amount: 0, coins: []int{3, 7}, want: 1},
		{name: "unit coin only", amount: 7, coins: []int{1}, want: 1},
		{name: "reuse a coin", amount: 9, coins: []int{3}, want: 1},
		{name: "all coins exceed amount", amount: 3, coins: []int{4, 5}, want: 0},
		{name: "unreachable despite smaller coins", amount: 7, coins: []int{2, 4}, want: 0},
		{name: "coprime coins but unreachable amount", amount: 4, coins: []int{3, 5}, want: 0},
		{name: "combinations not permutations", amount: 3, coins: []int{1, 2}, want: 2},
		{name: "count all combinations not just minimum coins", amount: 4, coins: []int{1, 2, 3}, want: 4},
		{name: "unsorted denominations", amount: 5, coins: []int{5, 1, 2}, want: 4},
		{name: "no unit coin", amount: 10, coins: []int{2, 5}, want: 2},
		{name: "multiple nonunit combinations", amount: 12, coins: []int{2, 3, 7}, want: 4},
		{name: "maximum amount with unit coin", amount: 5000, coins: []int{1}, want: 1},
		{name: "maximum coin denomination", amount: 5000, coins: []int{5000}, want: 1},
		{name: "maximum amount with two coins", amount: 5000, coins: []int{1, 2}, want: 2501},
		{name: "maximum amount unreachable", amount: 5000, coins: []int{3}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			coins := append([]int(nil), tt.coins...)
			if got := change(tt.amount, coins); got != tt.want {
				t.Errorf("change(%d, %v) = %d, want %d", tt.amount, tt.coins, got, tt.want)
			}
		})
	}
}

func TestChangeMaximumDenominationCount(t *testing.T) {
	coins := make([]int, 300)
	for i := range coins {
		coins[i] = i + 1
	}
	if got := change(5, coins); got != 7 {
		t.Errorf("change(5, denominations 1 through 300) = %d, want 7", got)
	}
}
