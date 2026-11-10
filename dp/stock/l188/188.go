package l188

import "math"

type E struct {
	holdStock []int // Holding, using at most the first purchase.
	noStock   []int // Not holding, after at most one transaction.
}

// https://leetcode.com/problems/best-time-to-buy-and-sell-stock-iv/
func maxProfit(k int, prices []int) int {
	n := len(prices)

	dp := make([]E, n)
	for i := 0; i < n; i++ {
		dp[i].holdStock = make([]int, k)
		dp[i].noStock = make([]int, k)
		for j := 0; j < k; j++ {
			dp[i].holdStock[j] = math.MinInt
			dp[i].noStock[j] = math.MinInt
		}
	}

	for j := 0; j < k; j++ {
		dp[0].holdStock[j] = -prices[0]
		dp[0].noStock[j] = 0
	}

	for j := 1; j < n; j++ {
		dp[j].holdStock[0] = max(-prices[j], dp[j-1].holdStock[0])
		dp[j].noStock[0] = max(dp[j-1].holdStock[0]+prices[j], dp[j-1].noStock[0])

		for c := 1; c < k; c++ {
			dp[j].holdStock[c] = max(dp[j-1].noStock[c-1]-prices[j], dp[j-1].holdStock[c])
			dp[j].noStock[c] = max(dp[j-1].holdStock[c]+prices[j], dp[j-1].noStock[c])

		}
	}

	return dp[n-1].noStock[k-1]
}
