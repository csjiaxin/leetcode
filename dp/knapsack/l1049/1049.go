package l1049

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

// https://leetcode.com/problems/last-stone-weight-ii/
// 求不超过一半的最大可达子集和，剩余差值为 total-2*chosen。
func lastStoneWeightII(stones []int) int {
	total := 0
	for _, value := range stones {
		total += value
	}
	dp := make([]bool, total/2+1)
	dp[0] = true
	Capacity(dp, stones, false, func(old, from bool) bool { return old || from })
	for chosen := total / 2; chosen >= 0; chosen-- {
		if dp[chosen] {
			return total - 2*chosen
		}
	}
	return 0
}
