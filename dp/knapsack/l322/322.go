package l322

import "math"

// https://leetcode.com/problems/coin-change/
func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 0; i <= amount; i++ {
		dp[i] = math.MaxInt
	}
	dp[0] = 0
	for _, c := range coins {
		for v := 0; v+c <= amount; v++ {
			newV := v + c
			if dp[v] != math.MaxInt {
				dp[newV] = min(dp[newV], dp[v]+1)
			}
		}
	}
	if dp[amount] == math.MaxInt {
		return -1
	}
	return dp[amount]
}
