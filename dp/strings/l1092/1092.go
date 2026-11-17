// Package l1092 solves LeetCode 1092: Shortest Common Supersequence.
// https://leetcode.com/problems/shortest-common-supersequence/
package l1092

// shortestCommonSupersequence returns a shortest string containing both inputs
// as subsequences. dp[i][j] is its minimum length for str1[:i] and str2[:j].
// The full length table takes O(m*n) time and space. Backtracking appends each
// output byte once, followed by an O(m+n) reversal. Strings contain ASCII letters.
func shortestCommonSupersequence(str1 string, str2 string) string {
	m, n := len(str1), len(str2)
	dp := Table[int](m, n,
		func(i, j int) int {
			return i + j // On a boundary, the other prefix is the answer.
		},
		func(i, j, up, left, diagonal int) int {
			if str1[i-1] == str2[j-1] {
				return diagonal + 1 // Share the matching final character.
			}
			if up < left {
				return up + 1
			}
			return left + 1
		},
	)

	// Walk from the final prefixes to the empty prefixes. The chosen characters
	// are discovered right to left, so collect them in a reverse buffer.
	reversed := make([]byte, 0, dp[m][n])
	i, j := m, n
	for i > 0 && j > 0 {
		if str1[i-1] == str2[j-1] {
			reversed = append(reversed, str1[i-1])
			i--
			j--
		} else if dp[i-1][j] < dp[i][j-1] {
			reversed = append(reversed, str1[i-1])
			i--
		} else {
			// Either side is valid on a tie; choose str2 consistently.
			reversed = append(reversed, str2[j-1])
			j--
		}
	}
	for i > 0 {
		i--
		reversed = append(reversed, str1[i])
	}
	for j > 0 {
		j--
		reversed = append(reversed, str2[j])
	}
	for left, right := 0, len(reversed)-1; left < right; left, right = left+1, right-1 {
		reversed[left], reversed[right] = reversed[right], reversed[left]
	}
	return string(reversed)
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
