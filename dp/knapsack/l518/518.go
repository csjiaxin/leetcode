package l518

// Capacity 修改调用方提供的初始表；reuse=false 倒序，true 正序。
// 0/1 允许零重量（每项仍只处理一次）；可重复使用的重量必须为正。
func Capacity[T any](dp []T, weights []int, reuse bool, combine func(old, from T) T) []T {
	for _, weight := range weights {
		if reuse {
			for c := weight; c < len(dp); c++ {
				dp[c] = combine(dp[c], dp[c-weight])
			}
		} else {
			for c := len(dp) - 1; c >= weight; c-- {
				dp[c] = combine(dp[c], dp[c-weight])
			}
		}
	}
	return dp
}

// https://leetcode.com/problems/coin-change-ii/
// 硬币在外层、金额正序，统计组合而不是选币排列。
func change(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1
	Capacity(dp, coins, true, func(old, from int) int { return old + from })
	return dp[amount]
}
