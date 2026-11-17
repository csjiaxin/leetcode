package stack

// Standalone nested-context subtemplate and problem adapter; see README.md.
type decoderFrame struct {
	prefix []byte
	times  int
}

// decodeNested restores the outer prefix and repeats a completed child.
func decodeNested(s string) string {
	stack := []decoderFrame{}
	current := []byte{}
	number := 0
	for i := 0; i < len(s); i++ {
		token := s[i]
		switch {
		case token >= '0' && token <= '9':
			number = number*10 + int(token-'0')
		case token == '[':
			stack = append(stack, decoderFrame{current, number})
			current = []byte{}
			number = 0
		case token == ']':
			frame := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			inner := current
			current = frame.prefix
			for repeat := 0; repeat < frame.times; repeat++ {
				current = append(current, inner...)
			}
		default:
			current = append(current, token)
		}
	}
	return string(current)
}

func decodeString(s string) string { return decodeNested(s) }
