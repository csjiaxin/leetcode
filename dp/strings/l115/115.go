// Package l115 solves LeetCode 115: Distinct Subsequences.
// https://leetcode.com/problems/distinct-subsequences/
package l115

import prefixdp "leetcode/dp/strings"

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

	dp := prefixdp.Table[int](len(s), len(t),
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
