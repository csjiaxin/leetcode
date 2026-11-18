package main

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

// 历史记忆化版本保留为独立包，统一使用同一无环求值骨架。
func stoneGameII(piles []int) int {
	n := len(piles)
	suffix := make([]int, n+1)
	for i := n - 1; i >= 0; i-- {
		suffix[i] = suffix[i+1] + piles[i]
	}
	type state struct{ start, m int }
	var diff func(state) int
	diff = MemoGame(func(s state) int {
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
			candidate := taken - diff(state{s.start + x, nextM})
			if !initialized || candidate > best {
				best, initialized = candidate, true
			}
		}
		return best
	})
	// 总分固定：先手分数 = (总分 + 先手分差) / 2。
	return (suffix[0] + diff(state{0, 1})) / 2
}
