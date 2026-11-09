package l931

import (
	"slices"
)

// minFallingPathSum returns the minimum sum of a falling path through matrix.
// A path starts anywhere in the first row and moves to the next row in the
// same column or one column to the left or right.
func minFallingPathSum(triangle [][]int) int {
	m := len(triangle)
	n := len(triangle[m-1])
	dp := make([][]int, m)

	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	// for i := 0; i < m; i++ {
	// 	for j := 0; j < n; j++ {
	// 		dp[i][j] = math.MaxInt
	// 	}
	// }

	dp[0][0] = triangle[0][0]
	for j := 0; j < n; j++ {
		dp[0][j] = triangle[0][j]
	}
	for j := 1; j < m; j++ {
		for k := 0; k < len(triangle[j]); k++ {

			dp[j][k] = dp[j-1][k]
			if k-1 >= 0 {
				dp[j][k] = min(dp[j][k], dp[j-1][k-1])
			}
			if k+1 < n {
				dp[j][k] = min(dp[j][k], dp[j-1][k+1])
			}
			dp[j][k] += triangle[j][k]
		}
	}
	return slices.Min(dp[m-1])
}
