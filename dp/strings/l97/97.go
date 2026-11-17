package l97

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

	dp := Table[bool](m, n,
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
