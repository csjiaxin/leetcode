package l312

// https://leetcode.com/problems/burst-balloons/
func maxCoins(nums []int) int {
	values := make([]int, len(nums)+2)
	values[0], values[len(values)-1] = 1, 1
	copy(values[1:], nums)
	n := len(values)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	for gap := 2; gap < n; gap++ {
		for i := 0; i+gap < n; i++ {
			j := i + gap
			// Keep i and j intact; choose the last balloon inside (i, j).
			for last := i + 1; last < j; last++ {
				coins := dp[i][last] + dp[last][j] + values[i]*values[last]*values[j]
				dp[i][j] = max(dp[i][j], coins)
			}
		}
	}
	return dp[0][n-1]
}
