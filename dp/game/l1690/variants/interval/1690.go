package l1690

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

// https://leetcode.com/problems/stone-game-vii/
// 区间迭代版：移除端点后取得的是剩余石子总和，减去对手分差。
func stoneGameVII(stones []int) int {
	n := len(stones)
	if n < 2 {
		return 0
	}
	prefix := make([]int, n+1)
	for i, v := range stones {
		prefix[i+1] = prefix[i] + v
	}
	dp := IntervalTable(n, func(int) int { return 0 }, func(l, r int, dp [][]int) int {
		return max(prefix[r+1]-prefix[l+1]-dp[l+1][r], prefix[r]-prefix[l]-dp[l][r-1])
	})
	return dp[0][n-1]
}
