package l1510

// MemoGame 为 evaluate 的结果添加缓存；每次调用创建独立缓存。
// evaluate 如需递归，应调用外部声明的缓存函数；状态转移必须无环。
func MemoGame[S comparable, V any](evaluate func(S) V) func(S) V {
	memo := make(map[S]V)
	return func(state S) V {
		if value, ok := memo[state]; ok {
			return value
		}
		value := evaluate(state)
		memo[state] = value
		return value
	}
}

// https://leetcode.com/problems/stone-game-iv/
// 状态单调变小；下面的适配器完整定义终止条件、合法动作与当前行动者收益。
func winnerSquareGame(n int) bool {
	var win func(int) bool
	win = MemoGame(func(remaining int) bool {
		for root := 1; root <= remaining/root; root++ {
			if !win(remaining - root*root) {
				return true
			}
		}
		return false
	})
	return win(n)
}
