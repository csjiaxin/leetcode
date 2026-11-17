package l130

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

// https://leetcode.com/problems/surrounded-regions/
// 非原地版本：不同于 LeetCode 的原地 solve 接口，本函数返回新棋盘。
func solve(board [][]byte) [][]byte {
	if board == nil {
		return nil
	}
	result := make([][]byte, len(board))
	for r := range board {
		if board[r] != nil {
			result[r] = append(make([]byte, 0, len(board[r])), board[r]...)
		}
	}
	if len(board) == 0 || len(board[0]) == 0 {
		return result
	}
	rows, cols := len(board), len(board[0])
	visited := make([][]bool, rows)
	for r := range visited {
		visited[r] = make([]bool, cols)
	}
	markSafe := func(r, c int) {
		FloodFill(rows, cols, r, c, visited,
			func(r, c int) bool { return board[r][c] == 'O' },
			func(r, c int) {},
		)
	}
	for r := 0; r < rows; r++ {
		markSafe(r, 0)
		markSafe(r, cols-1)
	}
	for c := 0; c < cols; c++ {
		markSafe(0, c)
		markSafe(rows-1, c)
	}
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if board[r][c] == 'O' && !visited[r][c] {
				result[r][c] = 'X'
			}
		}
	}
	return result
}
