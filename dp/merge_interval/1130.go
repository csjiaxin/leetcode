package main

import (
	"fmt"
)

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func mctFromLeafValues(arr []int) int {
	n := len(arr)
	dp := make([][]int, n)
	maxs := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
		maxs[i] = make([]int, n)
	}
	MAX := 99999999
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			dp[i][j] = MAX
		}
		dp[i][i] = 0
	}
	// for i := 0; i < n-1; i++ {
	// 	dp[i][i+1] = arr[i] * arr[i+1]
	// }

	for i := 0; i < n; i++ {
		m := arr[i]
		for j := i; j < n; j++ {
			m = max(m, arr[j])
			maxs[i][j] = m
		}
	}

	for l := 2; l <= n; l++ {
		for i := 0; i+l-1 < n; i++ {
			j := i + l - 1

			for left := i; left < j; left++ {
				//[i:left] -> left tree
				v := dp[i][left] + dp[left+1][j] + maxs[i][left]*maxs[left+1][j]
				dp[i][j] = min(dp[i][j], v)

			}
		}
	}
	return dp[0][n-1]
}

func main() {
	fmt.Println(mctFromLeafValues([]int{6, 2, 4}))
}
