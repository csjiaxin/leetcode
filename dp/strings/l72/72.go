package l72

import prefixdp "leetcode/dp/strings"

// https://leetcode.com/problems/edit-distance/
// dp[i][j] 表示 text1 的前 i 个字符变成 text2 的前 j 个字符所需的最少操作数。
func minDistance(text1 string, text2 string) int {
	m, n := len(text1), len(text2)
	dp := prefixdp.Table[int](m, n,
		func(i, j int) int {
			if i == 0 {
				return j // 空串变成目标前缀，需要插入 j 个字符。
			}
			return i // 源前缀变成空串，需要删除 i 个字符。
		},
		func(i, j int, up, left, diagonal int) int {
			if text1[i-1] == text2[j-1] {
				return diagonal
			}

			// up、left、diagonal 分别对应删除、插入、替换。
			best := up
			if left < best {
				best = left
			}
			if diagonal < best {
				best = diagonal
			}
			return best + 1
		},
	)
	return dp[m][n]
}
