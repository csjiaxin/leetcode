package l1937

import "slices"

func maxPoints(points [][]int) int64 {
	m := len(points)
	n := len(points[0])

	dp := make([][]int64, m)
	for j := range points {
		dp[j] = make([]int64, n)
	}
	for j := 0; j < n; j++ {
		dp[m-1][j] = int64(points[m-1][j])
	}

	for i := m - 2; i >= 0; i-- {
		left := dp[i+1][0]
		for j := 0; j < n; j++ {
			// Best next-row choice at or to the left of j.
			left = max(left-1, dp[i+1][j])
			dp[i][j] = left + int64(points[i][j])
		}

		right := dp[i+1][n-1]
		for j := n - 1; j >= 0; j-- {
			// Best next-row choice at or to the right of j.
			right = max(right-1, dp[i+1][j])
			dp[i][j] = max(dp[i][j], right+int64(points[i][j]))
		}
	}
	return slices.Max(dp[0])
}
