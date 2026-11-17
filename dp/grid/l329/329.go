package l329

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

// 严格递增保证无环；每格答案是从该格出发的最长路径，O(mn)。
func longestIncreasingPath(matrix [][]int) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	m, n := len(matrix), len(matrix[0])
	directions := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	solve := MemoGrid(func(s [2]int, next func([2]int) int) int {
		best := 1
		for _, d := range directions {
			r, c := s[0]+d[0], s[1]+d[1]
			if r >= 0 && r < m && c >= 0 && c < n && matrix[r][c] > matrix[s[0]][s[1]] {
				best = max(best, 1+next([2]int{r, c}))
			}
		}
		return best
	})
	answer := 0
	for r := 0; r < m; r++ {
		for c := 0; c < n; c++ {
			answer = max(answer, solve([2]int{r, c}))
		}
	}
	return answer
}
