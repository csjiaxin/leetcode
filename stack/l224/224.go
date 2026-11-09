package l224

type frame struct {
	sign rune
	val  int
}

func update(stack []frame, sign rune, number int) {
	last := &stack[len(stack)-1]
	if sign == '-' {
		last.val -= number
	} else {
		last.val += number
	}
}

// https://leetcode.com/problems/basic-calculator/
func calculate(s string) int {
	var (
		stack        []frame
		current_sign rune
		number       int
	)
	stack = append(stack, frame{
		val:  0,
		sign: '+',
	})
	for _, c := range s {
		if c == '(' {
			e := frame{
				val:  0,
				sign: current_sign,
			}
			stack = append(stack, e)
			current_sign = '\x00'

		} else if c == ')' {
			update(stack, current_sign, number)
			number = 0

			inner := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			update(stack, inner.sign, inner.val)

		} else if c == '+' || c == '-' {
			update(stack, current_sign, number)
			number = 0
			current_sign = c
		} else if c == ' ' {

		} else {
			number = 10*number + int(c-'0')
		}
	}
	update(stack, current_sign, number)
	return stack[0].val
}
