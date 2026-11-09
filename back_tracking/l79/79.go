package l79

// exist reports whether word can be formed by horizontally or vertically
// adjacent cells in board. Each cell may be used at most once per word.
func exist(board [][]byte, word string) bool {
	m := len(board)
	n := len(board[0])
	visit := make([][]bool, m)
	for i := 0; i < m; i++ {
		visit[i] = make([]bool, n)
	}
	L := len(word)
	dir4 := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	var dfs func(i, j int, wordidx int) bool
	dfs = func(i, j int, wordidx int) bool {
		if word[wordidx] != board[i][j] {
			return false
		}
		visit[i][j] = true
		if wordidx == L-1 {
			return true
		}
		for _, d := range dir4 {
			ni := i + d[0]
			nj := j + d[1]
			if ni >= 0 && ni < m && nj >= 0 && nj < n && visit[ni][nj] == false {
				if dfs(ni, nj, wordidx+1) {
					return true
				}
			}
		}
		visit[i][j] = false

		return false
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if dfs(i, j, 0) {
				return true
			}
		}
	}
	return false
}
