package main

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

// 历史记忆化版本保留为独立包，统一使用同一无环求值骨架。
func stoneGameII(piles []int) int {
	n := len(piles)
	suffix := make([]int, n+1)
	for i := n - 1; i >= 0; i-- {
		suffix[i] = suffix[i+1] + piles[i]
	}
	type state struct{ start, m int }
	diff := MemoGame(func(s state, solve func(state) int) int {
		if s.start == n {
			return 0
		}
		// 原题全为正数，能全拿时直接结束。
		if 2*s.m >= n-s.start {
			return suffix[s.start]
		}
		best, initialized := 0, false
		for x := 1; x <= 2*s.m && s.start+x <= n; x++ {
			nextM := s.m
			if x > nextM {
				nextM = x
			}
			taken := suffix[s.start] - suffix[s.start+x]
			candidate := taken - solve(state{s.start + x, nextM})
			if !initialized || candidate > best {
				best, initialized = candidate, true
			}
		}
		return best
	})
	// 总分固定：先手分数 = (总分 + 先手分差) / 2。
	return (suffix[0] + diff(state{0, 1})) / 2
}
