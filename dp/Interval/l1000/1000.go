package l1000

// https://leetcode.com/problems/minimum-cost-to-merge-stones/
func mergeStones(stones []int, k int) int {
	n := len(stones)
	if n <= 1 {
		return 0
	}
	if (n-1)%(k-1) != 0 {
		return -1
	}
	prefix := make([]int, n+1)
	dp := make([][]int, n)
	for i, stone := range stones {
		prefix[i+1] = prefix[i] + stone
		dp[i] = make([]int, n)
	}
	// Each interval is reduced to (length-1)%(k-1)+1 piles.
	for length := 2; length <= n; length++ {
		for i := 0; i+length <= n; i++ {
			j := i + length - 1
			best := int(^uint(0) >> 1)
			// The left interval must be reducible to exactly one pile.
			for mid := i; mid < j; mid += k - 1 {
				best = min(best, dp[i][mid]+dp[mid+1][j])
			}
			if (length-1)%(k-1) == 0 {
				best += prefix[j+1] - prefix[i]
			}
			dp[i][j] = best
		}
	}
	return dp[0][n-1]
}
