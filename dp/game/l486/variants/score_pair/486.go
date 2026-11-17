package l486

// IntervalTable 按闭区间长度递增；step 只能读取严格更短区间。
// base 初始化单点；空区间不在表中，由适配器自行处理。
func IntervalTable[T any](n int, base func(i int) T, step func(l, r int, dp [][]T) T) [][]T {
	dp := make([][]T, n)
	for i := range dp {
		dp[i] = make([]T, n)
		dp[i][i] = base(i)
	}
	for length := 2; length <= n; length++ {
		for l := 0; l+length <= n; l++ {
			r := l + length - 1
			dp[l][r] = step(l, r, dp)
		}
	}
	return dp
}

// https://leetcode.com/problems/predict-the-winner/
// 双方收益版本保留 E 状态；子区间当前行动者变为对手，因此角色互换。
type E struct{ firstMax, secondMax int }

func predictTheWinner(nums []int) bool {
	if len(nums) == 0 {
		return true
	}
	dp := IntervalTable(len(nums), func(i int) E { return E{nums[i], 0} }, func(l, r int, dp [][]E) E {
		left, right := nums[l]+dp[l+1][r].secondMax, nums[r]+dp[l][r-1].secondMax
		if left >= right {
			return E{left, dp[l+1][r].firstMax}
		}
		return E{right, dp[l][r-1].firstMax}
	})
	result := dp[0][len(nums)-1]
	return result.firstMax >= result.secondMax
}
