package l28

func buildPrefix(pattern string) []int {
	N := len(pattern)
	pi := make([]int, N)

	for i := 1; i < N; i++ {
		c := pattern[i]
		j := pi[i-1]
		for j > 0 && c != pattern[j] {
			j = pi[j-1]
		}
		if pattern[j] == c {
			j++
		}
		pi[i] = j
	}
	return pi
}

// https://leetcode.com/problems/find-the-index-of-the-first-occurrence-in-a-string/
func strStr(haystack string, pattern string) int {
	pi := buildPrefix(pattern)
	j := 0

	for i := 0; i < len(haystack); i++ {
		c := haystack[i]
		for j > 0 && pattern[j] != c {
			j = pi[j-1]
		}
		if c == pattern[j] {
			j++
		}
		if j == len(pattern) {
			return i - len(pattern) + 1
		}
	}
	return -1
}
