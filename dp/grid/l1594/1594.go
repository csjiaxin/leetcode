package l1594

// maxProductPath returns the maximum non-negative right/down path product,
// modulo 1e9+7, or -1 if every path has a negative product.
func maxProductPath(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	maxDP := make([][]int64, m)
	minDP := make([][]int64, m)
	for i := range grid {
		maxDP[i] = make([]int64, n)
		minDP[i] = make([]int64, n)
	}

	maxDP[0][0] = int64(grid[0][0])
	minDP[0][0] = int64(grid[0][0])
	for j := 1; j < n; j++ {
		maxDP[0][j] = maxDP[0][j-1] * int64(grid[0][j])
		minDP[0][j] = maxDP[0][j]
	}
	for i := 1; i < m; i++ {
		maxDP[i][0] = maxDP[i-1][0] * int64(grid[i][0])
		minDP[i][0] = maxDP[i][0]
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			v := int64(grid[i][j])
			// A negative value can turn the minimum predecessor into the maximum.
			a, b := maxDP[i-1][j]*v, minDP[i-1][j]*v
			c, d := maxDP[i][j-1]*v, minDP[i][j-1]*v
			maxDP[i][j] = max(a, b, c, d)
			minDP[i][j] = min(a, b, c, d)
		}
	}

	best := maxDP[m-1][n-1]
	if best < 0 {
		return -1
	}
	// Keep exact products during DP; modulo would change their ordering.
	return int(best % 1_000_000_007)
}
