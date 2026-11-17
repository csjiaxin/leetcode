// LeetCode 583: https://leetcode.com/problems/delete-operation-for-two-strings/
package l583

import prefixdp "leetcode/dp/strings"

// dp[i][j] 表示让 word1[:i] 与 word2[:j] 相等的最少删除次数。
func minDistance(word1 string, word2 string) int {
	table := prefixdp.Table(len(word1), len(word2),
		func(i, j int) int { return i + j },
		func(i, j int, up, left, diagonal int) int {
			if word1[i-1] == word2[j-1] {
				return diagonal
			}
			if up < left {
				return up + 1
			}
			return left + 1
		})
	return table[len(word1)][len(word2)]
}
