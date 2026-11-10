package l664

// https://leetcode.com/problems/strange-printer/
func strangePrinter(s string) int {
	if len(s) == 0 {
		return 0
	}

	chars := make([]byte, 0, len(s))
	for i := range s {
		if i == 0 || s[i] != s[i-1] {
			chars = append(chars, s[i])
		}
	}
	n := len(chars)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
		dp[i][i] = 1
	}
	for length := 2; length <= n; length++ {
		for i := 0; i+length-1 < n; i++ {
			j := i + length - 1
			dp[i][j] = dp[i][j-1] + 1

			for k := i; k < j; k++ {
				if chars[k] == chars[j] {
					dp[i][j] = min(dp[i][j], dp[i][k]+dp[k+1][j-1])
				}
			}
		}
	}

	return dp[0][n-1]
}
