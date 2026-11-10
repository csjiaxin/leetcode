package l416

// https://leetcode.com/problems/partition-equal-subset-sum/
func canPartition(nums []int) bool {
	sum := 0
	N := len(nums)

	for _, n := range nums {
		sum += n
	}
	if sum%2 != 0 {
		return false
	}
	half := sum / 2
	dp := make([]bool, half+1)
	dp[0] = true
	for i := 0; i < N; i++ {
		// for l := 0; l < half; l++ {
		for l := half; l >= nums[i]; l-- {
			dp[l] = dp[l] || dp[l-nums[i]]
		}
	}

	return dp[half]
}
