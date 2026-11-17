// Package l115 solves LeetCode 115: Distinct Subsequences.
// https://leetcode.com/problems/distinct-subsequences/
package l115

// numDistinct counts the ways to obtain t by deleting characters from s.
//
// dp[i][j] counts ways to form t[:j] from s[:i]. Intermediate prefix
// counts can exceed even int64 although the problem guarantees that the final
// answer fits a signed 32-bit integer. We therefore store min(count, limit).
// Every transition is a nonnegative addition, so capping preserves this
// invariant; states below the cap are exact, and the promised final answer
// remains exact. The guarded addition is safe on both 32-bit and 64-bit Go.
// Time and space: O(len(s) * len(t)). Strings contain ASCII letters.
func numDistinct(s string, t string) int {
	const limit = 1<<31 - 1

	dp := Table[int](len(s), len(t),
		func(i, j int) int {
			if j == 0 {
				return 1 // There is one way to form the empty target.
			}
			return 0 // An empty source cannot form a nonempty target.
		},
		func(i, j, up, left, diagonal int) int {
			if s[i-1] != t[j-1] {
				return up // Skip s[i-1].
			}
			// Either skip s[i-1], or use it to match t[j-1].
			if up > limit-diagonal {
				return limit
			}
			return up + diagonal
		},
	)
	return dp[len(s)][len(t)]
}

// Table 构建 (m+1) × (n+1) 的前缀状态表。
//
// dp[i][j] 对应第一个序列的前 i 个元素与第二个序列的前 j 个元素。
// base 只负责第 0 行、第 0 列；step 负责其余格子。
// step 的 up、left、diagonal 分别是 dp[i-1][j]、dp[i][j-1]、dp[i-1][j-1]。
//
// m、n 必须非负，回调必须在需要调用时有效。回调通过闭包读取原序列。
// 当回调为 O(1) 时，时间、空间均为 O((m+1)*(n+1))。
func Table[T any](
	m, n int,
	base func(i, j int) T,
	step func(i, j int, up, left, diagonal T) T,
) [][]T {
	dp := make([][]T, m+1)
	for i := range dp {
		dp[i] = make([]T, n+1)
	}

	for i := 0; i <= m; i++ {
		dp[i][0] = base(i, 0)
	}
	for j := 1; j <= n; j++ {
		dp[0][j] = base(0, j)
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			dp[i][j] = step(i, j, dp[i-1][j], dp[i][j-1], dp[i-1][j-1])
		}
	}
	return dp
}
