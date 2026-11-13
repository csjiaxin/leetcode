package l459

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

// https://leetcode.com/problems/repeated-substring-pattern/
func repeatedSubstringPattern(s string) bool {
	pi := buildPrefix(s)
	N := len(s)
	j := pi[N-1]
	diff := N - j
	return j > 0 && N%diff == 0
}
