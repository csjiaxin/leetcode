package l2684

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

// 每步列号增加，返回移动边数而不是访问格数；从第一列所有起点取最大。
func maxMoves(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	m, n := len(grid), len(grid[0])
	solve := MemoGrid(func(s [2]int, next func([2]int) int) int {
		best := 0
		for d := -1; d <= 1; d++ {
			r, c := s[0]+d, s[1]+1
			if r >= 0 && r < m && c < n && grid[r][c] > grid[s[0]][s[1]] {
				best = max(best, 1+next([2]int{r, c}))
			}
		}
		return best
	})
	answer := 0
	for r := 0; r < m; r++ {
		answer = max(answer, solve([2]int{r, 0}))
	}
	return answer
}
