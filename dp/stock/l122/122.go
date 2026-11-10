package l122

import "math"

type E struct {
	holdStock int
	noStock   int
}

// https://leetcode.com/problems/best-time-to-buy-and-sell-stock/
func maxProfit(prices []int) int {
	n := len(prices)

	dp := make([]E, n)
	for i := 0; i < n; i++ {
		dp[i].holdStock = math.MinInt
		dp[i].noStock = math.MinInt
	}

	dp[0] = E{
		holdStock: -prices[0],
		noStock:   0,
	}

	for j := 1; j < n; j++ {
		le := dp[j-1]
		e := E{
			holdStock: max(le.noStock-prices[j], le.holdStock),
			noStock:   max(le.holdStock+prices[j], le.noStock),
		}
		dp[j] = e
	}

	return dp[n-1].noStock
}
