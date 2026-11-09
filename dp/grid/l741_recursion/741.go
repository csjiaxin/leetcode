package l741

func cherryPickup(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])

	var valid func(i, j int) bool
	valid = func(i, j int) bool {
		return i >= 0 && i < m && j >= 0 && j < n

	}
	memo := make([][][]int, m)
	for i := 0; i < m; i++ {
		memo[i] = make([][]int, n)
		for j := 0; j < n; j++ {
			memo[i][j] = make([]int, m)

			for k := 0; k < m; k++ {
				memo[i][j][k] = -2
			}
		}
	}

	var dfs func(r1, c1, r2 int) int
	dfs = func(r1, c1, r2 int) int {
		c2 := r1 + c1 - r2
		if !valid(r1, c1) || !valid(r2, c2) {
			return -1
		}
		if grid[r1][c1] < 0 || grid[r2][c2] < 0 {
			return -1
		}
		if r1 == m-1 && c1 == n-1 {
			return grid[r1][c1]
		}
		val := memo[r1][c1][r2]

		if val == -2 {
			current := grid[r1][c1]
			if r2 != r1 {
				current += grid[r2][c2]
			}

			result := max(dfs(r1+1, c1, r2+1),
				dfs(r1+1, c1, r2),
				dfs(r1, c1+1, r2+1),
				dfs(r1, c1+1, r2))

			if result < 0 {
				val = result
			} else {
				val = result + current
			}
		}
		memo[r1][c1][r2] = val
		return val
	}
	return max(dfs(0, 0, 0), 0)
}
