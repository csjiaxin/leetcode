package l224

// Standalone nested-context subtemplate and problem adapter; see README.md.
type arithmeticFrame struct{ sum, sign int }

// evaluateSum restores sum/sign at each closing parenthesis; unary '-' is allowed.
func evaluateSum(s string) int {
	stack := []arithmeticFrame{}
	sum, number, sign := 0, 0, 1
	for i := range s {
		token := s[i]
		switch {
		case token >= '0' && token <= '9':
			number = number*10 + int(token-'0')
		case token == '+' || token == '-':
			sum += sign * number
			number = 0
			sign = 1
			if token == '-' {
				sign = -1
			}
		case token == '(':
			stack = append(stack, arithmeticFrame{sum, sign})
			sum, number, sign = 0, 0, 1
		case token == ')':
			sum += sign * number
			number = 0
			frame := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			sum = frame.sum + frame.sign*sum
		}
	}
	return sum + sign*number
}

func calculate(s string) int { return evaluateSum(s) }
