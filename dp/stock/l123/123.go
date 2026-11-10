package l123

import "math"

type E struct {
	holdStock  int // Holding, using at most the first purchase.
	noStock    int // Not holding, after at most one transaction.
	holdStock2 int // Holding, using at most the second purchase.
	noStock2   int // Not holding, after at most two transactions.
}

// https://leetcode.com/problems/best-time-to-buy-and-sell-stock/
func maxProfit(prices []int) int {
	n := len(prices)

	dp := make([]E, n)
	for i := 0; i < n; i++ {
		dp[i].holdStock = math.MinInt
		dp[i].noStock = math.MinInt
		dp[i].holdStock2 = math.MinInt
		dp[i].noStock2 = math.MinInt
	}

	dp[0] = E{
		holdStock:  -prices[0],
		noStock:    0,
		holdStock2: -prices[0],
		noStock2:   0,
	}

	for j := 1; j < n; j++ {
		le := dp[j-1]
		e := E{
			holdStock:  max(-prices[j], le.holdStock),
			noStock:    max(le.holdStock+prices[j], le.noStock),
			holdStock2: max(le.noStock-prices[j], le.holdStock2),
			noStock2:   max(le.holdStock2+prices[j], le.noStock2),
		}
		dp[j] = e
	}

	return dp[n-1].noStock2
}
