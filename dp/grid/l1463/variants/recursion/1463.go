package l1463

// MemoGrid 仅适用于依赖无环的状态；缓存的零值也是合法答案。
func MemoGrid[S comparable, V any](evaluate func(S, func(S) V) V) func(S) V {
	memo := make(map[S]V)
	var solve func(S) V
	solve = func(s S) V {
		if value, ok := memo[s]; ok {
			return value
		}
		value := evaluate(s, solve)
		memo[s] = value
		return value
	}
	return solve
}

// 递归变体仍使用行号和两列坐标，保留独立记忆化方法。
func cherryPickup(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	m, n := len(grid), len(grid[0])
	const invalid = -1 << 30
	type state struct{ row, c1, c2 int }
	solve := MemoGrid(func(s state, next func(state) int) int {
		if s.c1 < 0 || s.c1 >= n || s.c2 < 0 || s.c2 >= n {
			return invalid
		}
		if s.row == m {
			return 0
		}
		reward := grid[s.row][s.c1]
		if s.c1 != s.c2 {
			reward += grid[s.row][s.c2]
		}
		best := invalid
		for a := -1; a <= 1; a++ {
			for b := -1; b <= 1; b++ {
				best = max(best, next(state{s.row + 1, s.c1 + a, s.c2 + b}))
			}
		}
		return reward + best
	})
	return solve(state{0, 0, n - 1})
}
