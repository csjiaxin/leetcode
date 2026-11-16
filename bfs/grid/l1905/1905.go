package l1905

// FloodFill 使用共享访问表处理一个四连通块，不修改原网格。
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

// https://leetcode.com/problems/count-sub-islands/
// grid1、grid2 为同形矩形网格，二者均不修改。
func countSubIslands(grid1, grid2 [][]int) int {
	if len(grid2) == 0 || len(grid2[0]) == 0 {
		return 0
	}
	rows, cols := len(grid2), len(grid2[0])
	visited := make([][]bool, rows)
	for r := range visited {
		visited[r] = make([]bool, cols)
	}
	count := 0
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid2[r][c] != 1 || visited[r][c] {
				continue
			}
			valid := true
			FloodFill(rows, cols, r, c, visited,
				func(r, c int) bool { return grid2[r][c] == 1 },
				func(r, c int) {
					// 只改变判定，不提前停止整块遍历。
					if grid1[r][c] != 1 {
						valid = false
					}
				},
			)
			if valid {
				count++
			}
		}
	}
	return count
}
