package l1510

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

// https://leetcode.com/problems/stone-game-iv/
func winnerSquareGame(n int) bool {
	win := MemoGame(func(remaining int, solve func(int) bool) bool {
		for root := 1; root <= remaining/root; root++ {
			if !solve(remaining - root*root) {
				return true
			}
		}
		return false
	})
	return win(n)
}
