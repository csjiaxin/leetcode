package l1143

// https://leetcode.com/problems/longest-common-subsequence/
func longestCommonSubsequence(text1 string, text2 string) int {
	n1 := len(text1)
	n2 := len(text2)
	dp := make([][]int, n1+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, n2+1)
	}

	for i := 1; i < len(dp); i++ {
		iidx := i - 1
		for j := 1; j < n2+1; j++ {
			jidx := j - 1
			if text1[iidx] == text2[jidx] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[n1][n2]
}
