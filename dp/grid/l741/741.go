package l741

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

// 把往返改为两人同步前进；步数推出列坐标，状态仅保存两行坐标，不可达状态不能参与加分。
func cherryPickup(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 || grid[0][0] == -1 {
		return 0
	}
	m, n := len(grid), len(grid[0])
	const invalid = -1 << 30
	initial := Filled(m, m, invalid)
	initial[0][0] = grid[0][0]
	last := Layers(m+n-2, initial, func(layer int, previous [][]int) [][]int {
		step := layer + 1
		next := Filled(m, m, invalid)
		for r1 := 0; r1 < m; r1++ {
			for r2 := 0; r2 < m; r2++ {
				c1, c2 := step-r1, step-r2
				if c1 < 0 || c1 >= n || c2 < 0 || c2 >= n || grid[r1][c1] == -1 || grid[r2][c2] == -1 {
					continue
				}
				best := invalid
				for _, p1 := range []int{r1, r1 - 1} {
					for _, p2 := range []int{r2, r2 - 1} {
						if p1 >= 0 && p2 >= 0 {
							best = max(best, previous[p1][p2])
						}
					}
				}
				if best == invalid {
					continue
				}
				reward := grid[r1][c1]
				if r1 != r2 {
					reward += grid[r2][c2]
				}
				next[r1][r2] = best + reward
			}
		}
		return next
	})
	return max(0, last[m-1][m-1])
}
