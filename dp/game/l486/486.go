package l486

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

// https://leetcode.com/problems/predict-the-winner/
// 状态单调变小；下面的适配器完整定义终止条件、合法动作与当前行动者收益。
func predictTheWinner(nums []int) bool {
	type interval struct{ left, right int }
	diff := MemoGame(func(s interval, solve func(interval) int) int {
		if s.left > s.right {
			return 0
		}
		left := nums[s.left] - solve(interval{s.left + 1, s.right})
		right := nums[s.right] - solve(interval{s.left, s.right - 1})
		if left > right {
			return left
		}
		return right
	})
	return diff(interval{0, len(nums) - 1}) >= 0 // 平局也算先手赢。
}
