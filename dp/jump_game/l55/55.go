package l55

// https://leetcode.com/problems/jump-game/
func canJump(nums []int) bool {
	n := len(nums)
	dp := make([]bool, n)
	dp[n-1] = true
	for j := n - 2; j >= 0; j-- {
		for k := 1; k <= nums[j] && j+k < n && !dp[j]; k++ {
			dp[j] = dp[j] || dp[j+k]
		}
	}
	return dp[0]
}
