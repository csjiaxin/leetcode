package l1406

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

// https://leetcode.com/problems/stone-game-iii/
func stoneGameIII(stoneValue []int) string {
	diff := MemoGame(func(start int, solve func(int) int) int {
		if start == len(stoneValue) {
			return 0
		}
		// 不能用 0 初始化：必须取石子，即使所有选择都亏分。
		taken := stoneValue[start]
		best := taken - solve(start+1)
		for count := 2; count <= 3 && start+count <= len(stoneValue); count++ {
			taken += stoneValue[start+count-1]
			candidate := taken - solve(start+count)
			if candidate > best {
				best = candidate
			}
		}
		return best
	})
	result := diff(0)
	if result > 0 {
		return "Alice"
	}
	if result < 0 {
		return "Bob"
	}
	return "Tie"
}
