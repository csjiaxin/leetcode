package l443

import "slices"

// https://leetcode.com/problems/string-compression/
// compress writes the compressed result into chars and returns its length.
func compress(chars []byte) int {
	write := 0
	for read := 0; read < len(chars); {
		start := read
		char := chars[read]
		for read < len(chars) && chars[read] == char {
			read++
		}

		// Read the entire group before overwriting any of its bytes.
		chars[write] = char
		write++
		if count := read - start; count > 1 {
			digitsStart := write
			for ; count > 0; count /= 10 {
				chars[write] = '0' + byte(count%10)
				write++
			}
			// Digits were written least significant first.
			slices.Reverse(chars[digitsStart:write])
		}
	}
	return write
}
