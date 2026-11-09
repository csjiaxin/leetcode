package l1277

func countSquares(matrix [][]int) int {
	m := len(matrix)
	if m == 0 {
		return 0
	}
	n := len(matrix[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	ret := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 1 {
				dp[i][j] = 1
				ret += 1
			}
		}
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][j] == 1 {
				dp[i][j] = min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1
				ret += (dp[i][j] - 1)
			}
		}
	}
	return ret
}
