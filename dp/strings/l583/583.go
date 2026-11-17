// LeetCode 583: https://leetcode.com/problems/delete-operation-for-two-strings/
package l583

// dp[i][j] 表示让 word1[:i] 与 word2[:j] 相等的最少删除次数。
func minDistance(word1 string, word2 string) int {
	table := Table(len(word1), len(word2),
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
