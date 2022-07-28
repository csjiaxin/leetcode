package main

import "fmt"

func get(s string) (i0 int, i1 int) {
	for _, c := range s {
		if c == '0' {
			i0++
		} else {
			i1++
		}
	}
	return
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func findMaxForm(strs []string, m int, n int) int {
	dp := make([][]int, m+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, n+1)
	}
	for i := 0; i < len(strs); i++ {
		i0, i1 := get(strs[i])
		for mi := m; mi >= i0; mi-- {
			for ni := n; ni >= i1; ni-- {
				dp[mi][ni] = max(dp[mi-i0][ni-i1]+1, dp[mi][ni])
			}
		}
	}
	return dp[m][n]
}

func main() {
	fmt.Println(findMaxForm([]string{"10", "0001", "111001", "1", "0"}, 5, 3))
}
