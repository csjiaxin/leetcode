package l1000

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

// https://leetcode.com/problems/minimum-cost-to-merge-stones/
// 每个区间压到 (长度-1)%(k-1)+1 堆；合法分割步长 k-1，能压成一堆时才加合并代价。
func mergeStones(stones []int, k int) int {
	n := len(stones)
	if n <= 1 {
		return 0
	}
	if k <= 1 || (n-1)%(k-1) != 0 {
		return -1
	}
	prefix := make([]int, n+1)
	for i, value := range stones {
		prefix[i+1] = prefix[i] + value
	}
	dp := IntervalTable(n, func(int) int { return 0 }, func(l, r int, dp [][]int) int {
		best := int(^uint(0) >> 1)
		for split := l; split < r; split += k - 1 {
			best = min(best, dp[l][split]+dp[split+1][r])
		}
		if (r-l)%(k-1) == 0 {
			best += prefix[r+1] - prefix[l]
		}
		return best
	})
	return dp[0][n-1]
}
