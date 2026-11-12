package l76

// https://leetcode.com/problems/minimum-window-substring/
// minWindow returns the shortest substring covering t's character multiplicities,
// or an empty string if none exists. Letters are case-sensitive.
func minWindow(s string, t string) string {
	if len(t) == 0 || len(t) > len(s) {
		return ""
	}

	var need, window [256]int
	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}

	// matched counts filled demand slots, including repeated characters in t.
	left, matched := 0, 0
	bestStart, bestLen := 0, len(s)+1
	for right := 0; right < len(s); right++ {
		c := s[right]
		window[c]++
		if window[c] <= need[c] {
			matched++
		}

		for matched == len(t) {
			if length := right - left + 1; length < bestLen {
				bestStart, bestLen = left, length
			}

			removed := s[left]
			// Check before decrementing: excess copies do not fill demand slots.
			if window[removed] <= need[removed] {
				matched--
			}
			window[removed]--
			left++
		}
	}

	if bestLen > len(s) {
		return ""
	}
	return s[bestStart : bestStart+bestLen]
}
