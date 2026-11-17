package l321

// https://leetcode.com/problems/create-maximum-number/
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
func suffixGreater(a, b []int) bool {
	i := 0
	for i < len(a) && i < len(b) && a[i] == b[i] {
		i++
	}
	return i < len(a) && (i == len(b) || a[i] > b[i])
}
func mergeLargest(a, b []int) []int {
	result := make([]int, 0, len(a)+len(b))
	for len(a) > 0 || len(b) > 0 {
		if suffixGreater(a, b) {
			result = append(result, a[0])
			a = a[1:]
		} else {
			result = append(result, b[0])
			b = b[1:]
		}
	}
	return result
}

func maxNumber(nums1 []int, nums2 []int, k int) []int {
	best := []int{}
	for take := 0; take <= k; take++ {
		if take > len(nums1) || k-take > len(nums2) {
			continue
		}
		a := greedySubsequence(nums1, take, func(current, previous int) bool { return current > previous })
		b := greedySubsequence(nums2, k-take, func(current, previous int) bool { return current > previous })
		candidate := mergeLargest(a, b)
		if suffixGreater(candidate, best) {
			best = candidate
		}
	}
	return best
}
