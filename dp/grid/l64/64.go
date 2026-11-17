package l64

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

// 缺失前驱不能当成零代价捷径；第一行和第一列分别处理。
func minPathSum(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	m, n := len(grid), len(grid[0])
	dp := MonotoneGrid(m, n, func(r, c, up, left int) int {
		if r == 0 && c == 0 {
			return grid[r][c]
		}
		if r == 0 {
			return left + grid[r][c]
		}
		if c == 0 {
			return up + grid[r][c]
		}
		return min(up, left) + grid[r][c]
	})
	return dp[m-1][n-1]
}
