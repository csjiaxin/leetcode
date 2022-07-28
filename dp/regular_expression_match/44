package main

import "fmt"

var dp [][]bool

func isMatch(s string, p string) bool {
	m := len(s)
	n := len(p)
	dp = make([][]bool, len(s)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]bool, len(p)+1)
	}
	dp[0][0] = true
	for j := 1; j <= n && p[j-1] == '*'; j++ {
		dp[0][j] = true
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			ci := i - 1
			cj := j - 1
			if p[cj] == '*' {
				for k := 0; k <= i; k++ {
					dp[i][j] = dp[i][j] || dp[i-k][j-1]
				}
			} else {
				if s[ci] == p[cj] || p[cj] == '?' {
					dp[i][j] = dp[i][j] || dp[i-1][j-1]
				}
			}

		}
	}

	return dp[m][n]
}

func main() {
	fmt.Println(isMatch("adceb", "*a*b")) //true
	// fmt.Println(isMatch("aa", "*")) //true
	// // fmt.Println(isMatch("ab", ".*..")) //true
	// // fmt.Println(isMatch("a", "ab*"))     //true
	// fmt.Println(isMatch("aab", "c*a*b")) //true
	// fmt.Println(isMatch("abc", ".*"))
}
