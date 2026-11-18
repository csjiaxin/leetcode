package l486

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

// https://leetcode.com/problems/predict-the-winner/
// 状态单调变小；下面的适配器完整定义终止条件、合法动作与当前行动者收益。
func predictTheWinner(nums []int) bool {
	type interval struct{ left, right int }
	var diff func(interval) int
	diff = MemoGame(func(s interval) int {
		if s.left > s.right {
			return 0
		}
		left := nums[s.left] - diff(interval{s.left + 1, s.right})
		right := nums[s.right] - diff(interval{s.left, s.right - 1})
		if left > right {
			return left
		}
		return right
	})
	return diff(interval{0, len(nums) - 1}) >= 0 // 平局也算先手赢。
}
