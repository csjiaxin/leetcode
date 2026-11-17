package l931

// Layers 的 transition 读取上一层，返回独立的新一层；不能覆盖 previous 的切片内容。
func Layers[S any](count int, initial S, transition func(layer int, previous S) S) S {
	state := initial
	for layer := 0; layer < count; layer++ {
		state = transition(layer, state)
	}
	return state
}

// 上一行同列及左右相邻列，读取独立旧层。
func minFallingPathSum(matrix [][]int) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	cols := len(matrix[0])
	last := Layers(len(matrix), []int(nil), func(r int, previous []int) []int {
		row := make([]int, cols)
		for c, value := range matrix[r] {
			if r == 0 {
				row[c] = value
				continue
			}
			best := previous[c]
			if c > 0 {
				best = min(best, previous[c-1])
			}
			if c+1 < cols {
				best = min(best, previous[c+1])
			}
			row[c] = best + value
		}
		return row
	})
	answer := last[0]
	for _, value := range last {
		answer = min(answer, value)
	}
	return answer
}
