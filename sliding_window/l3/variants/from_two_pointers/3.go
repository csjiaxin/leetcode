package l3

// https://leetcode.com/problems/longest-substring-without-repeating-characters/
// The standalone helper below matches the named subtemplate in README.md.
// windowLongest maintains a valid suffix after removing invalid left endpoints.
func windowLongest(n int, add, remove func(int), invalid func() bool) int {
	left, best := 0, 0
	for right := 0; right < n; right++ {
		add(right)
		for left <= right && invalid() {
			remove(left)
			left++
		}
		if length := right - left + 1; length > best {
			best = length
		}
	}
	return best
}

func lengthOfLongestSubstring(s string) int {
	var count [256]int
	duplicates := 0
	return windowLongest(len(s),
		func(i int) {
			count[s[i]]++
			if count[s[i]] == 2 {
				duplicates++
			}
		},
		func(i int) {
			if count[s[i]] == 2 {
				duplicates--
			}
			count[s[i]]--
		},
		func() bool { return duplicates > 0 })
}
