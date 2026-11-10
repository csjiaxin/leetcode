package l1312

import "math"

// https://leetcode.com/problems/minimum-insertion-steps-to-make-a-string-palindrome/
func minInsertions(s string) int {
	L := len(s)
	dp := make([][]int, L)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, L)
		for j := 0; j < len(dp[i]); j++ {
			dp[i][j] = math.MaxInt
		}
		dp[i][i] = 0
	}
	for i := len(dp) - 1; i >= 0; i-- {
		for j := i + 1; j < len(dp[0]); j++ {

			if s[i] == s[j] {
				if j == i+1 {
					dp[i][j] = 0
				} else {
					dp[i][j] = dp[i+1][j-1]
				}
			} else {
				dp[i][j] = min(dp[i+1][j], dp[i][j-1]) + 1
			}

		}
	}

	return dp[0][L-1]
}
