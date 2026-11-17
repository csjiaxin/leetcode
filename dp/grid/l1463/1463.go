package l1463

// Layers 的 transition 读取上一层，返回独立的新一层；不能覆盖 previous 的切片内容。
func Layers[S any](count int, initial S, transition func(layer int, previous S) S) S {
	state := initial
	for layer := 0; layer < count; layer++ {
		state = transition(layer, state)
	}
	return state
}

func Filled(rows, cols, value int) [][]int {
	result := make([][]int, rows)
	for r := range result {
		result[r] = make([]int, cols)
		for c := range result[r] {
			result[r][c] = value
		}
	}
	return result
}

// 层是行号，状态是两列位置；同格奖励只计算一次；O(mn²)。
func cherryPickup(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	m, n := len(grid), len(grid[0])
	const invalid = -1 << 30
	initial := Filled(n, n, invalid)
	initial[0][n-1] = grid[0][0]
	if n > 1 {
		initial[0][n-1] += grid[0][n-1]
	}
	last := Layers(m-1, initial, func(layer int, previous [][]int) [][]int {
		r := layer + 1
		next := Filled(n, n, invalid)
		for c1 := 0; c1 < n; c1++ {
			for c2 := 0; c2 < n; c2++ {
				if previous[c1][c2] == invalid {
					continue
				}
				for d1 := -1; d1 <= 1; d1++ {
					for d2 := -1; d2 <= 1; d2++ {
						a, b := c1+d1, c2+d2
						if a < 0 || a >= n || b < 0 || b >= n {
							continue
						}
						reward := grid[r][a]
						if a != b {
							reward += grid[r][b]
						}
						next[a][b] = max(next[a][b], previous[c1][c2]+reward)
					}
				}
			}
		}
		return next
	})
	answer := 0
	for _, row := range last {
		for _, value := range row {
			answer = max(answer, value)
		}
	}
	return answer
}
