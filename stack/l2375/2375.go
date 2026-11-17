package l2375

// Standalone nested-context subtemplate and problem adapter; see README.md.
// patternStackDigits flushes an increasing pending segment at each I/end.
func patternStackDigits(pattern string) string {
	pending, result := []byte{}, []byte{}
	for i := 0; i <= len(pattern); i++ {
		pending = append(pending, byte('1'+i))
		if i == len(pattern) || pattern[i] == 'I' {
			for len(pending) > 0 {
				result = append(result, pending[len(pending)-1])
				pending = pending[:len(pending)-1]
			}
		}
	}
	return string(result)
}

func smallestNumber(pattern string) string { return patternStackDigits(pattern) }
