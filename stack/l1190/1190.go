package l1190

// Standalone nested-context subtemplate and problem adapter; see README.md.
// reverseNestedPairs jumps between paired delimiters and reverses traversal.
func reverseNestedPairs(s string) string {
	pair := make([]int, len(s))
	stack := []int{}
	for i := range s {
		if s[i] == '(' {
			stack = append(stack, i)
		} else if s[i] == ')' {
			j := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			pair[i] = j
			pair[j] = i
		}
	}
	result := make([]byte, 0, len(s))
	for i, direction := 0, 1; i >= 0 && i < len(s); i += direction {
		if s[i] == '(' || s[i] == ')' {
			i = pair[i]
			direction = -direction
		} else {
			result = append(result, s[i])
		}
	}
	return string(result)
}

func reverseParentheses(s string) string { return reverseNestedPairs(s) }
