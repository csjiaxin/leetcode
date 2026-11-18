package l1140

import "math"

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

// https://leetcode.com/problems/stone-game-ii/
// 状态单调变小；下面的适配器完整定义终止条件、合法动作与当前行动者收益。
func stoneGameII(piles []int) int {
	n := len(piles)
	suffix := make([]int, n+1)
	for i := n - 1; i >= 0; i-- {
		suffix[i] = suffix[i+1] + piles[i]
	}
	type state struct{ start, m int }
	var diff func(state) int
	diff = MemoGame(func(s state) int {
		if s.start >= n {
			return 0
		}
		best := math.MinInt
		for X := 1; X <= 2*s.m && s.start+X-1 < n; X++ {
			end := s.start + X - 1
			score := suffix[s.start] - suffix[end+1] - diff(state{
				start: end + 1,
				m:     max(s.m, X),
			})
			best = max(best, score)
		}
		return best
	})
	// 总分固定：先手分数 = (总分 + 先手分差) / 2。
	return (suffix[0] + diff(state{0, 1})) / 2
}
