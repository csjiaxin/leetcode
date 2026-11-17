package l70

// climbStairs returns the number of distinct ways to climb n stairs,
// taking either one or two steps at a time.
func climbStairs(n int) int {
	dp := make([]int, n)
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	dp[0] = 1
	dp[1] = 2
	for j := 2; j < n; j++ {
		dp[j] = dp[j-1] + dp[j-2]
	}
	return dp[n-1]
}
