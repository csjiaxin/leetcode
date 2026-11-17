package stack

import "strings"

type frame struct {
	times  int
	prefix string
}

func isDigit(r rune) bool {
	return r >= '0' && r <= '9'
}

func decodeString(s string) string {
	var (
		stack   []frame
		times   int
		current strings.Builder
	)

	for _, c := range s {
		if isDigit(c) {
			times = times*10 + int(c-'0')
		} else if c == '[' {
			stack = append(stack, frame{
				times:  times,
				prefix: current.String(),
			})
			times = 0
			current.Reset()
		} else if c == ']' {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			repeated := strings.Repeat(current.String(), top.times)
			current.Reset()
			current.WriteString(top.prefix)
			current.WriteString(repeated)
		} else {
			current.WriteRune(c)
		}
	}
	return current.String()
}

// 3[a3[b]c4[d]]
// 3,
