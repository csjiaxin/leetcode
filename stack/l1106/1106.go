package l1106

// Standalone nested-context subtemplate and problem adapter; see README.md.
type booleanFrame struct {
	operator          byte
	anyTrue, anyFalse bool
}

// evaluateBoolean folds each operator's direct children, never skipping syntax.
func evaluateBoolean(expression string) bool {
	stack := []booleanFrame{}
	operator := byte(0)
	result := false
	feed := func(value bool) {
		if len(stack) == 0 {
			result = value
			return
		}
		frame := &stack[len(stack)-1]
		frame.anyTrue = frame.anyTrue || value
		frame.anyFalse = frame.anyFalse || !value
	}
	for i := range expression {
		token := expression[i]
		switch token {
		case '!', '&', '|':
			operator = token
		case '(':
			stack = append(stack, booleanFrame{operator: operator})
		case 't':
			feed(true)
		case 'f':
			feed(false)
		case ')':
			frame := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			value := false
			switch frame.operator {
			case '!':
				value = frame.anyFalse
			case '&':
				value = !frame.anyFalse
			case '|':
				value = frame.anyTrue
			}
			feed(value)
		}
	}
	return result
}

func parseBoolExpr(expression string) bool { return evaluateBoolean(expression) }
