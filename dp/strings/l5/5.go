package l5

func longestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}
	start, best := 0, 1
	IntervalTable(len(s), true, func(i int) bool { return true }, func(l, r int, withoutLeft, withoutRight, inside bool) bool {
		valid := s[l] == s[r] && inside
		if valid && r-l+1 > best {
			start, best = l, r-l+1
		}
		return valid
	})
	return s[start : start+best]
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
