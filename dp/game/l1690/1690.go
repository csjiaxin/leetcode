package l1690

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

// https://leetcode.com/problems/stone-game-vii/
// 状态单调变小；下面的适配器完整定义终止条件、合法动作与当前行动者收益。
func stoneGameVII(stones []int) int {
	prefix := make([]int, len(stones)+1)
	for i, value := range stones {
		prefix[i+1] = prefix[i] + value
	}
	type interval struct{ left, right int }
	var diff func(interval) int
	diff = MemoGame(func(s interval) int {
		if s.left > s.right {
			return 0
		}
		left := prefix[s.right+1] - prefix[s.left+1] - diff(interval{
			left:  s.left + 1,
			right: s.right,
		})
		right := prefix[s.right] - prefix[s.left] - diff(interval{
			left:  s.left,
			right: s.right - 1,
		})
		return max(left, right)
	})
	return diff(interval{0, len(stones) - 1})
}
