package l375

// https://leetcode.com/problems/guess-number-higher-or-lower-ii/
func getMoneyAmount(n int) int {
	dp := make([][]int, n+2)
	for i := range dp {
		dp[i] = make([]int, n+2)
	}
	for length := 2; length <= n; length++ {
		for i := 1; i+length-1 <= n; i++ {
			j := i + length - 1
			best := int(^uint(0) >> 1)
			for guess := i; guess <= j; guess++ {
				cost := guess + max(dp[i][guess-1], dp[guess+1][j])
				best = min(best, cost)
			}
			dp[i][j] = best
		}
	}
	return dp[1][n]
}
