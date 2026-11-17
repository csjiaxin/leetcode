package l120

import (
	"math"
	"slices"
)

// minimumTotal returns the minimum top-to-bottom path sum in triangle.
// From index j in one row, the next step may use index j or j+1 in the next row.
func minimumTotal(triangle [][]int) int {
	m := len(triangle)
	n := len(triangle[m-1])
	dp := make([][]int, m)

	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			dp[i][j] = math.MaxInt
		}
	}

	dp[0][0] = triangle[0][0]
	for j := 1; j < n; j++ {
		dp[j][0] = dp[j-1][0] + triangle[j][0]
	}
	for j := 1; j < m; j++ {
		for k := 1; k < len(triangle[j]); k++ {
			dp[j][k] = min(dp[j-1][k], dp[j-1][k-1]) + triangle[j][k]
		}
	}
	return slices.Min(dp[m-1])
}
