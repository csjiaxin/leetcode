package l494

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

// https://leetcode.com/problems/target-sum/
// P-N=target、P+N=total，所以 P=(total+target)/2；零元素使计数翻倍。
func findTargetSumWays(nums []int, target int) int {
	total := 0
	for _, value := range nums {
		total += value
	}
	if target < -total || target > total || (total+target)%2 != 0 {
		return 0
	}
	chosen := (total + target) / 2
	dp := make([]int, chosen+1)
	dp[0] = 1
	Capacity(dp, nums, false, func(old, from int) int { return old + from })
	return dp[chosen]
}
