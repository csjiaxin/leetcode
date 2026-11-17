package l62

// MonotoneGrid 按行填表；cell 只能读取已完成的上、左前驱。
// 缺失前驱为零值，是否可用必须由 cell 根据 r/c 判断。
func MonotoneGrid[T any](rows, cols int, cell func(r, c int, up, left T) T) [][]T {
	dp := make([][]T, rows)
	for r := 0; r < rows; r++ {
		dp[r] = make([]T, cols)
		for c := 0; c < cols; c++ {
			var up, left T
			if r > 0 {
				up = dp[r-1][c]
			}
			if c > 0 {
				left = dp[r][c-1]
			}
			dp[r][c] = cell(r, c, up, left)
		}
	}
	return dp
}

// 起点为 1，其他格累计上方和左方；边界缺失前驱贡献 0。
func uniquePaths(m int, n int) int {
	if m <= 0 || n <= 0 {
		return 0
	}
	dp := MonotoneGrid(m, n, func(r, c, up, left int) int {
		if r == 0 && c == 0 {
			return 1
		}
		return up + left
	})
	return dp[m-1][n-1]
}
