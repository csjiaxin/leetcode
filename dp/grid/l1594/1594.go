package l1594

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

// 负数会交换极大与极小乘积；路径比较前不能取模，最后才对非负最大值取模。
func maxProductPath(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return -1
	}
	type product struct {
		low, high int64
		valid     bool
	}
	m, n := len(grid), len(grid[0])
	dp := MonotoneGrid(m, n, func(r, c int, up, left product) product {
		value := int64(grid[r][c])
		if r == 0 && c == 0 {
			return product{value, value, true}
		}
		result := product{}
		for _, source := range []product{up, left} {
			if !source.valid {
				continue
			}
			a, b := source.low*value, source.high*value
			if !result.valid {
				result = product{min(a, b), max(a, b), true}
			} else {
				result.low = min(result.low, a, b)
				result.high = max(result.high, a, b)
			}
		}
		return result
	})
	answer := dp[m-1][n-1].high
	if answer < 0 {
		return -1
	}
	return int(answer % 1000000007)
}
