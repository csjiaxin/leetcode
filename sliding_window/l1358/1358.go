package l1358

// https://leetcode.com/problems/number-of-substrings-containing-all-three-characters/
// The standalone helper below matches the named subtemplate in README.md.
// windowCoverCount counts all future right endpoints for each valid left endpoint.
func windowCoverCount(n int, add, remove func(int), valid func() bool) int {
	left, count := 0, 0
	for right := 0; right < n; right++ {
		add(right)
		for left <= right && valid() {
			count += n - right
			remove(left)
			left++
		}
	}
	return count
}

func numberOfSubstrings(s string) int {
	var count [3]int
	return windowCoverCount(len(s),
		func(i int) { count[s[i]-'a']++ },
		func(i int) { count[s[i]-'a']-- },
		func() bool { return count[0] > 0 && count[1] > 0 && count[2] > 0 })
}
