package l1690

// https://leetcode.com/problems/stone-game-vii/
func stoneGameVII(stones []int) int {
	n := len(stones)
	if n < 2 {
		return 0
	}
	prefix := make([]int, n+1)
	dp := make([][]int, n)
	for i, stone := range stones {
		prefix[i+1] = prefix[i] + stone
		dp[i] = make([]int, n)
	}
	// dp[i][j] is the best current-player minus opponent score.
	for length := 2; length <= n; length++ {
		for i := 0; i+length <= n; i++ {
			j := i + length - 1
			left := prefix[j+1] - prefix[i+1] - dp[i+1][j]
			right := prefix[j] - prefix[i] - dp[i][j-1]
			dp[i][j] = max(left, right)
		}
	}
	return dp[0][n-1]
}
