package l827

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

// https://leetcode.com/problems/making-a-large-island/
func largestIsland(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	rows, cols := len(grid), len(grid[0])
	visited := make([][]bool, rows)
	labels := make([][]int, rows)
	for r := range visited {
		visited[r] = make([]bool, cols)
		labels[r] = make([]int, cols)
	}
	areas := map[int]int{}
	id, best := 0, 0
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] != 1 || visited[r][c] {
				continue
			}
			id++
			FloodFill(rows, cols, r, c, visited,
				func(r, c int) bool { return grid[r][c] == 1 },
				func(r, c int) {
					labels[r][c] = id
					areas[id]++
				},
			)
			if areas[id] > best {
				best = areas[id]
			}
		}
	}
	directions := [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] != 0 {
				continue
			}
			area := 1
			seen := map[int]bool{}
			for _, d := range directions {
				nr, nc := r+d[0], c+d[1]
				if nr < 0 || nr >= rows || nc < 0 || nc >= cols {
					continue
				}
				neighborID := labels[nr][nc]
				// 按编号去重：同一块不能重复，不同块面积相同也都要加。
				if neighborID != 0 && !seen[neighborID] {
					seen[neighborID] = true
					area += areas[neighborID]
				}
			}
			if area > best {
				best = area
			}
		}
	}
	return best
}
