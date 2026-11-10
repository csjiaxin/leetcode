package l1049

// https://leetcode.com/problems/last-stone-weight-ii/
// Partition the stones into two groups with the smallest weight difference.
// Time: O(n * total), space: O(total), where total is the sum of weights.
func lastStoneWeightII(stones []int) int {
	total := 0
	for _, stone := range stones {
		total += stone
	}

	half := total / 2
	dp := make([]bool, half+1)
	dp[0] = true
	for _, stone := range stones {
		// Descending sums ensure each stone is used at most once.
		for sum := half; sum >= stone; sum-- {
			dp[sum] = dp[sum] || dp[sum-stone]
		}
	}

	for sum := half; sum >= 0; sum-- {
		if dp[sum] {
			return total - 2*sum
		}
	}
	return total
}
