package l856

// Standalone nested-context subtemplate and problem adapter; see README.md.
// parenthesesFold stores each open layer's sum of completed children.
func parenthesesFold(s string, wrap func(int) int) int {
	stack := []int{0}
	for _, token := range s {
		if token == '(' {
			stack = append(stack, 0)
			continue
		}
		inner := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		stack[len(stack)-1] += wrap(inner)
	}
	return stack[0]
}

func scoreOfParentheses(s string) int {
	return parenthesesFold(s, func(inner int) int {
		if inner == 0 {
			return 1
		}
		return 2 * inner
	})
}
