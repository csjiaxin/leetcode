package main

// MemoWays 仅适用于依赖无环的状态；缓存的零值也是合法答案。
func MemoWays[S comparable, V any](evaluate func(S, func(S) V) V) func(S) V {
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

// 递归版固定硬币下标，只在同一下标重复选币，避免排列重复计数。
func change(amount int, coins []int) int {
	type state struct{ index, remaining int }
	solve := MemoWays(func(s state, next func(state) int) int {
		if s.remaining == 0 {
			return 1
		}
		if s.index == len(coins) || s.remaining < 0 {
			return 0
		}
		return next(state{s.index + 1, s.remaining}) + next(state{s.index, s.remaining - coins[s.index]})
	})
	return solve(state{0, amount})
}
