package l1937

import "slices"

func absInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

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
		for j := 0; j < n; j++ {
			next := dp[i+1][j]

			for k := 0; k < n; k++ {
				val := dp[i+1][k] - int64(absInt(j-k))
				next = max(next, val)
			}
			dp[i][j] = next + int64(points[i][j])
		}
	}
	return slices.Max(dp[0])
}
