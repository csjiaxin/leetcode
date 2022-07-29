package main

import "fmt"

var ret int

func dfs(s string, i, j int) {
	for ; i >= 0 && j < len(s) && s[i] == s[j]; i, j = i-1, j+1 {
		ret += 1
	}
}
func countSubstrings(s string) int {
	ret = 0
	for i := 0; i < len(s); i++ {
		dfs(s, i, i)
		dfs(s, i, i+1)
	}
	return ret
}

func main() {
	fmt.Println(longestPalindromeSubseq("abc"))
}
