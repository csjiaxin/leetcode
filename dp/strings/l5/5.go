package l5

import "math"

// https://leetcode.com/problems/longest-palindromic-substring/
func longestPalindrome(s string) string {
	L := len(s)
	dp := make([][]bool, L)
	val := math.MinInt
	valStr := ""
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]bool, L)
		dp[i][i] = true
		val = 1
		valStr = s[i : i+1]
	}
	for i := len(dp) - 1; i >= 0; i-- {
		for j := i + 1; j < len(dp[0]); j++ {

			dp[i][j] = s[i] == s[j] && (j-i == 1 || dp[i+1][j-1])
			if dp[i][j] && j-i+1 > val {
				val = j - i + 1
				valStr = s[i : j+1]
			}
		}
	}

	return valStr
}
