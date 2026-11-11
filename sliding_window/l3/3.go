package l3

// https://leetcode.com/problems/longest-substring-without-repeating-characters/
// lengthOfLongestSubstring returns the longest substring length without repeats.
// The problem's letters, digits, symbols and spaces can be handled as bytes.
func lengthOfLongestSubstring(s string) int {
	N := len(s)
	window := [256]int{}
	left := 0
	ret := 0

	for i := 0; i < N; i++ {
		c := s[i]
		window[c] += 1

		for window[c] > 1 {
			window[s[left]] -= 1
			left += 1
		}

		ret = max(ret, i-left+1)
	}
	return ret
}
