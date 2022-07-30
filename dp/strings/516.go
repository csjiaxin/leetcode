package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func LCS(text1, text2 string) int {
	n1 := len(text1)
	n2 := len(text2)

	dp := make([][]int, n1+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, n2+1)
	}
	for i := 1; i <= n1; i++ {
		for j := 1; j <= n2; j++ {
			_i := i - 1
			_j := j - 1
			if text1[_i] == text2[_j] {
				dp[i][j] = dp[i-1][j-1] + 1
			}
			dp[i][j] = max(dp[i][j], dp[i-1][j])
			dp[i][j] = max(dp[i][j], dp[i][j-1])
		}
	}
	return dp[n1][n2]
}
func reverse(s string) string {
	b := []byte(s)
	for i, j := 0, len(s)-1; i <= j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return string(b)
}
func longestPalindromeSubseq(s string) int {
	s2 := reverse(s)
	return LCS(s, s2)
}

func main() {
	fmt.Println(longestPalindromeSubseq("bbbab"))
}
