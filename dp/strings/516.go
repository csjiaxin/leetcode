package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

var ret int

func dfs(s string, i, j int) {
	for ; i >= 0 && j < len(s) && s[i] == s[j]; i, j = i-1, j+1 {
		ret = max(ret, j-i+1)
	}
}
func longestPalindromeSubseq(s string) int {
	ret = 0
	for i := 0; i < len(s); i++ {
		dfs(s, i, i)
		dfs(s, i, i+1)
	}
	return ret
}

func main() {
	fmt.Println(longestPalindromeSubseq("bbbab"))
}
