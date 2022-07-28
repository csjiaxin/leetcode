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
func integerBreak(n int) int {
	dp := make([][]int, n+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, n+1)
	}
	return dp[m][n]
}

func main() {
	fmt.Println(findMaxForm([]string{"10", "0001", "111001", "1", "0"}, 5, 3))
}
