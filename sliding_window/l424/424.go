package l424

// https://leetcode.com/problems/longest-repeating-character-replacement/
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

func characterReplacement(s string, k int) int {
	var count [26]int
	length := 0
	return windowLongest(len(s),
		func(i int) { count[s[i]-'A']++; length++ },
		func(i int) { count[s[i]-'A']--; length-- },
		func() bool {
			highest := 0
			for _, c := range count {
				if c > highest {
					highest = c
				}
			}
			return length-highest > k
		})
}
