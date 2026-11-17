package l1277

// SquareSides[r+1][c+1] 表示以 (r,c) 为右下角的全一正方形最大边长。
func SquareSides(rows, cols int, isOne func(r, c int) bool) [][]int {
	side := make([][]int, rows+1)
	for r := range side {
		side[r] = make([]int, cols+1)
	}
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if isOne(r, c) {
				side[r+1][c+1] = 1 + min(side[r][c+1], side[r+1][c], side[r][c])
			}
		}
	}
	return side
}

// 同一边长状态，累加每个格子可贡献的 1..side 个正方形。
func countSquares(matrix [][]int) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	side := SquareSides(len(matrix), len(matrix[0]), func(r, c int) bool { return matrix[r][c] == 1 })
	answer := 0
	for _, row := range side {
		for _, value := range row {
			answer += value
		}
	}
	return answer
}
