package l1254

// FloodFill 访问起点所在的四连通块，先标记再调用 visit。
// visited 必须覆盖 rows*cols，canEnter 和 visit 必须非 nil。
// 跨多次调用共享 visited 可以枚举连通块；不会撤销访问标记。
// 回调不应修改 visited 或使其他格子的可进入条件发生变化。
// 本函数不接收原网格，是否修改输入由调用方的回调决定。
func FloodFill(rows, cols, startR, startC int, visited [][]bool, canEnter func(r, c int) bool, visit func(r, c int)) {
	var dfs func(int, int)
	dfs = func(r, c int) {
		if r < 0 || r >= rows || c < 0 || c >= cols {
			return
		}
		if visited[r][c] || !canEnter(r, c) {
			return
		}
		visited[r][c] = true
		visit(r, c)
		dfs(r-1, c)
		dfs(r+1, c)
		dfs(r, c-1)
		dfs(r, c+1)
	}
	dfs(startR, startC)
}

// https://leetcode.com/problems/number-of-closed-islands/
func closedIsland(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	rows, cols := len(grid), len(grid[0])
	visited := make([][]bool, rows)
	for r := range visited {
		visited[r] = make([]bool, cols)
	}
	fill := func(r, c int) {
		FloodFill(rows, cols, r, c, visited,
			func(r, c int) bool { return grid[r][c] == 0 },
			func(r, c int) {},
		)
	}
	// 此题 0 是陆地，先标记与边界相连的陆地。
	for r := 0; r < rows; r++ {
		fill(r, 0)
		fill(r, cols-1)
	}
	for c := 0; c < cols; c++ {
		fill(0, c)
		fill(rows-1, c)
	}
	count := 0
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == 0 && !visited[r][c] {
				count++
				fill(r, c)
			}
		}
	}
	return count
}
