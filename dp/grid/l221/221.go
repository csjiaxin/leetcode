package l221

// maximalSquare returns the area of the largest square containing only '1's.
// It uses O(m*n) time and O(m*n) extra space without modifying matrix.
func maximalSquare(matrix [][]byte) int {
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
			if matrix[i][j] == '1' {
				dp[i][j] = 1
				ret = 1
			}
		}
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][j] == '1' {
				dp[i][j] += min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1])
				ret = max(dp[i][j], ret)

			}
		}
	}
	return ret * ret
}
