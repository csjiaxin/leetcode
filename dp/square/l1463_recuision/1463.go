package l1463

func cherryPickup(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])

	var valid func(i, j int) bool
	valid = func(i, j int) bool {
		return i >= 0 && i < m && j >= 0 && j < n

	}

	dir := []int{-1, 0, 1}
	memo := make([][][]int, m)
	for i := 0; i < m; i++ {
		memo[i] = make([][]int, n)
		for j := 0; j < n; j++ {
			memo[i][j] = make([]int, n)

			for k := 0; k < n; k++ {
				memo[i][j][k] = -1
			}
		}
	}

	var dfs func(r1, c1, c2 int) int
	dfs = func(r1, c1, c2 int) int {
		r2 := r1
		if !valid(r1, c1) || !valid(r2, c2) {
			return -1
		}
		val := memo[r1][c1][c2]

		if val < 0 {
			val = grid[r1][c1]
			if c1 != c2 {
				val += grid[r2][c2]
			}
			if r1 < m-1 {
				next := 0
				for _, d1 := range dir {
					for _, d2 := range dir {
						ret := dfs(r1+1, c1+d1, c2+d2)
						next = max(next, ret)
					}
				}
				val += next
			}
			memo[r1][c1][c2] = val
		}
		return val
	}
	return dfs(0, 0, n-1)
}
