package l1106

type frame struct {
	val  bool
	sign rune
}

func update(stack []frame, val bool) {
	end := len(stack) - 1
	top := stack[end]

	if top.sign == '|' {
		top.val = top.val || val
	} else if top.sign == '&' {
		top.val = top.val && val
	} else {
		top.val = !val
	}

	stack[end] = top
}

// https://leetcode.com/problems/parsing-a-boolean-expression/
func parseBoolExpr(expression string) bool {

	var (
		stack        []frame
		current_sign rune
	)

	for _, c := range expression {
		if c == '(' {
			e := frame{
				val:  false,
				sign: current_sign,
			}
			if current_sign == '&' {
				e.val = true
			}
			stack = append(stack, e)
		} else if c == ')' {
			inner := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				return inner.val
			}
			update(stack, inner.val)

		} else if c == '!' || c == '&' || c == '|' {
			current_sign = c
		} else if c == 'f' {
			if len(stack) == 0 {
				return false
			}
			update(stack, false)
		} else if c == 't' {
			if len(stack) == 0 {
				return true
			}
			update(stack, true)

		}
	}
	return false
}
