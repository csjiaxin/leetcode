package l664

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

// https://leetcode.com/problems/strange-printer/
// 相同字符可以复用一次打印；把 s[l] 与某个 s[k] 的打印合并，不能只做普通加法切分。
func strangePrinter(s string) int {
	if len(s) == 0 {
		return 0
	}
	dp := IntervalTable(len(s), func(int) int { return 1 }, func(l, r int, dp [][]int) int {
		best := 1 + dp[l+1][r]
		for k := l + 1; k <= r; k++ {
			if s[k] != s[l] {
				continue
			}
			middle := 0
			if k > l+1 {
				middle = dp[l+1][k-1]
			}
			best = min(best, middle+dp[k][r])
		}
		return best
	})
	return dp[0][len(s)-1]
}
