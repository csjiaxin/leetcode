package l1190

import "strings"

func reverseString(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func reverseParentheses(s string) string {

	var (
		stack   []string
		current strings.Builder
	)

	for _, c := range s {
		if c == '(' {
			stack = append(stack, current.String())
			current.Reset()
		} else if c == ')' {
			// layer := len(stack)
			inner := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			current_string := reverseString(current.String())
			// if layer%2 == 1 {
			// 	current_string = reverseString(current_string)
			// }
			current.Reset()
			current.WriteString(inner)
			current.WriteString(current_string)
		} else {
			current.WriteRune(c)
		}
	}
	return current.String()
}
