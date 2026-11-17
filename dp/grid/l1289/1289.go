package l1289

import (
	"math"
	"slices"
)

func minFallingPathSum(matrix [][]int) int {
	m := len(matrix)
	if m == 0 {
		return 0
	}
	n := len(matrix[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	for j := 0; j < n; j++ {
		dp[0][j] = matrix[0][j]
	}
	for i := 1; i < m; i++ {
		for j := 0; j < n; j++ {
			minVal := math.MaxInt

			for k := 0; k < n; k++ {
				if k != j {
					minVal = min(minVal, dp[i-1][k])
				}
			}
			dp[i][j] = minVal + matrix[i][j]
		}
	}
	return slices.Min(dp[m-1])
}
