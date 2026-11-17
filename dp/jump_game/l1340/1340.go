package l1340

// MemoJump 仅适用于依赖无环的状态；缓存的零值也是合法答案。
func MemoJump[S comparable, V any](evaluate func(S, func(S) V) V) func(S) V {
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

// https://leetcode.com/problems/jump-game-v/
// 只能严格下降，故依赖无环；遇到不矮的柱子必须停止该方向，答案包含起点。
func maxJumps(arr []int, d int) int {
	solve := MemoJump(func(i int, next func(int) int) int {
		best := 1
		for _, direction := range []int{-1, 1} {
			for step := 1; step <= d; step++ {
				j := i + direction*step
				if j < 0 || j >= len(arr) || arr[j] >= arr[i] {
					break
				}
				best = max(best, 1+next(j))
			}
		}
		return best
	})
	answer := 0
	for i := range arr {
		answer = max(answer, solve(i))
	}
	return answer
}
