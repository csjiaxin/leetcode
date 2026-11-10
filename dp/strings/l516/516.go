package l516

// https://leetcode.com/problems/longest-palindromic-subsequence/
func longestPalindromeSubseq(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}

	// dp[i][j] is the longest palindromic subsequence length in s[i:j+1].
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
		dp[i][i] = 1
	}

	// Fill from shorter inner intervals to longer outer intervals.
	for i := n - 2; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = max(dp[i+1][j], dp[i][j-1])
			}
		}
	}
	return dp[0][n-1]
}
