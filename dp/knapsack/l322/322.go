package l322

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

// https://leetcode.com/problems/coin-change/
// 完全背包求最少个数；不可达状态不能继续加一。
func coinChange(coins []int, amount int) int {
	inf := amount + 1
	dp := make([]int, amount+1)
	for c := 1; c <= amount; c++ {
		dp[c] = inf
	}
	Capacity(dp, coins, true, func(old, from int) int {
		if from == inf {
			return old
		}
		return min(old, from+1)
	})
	if dp[amount] == inf {
		return -1
	}
	return dp[amount]
}
