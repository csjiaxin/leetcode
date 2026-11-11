package l556

import (
	"slices"
	"strconv"
)

// https://leetcode.com/problems/next-greater-element-iii/
// nextGreaterElement returns the smallest integer greater than n with the same
// digits, or -1 if none exists within the signed 32-bit integer range.
func nextGreaterElement(n int) int {
	digits := []byte(strconv.Itoa(n))
	N := len(digits)

	for i := N - 1; i > 0; i-- {
		if digits[i-1] < digits[i] {
			for j := N - 1; j >= i; j-- {
				if digits[j] > digits[i-1] {
					digits[i-1], digits[j] = digits[j], digits[i-1]
					slices.Reverse(digits[i:])

					// Use int64 so conversion is safe on 32-bit platforms too.
					var ret int64
					for _, digit := range digits {
						ret = ret*10 + int64(digit-'0')
						if ret > 1<<31-1 {
							return -1
						}
					}
					return int(ret)
				}
			}
		}
	}
	return -1
}
