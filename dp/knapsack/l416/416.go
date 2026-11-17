package l416

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

// https://leetcode.com/problems/partition-equal-subset-sum/
// 0/1 可达性：每个元素最多用一次，查询总和的一半。
func canPartition(nums []int) bool {
	total := 0
	for _, value := range nums {
		total += value
	}
	if total%2 != 0 {
		return false
	}
	dp := make([]bool, total/2+1)
	dp[0] = true
	Capacity(dp, nums, false, func(old, from bool) bool { return old || from })
	return dp[total/2]
}
