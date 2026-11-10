package l2684

func maxMoves(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])

	valid := func(i, j int) bool {
		return (i >= 0 && i < m && j >= 0 && j < n)
	}

	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)

		for j := 0; j < n; j++ {
			dp[i][j] = -1
		}
	}
	dir := []int{1, 0, -1}

	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		val := dp[i][j]
		if val < 0 {
			val = 0
			for _, d := range dir {
				ni := i + d
				nj := j + 1
				if valid(ni, nj) && grid[i][j] < grid[ni][nj] {
					val = max(val, dfs(ni, nj)+1)
				}
			}
			dp[i][j] = val
		}

		return dp[i][j]
	}
	ret := 0
	for i := 0; i < m; i++ {
		ret = max(ret, dfs(i, 0))
	}
	return ret
}
