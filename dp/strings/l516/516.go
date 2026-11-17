package l516

func longestPalindromeSubseq(s string) int {
	if len(s) == 0 {
		return 0
	}
	dp := IntervalTable(len(s), 0, func(i int) int { return 1 }, func(l, r, withoutLeft, withoutRight, inside int) int {
		if s[l] == s[r] {
			return inside + 2
		}
		return max(withoutLeft, withoutRight)
	})
	return dp[0][len(s)-1]
}

// IntervalTable 按区间长度填表；insideEmpty 是空内部区间的单位值。
func IntervalTable[T any](n int, insideEmpty T, base func(i int) T, step func(l, r int, withoutLeft, withoutRight, inside T) T) [][]T {
	dp := make([][]T, n)
	for i := range dp {
		dp[i] = make([]T, n)
		dp[i][i] = base(i)
	}
	for length := 2; length <= n; length++ {
		for l := 0; l+length <= n; l++ {
			r := l + length - 1
			inside := insideEmpty
			if length > 2 {
				inside = dp[l+1][r-1]
			}
			dp[l][r] = step(l, r, dp[l+1][r], dp[l][r-1], inside)
		}
	}
	return dp
}
