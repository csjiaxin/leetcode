package l45

import "math"

// https://leetcode.com/problems/jump-game-ii/
func jump(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	for j := 0; j < n; j++ {
		dp[j] = math.MaxInt
	}
	dp[n-1] = 0
	for j := n - 2; j >= 0; j-- {
		for k := 1; k <= nums[j] && j+k < n; k++ {
			if dp[j+k] != math.MaxInt {
				dp[j] = min(dp[j], dp[j+k]+1)
			}
		}
	}
	return dp[0]
}
