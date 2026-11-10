package l1547

import "sort"

// https://leetcode.com/problems/minimum-cost-to-cut-a-stick/
func minCost(n int, cuts []int) int {
	points := make([]int, len(cuts)+2)
	copy(points[1:], cuts)
	points[len(points)-1] = n
	sort.Ints(points)
	m := len(points)
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, m)
	}
	for gap := 2; gap < m; gap++ {
		for i := 0; i+gap < m; i++ {
			j := i + gap
			best := int(^uint(0) >> 1)
			// Choose the first cut within these fixed endpoints.
			for cut := i + 1; cut < j; cut++ {
				best = min(best, dp[i][cut]+dp[cut][j]+points[j]-points[i])
			}
			dp[i][j] = best
		}
	}
	return dp[0][m-1]
}
