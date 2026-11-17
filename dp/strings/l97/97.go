package l97

import prefixdp "leetcode/dp/strings"

// https://leetcode.com/problems/interleaving-string/
// dp[i][j] 表示 s1[:i] 和 s2[:j] 能否交错组成 s3[:i+j]。
func isInterleave(s1 string, s2 string, s3 string) bool {
	m, n := len(s1), len(s2)
	if m+n != len(s3) {
		return false
	}

	// 边界只有一个源字符串可用，预处理前缀匹配，避免反复比较子串。
	firstOnly := make([]bool, m+1)
	secondOnly := make([]bool, n+1)
	firstOnly[0], secondOnly[0] = true, true
	for i := 1; i <= m; i++ {
		firstOnly[i] = firstOnly[i-1] && s1[i-1] == s3[i-1]
	}
	for j := 1; j <= n; j++ {
		secondOnly[j] = secondOnly[j-1] && s2[j-1] == s3[j-1]
	}

	dp := prefixdp.Table[bool](m, n,
		func(i, j int) bool {
			if j == 0 {
				return firstOnly[i]
			}
			return secondOnly[j]
		},
		func(i, j int, up, left, _ bool) bool {
			last := s3[i+j-1]
			// 最后一个字符来自 s1，或者来自 s2；各自内部顺序不能改变。
			return (up && s1[i-1] == last) ||
				(left && s2[j-1] == last)
		},
	)
	return dp[m][n]
}
