package l741

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

// 递归变体使用两个同步移动者的三维状态，非法状态不进入得分相加。
func cherryPickup(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	m, n := len(grid), len(grid[0])
	const invalid = -1 << 30
	type state struct{ r1, c1, r2 int }
	solve := MemoGrid(func(s state, next func(state) int) int {
		c2 := s.r1 + s.c1 - s.r2
		if s.r1 < 0 || s.r1 >= m || s.c1 < 0 || s.c1 >= n || s.r2 < 0 || s.r2 >= m || c2 < 0 || c2 >= n {
			return invalid
		}
		if grid[s.r1][s.c1] == -1 || grid[s.r2][c2] == -1 {
			return invalid
		}
		if s.r1 == m-1 && s.c1 == n-1 {
			return grid[s.r1][s.c1]
		}
		best := max(next(state{s.r1 + 1, s.c1, s.r2 + 1}), next(state{s.r1 + 1, s.c1, s.r2}),
			next(state{s.r1, s.c1 + 1, s.r2 + 1}), next(state{s.r1, s.c1 + 1, s.r2}))
		if best == invalid {
			return invalid
		}
		reward := grid[s.r1][s.c1]
		if s.r1 != s.r2 {
			reward += grid[s.r2][c2]
		}
		return reward + best
	})
	return max(0, solve(state{0, 0, 0}))
}
