package main

import "fmt"

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

//s[i] == s[j]
func equal(s string, i, j int) bool {
	return s[i:i+1] == s[j:j+1]
}

var dp []map[int]bool

func try_update(end1, end2 int) {
	for k, _ := range dp[end2] {
		dp[end1+1][k+1] = true
	}
}
func checkPartitioning(s string) bool {
	n := len(s)
	dp = make([]map[int]bool, n+1) // number of cuts for the first k characters
	for i := 0; i <= n; i++ {
		dp[i] = map[int]bool{}
	}
	dp[0] = map[int]bool{0: true}

	for i := 0; i < n; i++ {
		for j := 0; i+j < n && i-j >= 0; j++ {
			if equal(s, i-j, i+j) {
				try_update(i+j, i-j)
			} else {
				break
			}
		}
		for j := 1; i+j < n && i-(j-1) >= 0; j++ {
			if equal(s, i+j, i-(j-1)) {
				try_update(i+j, i-(j-1))
			} else {
				break
			}
		}
	}

	if _, ok := dp[n][3]; ok {
		return true
	}
	return false
}

func main() {
	fmt.Println(checkPartitioning("bcbddxy"))
}
