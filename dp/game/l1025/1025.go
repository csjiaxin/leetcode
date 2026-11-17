package l1025

// MemoGame 仅缓存无环状态的求值；每次调用创建独立缓存。
func MemoGame[S comparable, V any](evaluate func(S, func(S) V) V) func(S) V {
	memo := make(map[S]V)
	var solve func(S) V
	solve = func(state S) V {
		if value, ok := memo[state]; ok {
			return value
		}
		value := evaluate(state, solve)
		memo[state] = value
		return value
	}
	return solve
}

// https://leetcode.com/problems/divisor-game/
// 状态单调变小；下面的适配器完整定义终止条件、合法动作与当前行动者收益。
func divisorGame(n int) bool {
	win := MemoGame(func(remaining int, solve func(int) bool) bool {
		for x := 1; x < remaining; x++ {
			if remaining%x == 0 && !solve(remaining-x) {
				return true
			}
		}
		return false // 无法操作，当前玩家失败。
	})
	return win(n)
}
