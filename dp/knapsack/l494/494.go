package l494

// https://leetcode.com/problems/target-sum/
func findTargetSumWays(nums []int, target int) int {
	total := 0
	for _, v := range nums {
		total += v
	}
	if (total+target)%2 != 0 {
		return 0
	}

	left := (total + target) / 2
	if left < 0 {
		return 0
	}
	dp := make([]int, left+1)
	dp[0] = 1

	for _, v := range nums {
		for volume := left; volume >= v; volume-- {
			dp[volume] += dp[volume-v]
		}
	}

	return dp[left]
}
