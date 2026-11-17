// Package grid 提供共享访问表的四方向连通块遍历。
package grid

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
