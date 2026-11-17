package l1254

import "leetcode/graph/dfs/grid"

// FloodFill 使用共享访问表处理一个四连通块，不修改原网格。
func FloodFill(rows, cols, startR, startC int, visited [][]bool, canEnter func(r, c int) bool, visit func(r, c int)) {
	grid.FloodFill(rows, cols, startR, startC, visited, canEnter, visit)
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
