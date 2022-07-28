package main

import "fmt"

var dp [][]int

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func minDistance(s string, p string) int {
	m := len(s)
	n := len(p)
	dp = make([][]int, len(s)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(p)+1)
	}
	dp[0][0] = 0
	for i := 0; i <= m; i++ {
		dp[i][0] = i
	}
	for j := 1; j <= n; j++ {
		dp[0][j] = j
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			ci := i - 1
			cj := j - 1
			if s[ci] == p[cj] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				//insert, //replace, //remove
				dp[i][j] = min(dp[i][j-1]+1, dp[i-1][j-1]+1)
				dp[i][j] = min(dp[i][j], dp[i-1][j]+1)
			}
		}
	}

	return dp[m][n]
}

func main() {
	fmt.Println(minDistance("a", "ab"))
	// fmt.Println(minDistance("intention", "execution")) //5
}
