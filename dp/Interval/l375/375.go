package l375

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

// https://leetcode.com/problems/guess-number-higher-or-lower-ii/
// 先对猜错后的分支取最坏 max，再对猜测值取最好 min；下标对应数值 guess+1。
func getMoneyAmount(n int) int {
	if n <= 1 {
		return 0
	}
	dp := IntervalTable(n, func(int) int { return 0 }, func(l, r int, dp [][]int) int {
		best := int(^uint(0) >> 1)
		for guess := l; guess <= r; guess++ {
			left, right := 0, 0
			if guess > l {
				left = dp[l][guess-1]
			}
			if guess < r {
				right = dp[guess+1][r]
			}
			best = min(best, guess+1+max(left, right))
		}
		return best
	})
	return dp[0][n-1]
}
