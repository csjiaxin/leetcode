package l76

// https://leetcode.com/problems/minimum-window-substring/
// The standalone helper below matches the named subtemplate in README.md.
// windowShortest records valid windows before shrinking; length n+1 means absent.
func windowShortest(n int, add, remove func(int), valid func() bool) (int, int) {
	left, start, length := 0, 0, n+1
	for right := 0; right < n; right++ {
		add(right)
		for left <= right && valid() {
			if right-left+1 < length {
				start, length = left, right-left+1
			}
			remove(left)
			left++
		}
	}
	return start, length
}

func minWindow(s string, t string) string {
	if len(t) == 0 || len(t) > len(s) {
		return ""
	}
	var need, count [256]int
	missing := len(t)
	for i := range t {
		need[t[i]]++
	}
	start, length := windowShortest(len(s),
		func(i int) {
			c := s[i]
			if count[c] < need[c] {
				missing--
			}
			count[c]++
		},
		func(i int) {
			c := s[i]
			count[c]--
			if count[c] < need[c] {
				missing++
			}
		},
		func() bool { return missing == 0 })
	if length == len(s)+1 {
		return ""
	}
	return s[start : start+length]
}
