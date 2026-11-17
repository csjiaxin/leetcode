package l1547

import (
	"math"
	"sort"
)

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
	m := len(cuts)
	if m == 0 {
		return 0
	}
	positions := append([]int(nil), cuts...)
	sort.Ints(positions)
	// 不补边界：下标 -1 对应木棍左端 0，下标 m 对应右端 n。
	boundary := func(i int) int {
		if i < 0 {
			return 0
		}
		if i >= m {
			return n
		}
		return positions[i]
	}

	// l、r 是排序后的切点下标；当前木段的端点是相邻切点 l-1、r+1。
	dp := IntervalTable(m, func(i int) int {
		return boundary(i+1) - boundary(i-1)
	}, func(l, r int, dp [][]int) int {
		best := math.MaxInt
		for first := l; first <= r; first++ {
			cost := boundary(r+1) - boundary(l-1)
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
	return dp[0][m-1]
}
