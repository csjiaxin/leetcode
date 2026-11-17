package l312

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

// https://leetcode.com/problems/burst-balloons/
// dp[l][r] 清空原下标闭区间 l..r，外邻居 nums[l-1]、nums[r+1] 固定保留；越界邻居视为 1。
func maxCoins(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	// 直接使用原下标，不创建补边界数组。
	neighbor := func(i int) int {
		if i < 0 || i >= n {
			return 1
		}
		return nums[i]
	}
	dp := IntervalTable(n, func(i int) int {
		return neighbor(i-1) * nums[i] * neighbor(i+1)
	}, func(l, r int, dp [][]int) int {
		best := 0
		for last := l; last <= r; last++ {
			// last 最后戳破时，区间内其余气球已消失，只剩两个外邻居。
			score := neighbor(l-1) * nums[last] * neighbor(r+1)
			if last > l {
				score += dp[l][last-1]
			}
			if last < r {
				score += dp[last+1][r]
			}
			best = max(best, score)
		}
		return best
	})
	return dp[0][n-1]
}
