package l309

import "math"

type E struct {
	holdStock int // 第 j 天结束时，持有股票的最大收益
	noStock   int // 第 j 天结束时，不持有股票的最大收益 (noStock 可能包含当天刚卖出股票的情况，所以不能直接用它在第二天买入。)
}

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
	if len(dp) <= 1 {
		return 0
	}
	dp[1] = E{
		holdStock: max(-prices[1], dp[0].holdStock),
		noStock:   max(prices[1]+dp[0].holdStock, dp[0].noStock),
	}

	for j := 2; j < n; j++ {
		le := dp[j-1]
		e := E{
			holdStock: max(dp[j-2].noStock-prices[j], le.holdStock),
			noStock:   max(le.holdStock+prices[j], le.noStock),
		}
		dp[j] = e
	}

	return dp[n-1].noStock
}

// https://leetcode.com/problems/best-time-to-buy-and-sell-stock-with-cooldown/
