package l402

// https://leetcode.com/problems/remove-k-digits/
// Embedded subtemplate: see README.md for its invariant and tie policy.
// greedySubsequence keeps k elements; better may replace a worse previous value.
func greedySubsequence[T any](values []T, k int, better func(current, previous T) bool) []T {
	if k < 0 {
		k = 0
	}
	if k > len(values) {
		k = len(values)
	}
	drop := len(values) - k
	stack := make([]T, 0, len(values))
	for _, value := range values {
		for drop > 0 && len(stack) > 0 && better(value, stack[len(stack)-1]) {
			stack = stack[:len(stack)-1]
			drop--
		}
		stack = append(stack, value)
	}
	return stack[:k]
}

func removeKdigits(num string, k int) string {
	digits := greedySubsequence([]byte(num), len(num)-k, func(current, previous byte) bool { return current < previous })
	start := 0
	for start < len(digits) && digits[start] == '0' {
		start++
	}
	if start == len(digits) {
		return "0"
	}
	return string(digits[start:])
}
