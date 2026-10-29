package l856

// The score of a balanced parentheses string is based on the following rule:

// "()" has score 1.
// AB has score A + B, where A and B are balanced parentheses strings.
// (A) has score 2 * A, where A is a balanced parentheses string.

// Example 1:

// Input: s = "()"
// Output: 1
// Example 2:

// Input: s = "(())"
// Output: 2
// Example 3:

// Input: s = "()()()()"
// Output: 2

//()()
//(())

type frame struct {
	times  int
	prefix string
}

func isDigit(r rune) bool {
	return r >= '0' && r <= '9'
}

func scoreOfParentheses(s string) int {
	var (
		stack []int
	)
	stack = append(stack, 0)
	for _, c := range s {
		if c == '(' {
			stack = append(stack, 0)
		} else {
			inner := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			this_layer := 1
			if inner == 0 {

			} else {
				this_layer = inner * 2
			}

			stack[len(stack)-1] += this_layer
		}
	}
	return stack[0]
}

// 3[a3[b]c4[d]]
// 3,
