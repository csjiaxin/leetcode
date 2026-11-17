package l567

// https://leetcode.com/problems/permutation-in-string/
// checkInclusion reports whether s2 contains a permutation of s1.
// Both strings contain only lowercase English letters.
func checkInclusion(s1 string, s2 string) bool {
	windowSize := len(s1)
	if windowSize > len(s2) {
		return false
	}

	var need, window [26]int
	for i := 0; i < windowSize; i++ {
		need[s1[i]-'a']++
		window[s2[i]-'a']++
	}
	if window == need {
		return true
	}

	for right := windowSize; right < len(s2); right++ {
		window[s2[right]-'a']++
		window[s2[right-windowSize]-'a']--
		if window == need {
			return true
		}
	}
	return false
}
