package l733

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

// https://leetcode.com/problems/flood-fill/
// 返回逐行复制的新图像，即使新旧颜色相同也不与输入共享行。
func floodFill(image [][]int, sr, sc, color int) [][]int {
	if image == nil {
		return nil
	}
	result := make([][]int, len(image))
	for r := range image {
		if image[r] != nil {
			result[r] = append(make([]int, 0, len(image[r])), image[r]...)
		}
	}
	if len(image) == 0 || len(image[0]) == 0 {
		return result
	}
	rows, cols := len(image), len(image[0])
	if sr < 0 || sr >= rows || sc < 0 || sc >= cols {
		return result
	}
	visited := make([][]bool, rows)
	for r := range visited {
		visited[r] = make([]bool, cols)
	}
	original := image[sr][sc]
	FloodFill(rows, cols, sr, sc, visited,
		func(r, c int) bool { return image[r][c] == original },
		func(r, c int) { result[r][c] = color },
	)
	return result
}
