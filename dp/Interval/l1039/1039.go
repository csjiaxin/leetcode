package main

import (
	"fmt"
	"math"
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

func minScoreTriangulation(values []int) int {
	n := len(values)
	N := n + 2
	V := make([]int, N)
	V[0] = values[n-1]
	copy(V[1:], values)
	V[N-1] = values[0]

	dp := make([][]int, N)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, N)
		for j := 0; j < len(dp[i]); j++ {
			dp[i][j] = math.MaxInt
		}
	}
	for i := 1; i < N-1; i++ {
		dp[i][i] = V[i-1] * V[i] * V[i+1]
	}
	for l := 2; l <= n; l++ {
		for i := 1; i+l-1 < N-1; i++ {
			j := i + l - 1
			bond := V[i-1] * V[j+1]
			for last := i; last <= j; last++ {
				cost := V[last] * bond
				if last-1 >= i {
					cost += dp[i][last-1]
				}
				if last+1 <= j {
					cost += dp[last+1][j]
				}
				dp[i][j] = min(dp[i][j], cost)
			}
		}
	}
	remove_len := n - 2
	ret := math.MaxInt
	for i := 1; i+remove_len-1 <= N-1; i++ {
		j := i + remove_len - 1
		ret = min(ret, dp[i][j])
	}
	return ret
}

func main() {
	fmt.Println(minScoreTriangulation([]int{1, 2, 3}))
	fmt.Println(minScoreTriangulation([]int{3, 7, 4, 5}))
}
