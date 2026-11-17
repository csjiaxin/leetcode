package main

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

// 顶点闭区间 l..r；边界边 (l,r) 对应三角形第三点 k，两个子区间共享 k。
func minScoreTriangulation(values []int) int {
	n := len(values)
	if n < 3 {
		return 0
	}
	dp := IntervalTable(n, func(int) int { return 0 }, func(l, r int, dp [][]int) int {
		if r-l < 2 {
			return 0
		}
		best := int(^uint(0) >> 1)
		for k := l + 1; k < r; k++ {
			best = min(best, dp[l][k]+dp[k][r]+values[l]*values[k]*values[r])
		}
		return best
	})
	return dp[0][n-1]
}
