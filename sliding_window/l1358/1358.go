package l1358

// https://leetcode.com/problems/number-of-substrings-containing-all-three-characters/
// numberOfSubstrings counts substrings containing a, b and c at least once each.
// s contains only a, b and c.
func numberOfSubstrings(s string) int {

	N := len(s)
	window := map[byte]int{}

	validWindow := func() bool {
		return window['a'] > 0 && window['b'] > 0 && window['c'] > 0
	}
	ret := 0
	left := 0
	for right := 0; right < N; right++ {
		c := s[right]
		window[c] += 1
		for validWindow() {
			ret += N - 1 - right + 1
			window[s[left]] -= 1
			left += 1
		}
	}
	return ret

}
