package main

// PrefixTable 的 dp[0] 是空前缀；step(i, dp[:i]) 只能读取已完成状态。
func PrefixTable[T any](n int, empty T, step func(i int, previous []T) T) []T {
	dp := make([]T, n+1)
	dp[0] = empty
	for i := 1; i <= n; i++ {
		dp[i] = step(i, dp[:i])
	}
	return dp
}

// 至少拆成两项；剩余部分可以不继续拆，也可以使用旧状态。
func integerBreak(n int) int {
	dp := PrefixTable(n, 0, func(i int, old []int) int {
		best := 0
		for first := 1; first < i; first++ {
			best = max(best, first*max(i-first, old[i-first]))
		}
		return best
	})
	return dp[n]
}
