package l1547

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

// https://leetcode.com/problems/minimum-cost-to-cut-a-stick/
// 状态是尚待切割的切点闭区间；枚举第一次切割，段长来自外侧固定端点，不修改 cuts。
func minCost(n int, cuts []int) int {
	if len(cuts) == 0 {
		return 0
	}
	positions := append([]int{0}, cuts...)
	positions = append(positions, n)
	// 切点最多 100 个；插入排序避免模板需要额外导入。
	for i := 1; i < len(positions); i++ {
		for j := i; j > 0 && positions[j] < positions[j-1]; j-- {
			positions[j], positions[j-1] = positions[j-1], positions[j]
		}
	}
	count := len(cuts)
	dp := IntervalTable(count, func(i int) int { return positions[i+2] - positions[i] },
		func(l, r int, dp [][]int) int {
			best := int(^uint(0) >> 1)
			for first := l; first <= r; first++ {
				cost := positions[r+2] - positions[l]
				if first > l {
					cost += dp[l][first-1]
				}
				if first < r {
					cost += dp[first+1][r]
				}
				best = min(best, cost)
			}
			return best
		})
	return dp[0][count-1]
}
