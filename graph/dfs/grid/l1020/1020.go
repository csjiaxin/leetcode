package l1020

import "leetcode/graph/dfs/grid"

// FloodFill 使用共享访问表处理一个四连通块，不修改原网格。
func FloodFill(rows, cols, startR, startC int, visited [][]bool, canEnter func(r, c int) bool, visit func(r, c int)) {
	grid.FloodFill(rows, cols, startR, startC, visited, canEnter, visit)
}

// https://leetcode.com/problems/number-of-enclaves/
func numEnclaves(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	rows, cols := len(grid), len(grid[0])
	visited := make([][]bool, rows)
	for r := range visited {
		visited[r] = make([]bool, cols)
	}
	markBoundary := func(r, c int) {
		FloodFill(rows, cols, r, c, visited,
			func(r, c int) bool { return grid[r][c] == 1 },
			func(r, c int) {},
		)
	}
	for r := 0; r < rows; r++ {
		markBoundary(r, 0)
		markBoundary(r, cols-1)
	}
	for c := 0; c < cols; c++ {
		markBoundary(0, c)
		markBoundary(rows-1, c)
	}
	count := 0
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			// 数剩余格子，不是连通块数。
			if grid[r][c] == 1 && !visited[r][c] {
				count++
			}
		}
	}
	return count
}
