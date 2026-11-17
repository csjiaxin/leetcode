package l70

// PrefixTable 的 dp[0] 是空前缀；step(i, dp[:i]) 只能读取已完成状态。
func PrefixTable[T any](n int, empty T, step func(i int, previous []T) T) []T {
	dp := make([]T, n+1)
	dp[0] = empty
	for i := 1; i <= n; i++ {
		dp[i] = step(i, dp[:i])
	}
	return dp
}

// 空阶梯有一种走法；每次来自前一或前两级。
func climbStairs(n int) int {
	dp := PrefixTable(n, 1, func(i int, old []int) int {
		ways := old[i-1]
		if i >= 2 {
			ways += old[i-2]
		}
		return ways
	})
	return dp[n]
}
