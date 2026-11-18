package l1406

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

// https://leetcode.com/problems/stone-game-iii/
// 状态单调变小；下面的适配器完整定义终止条件、合法动作与当前行动者收益。
func stoneGameIII(stoneValue []int) string {
	n := len(stoneValue)
	var diff func(int) int
	diff = MemoGame(func(start int) int {
		if start >= n {
			return 0
		}
		best := math.MinInt
		sum := 0
		for end := start; end-start+1 <= 3 && end < n; end++ {
			sum += stoneValue[end]
			score := sum - diff(end+1)
			best = max(score, best)
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
