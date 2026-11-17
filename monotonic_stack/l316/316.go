package l316

// https://leetcode.com/problems/remove-duplicate-letters/
// Embedded subtemplate: see README.md for its invariant and tie policy.
// distinctSubsequence pops a larger symbol only when another copy remains.
func distinctSubsequence(s string) string {
	var remaining [256]int
	var used [256]bool
	for i := range s {
		remaining[s[i]]++
	}
	stack := []byte{}
	for i := range s {
		value := s[i]
		remaining[value]--
		if used[value] {
			continue
		}
		for len(stack) > 0 && stack[len(stack)-1] > value && remaining[stack[len(stack)-1]] > 0 {
			used[stack[len(stack)-1]] = false
			stack = stack[:len(stack)-1]
		}
		used[value] = true
		stack = append(stack, value)
	}
	return string(stack)
}

func removeDuplicateLetters(s string) string { return distinctSubsequence(s) }
