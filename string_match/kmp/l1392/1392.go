package l1392

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

// https://leetcode.com/problems/longest-happy-prefix/
func longestPrefix(s string) string {
	pi := buildPrefix(s)
	l := pi[len(s)-1]
	return s[0:l]
}
