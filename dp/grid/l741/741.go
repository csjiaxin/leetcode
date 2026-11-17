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
				memo[i][j][k] = -1
			}
		}
	}
	memo[0][0][0] = grid[0][0]
	dir4 := [][]int{{-1, 0, 0}, {-1, 0, -1}, {0, -1, 0}, {0, -1, -1}}
	for r1 := 0; r1 < m; r1++ {
		for r2 := 0; r2 < m; r2++ {
			for c1 := 0; c1 < n; c1++ {

				if r1 == 0 && r2 == 0 && c1 == 0 {
					continue
				}
				c2 := r1 + c1 - r2
				if !valid(r1, c1) || !valid(r2, c2) {
					continue
				}
				if grid[r1][c1] == -1 || grid[r2][c2] == -1 {
					continue
				}
				val := -1
				for _, d := range dir4 {
					lr1 := r1 + d[0]
					lc1 := c1 + d[1]
					lr2 := r2 + d[2]
					lc2 := lr1 + lc1 - lr2

					if valid(lr1, lc1) && valid(lr2, lc2) {
						val = max(memo[lr1][lc1][lr2], val)
					}

				}
				if val != -1 {
					val = val + grid[r1][c1]
					if r1 != r2 {
						val += grid[r2][c2]
					}
				}
				memo[r1][c1][r2] = val
			}

		}
	}

	return max(memo[m-1][n-1][m-1], 0)
}
