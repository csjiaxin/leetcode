package l2104

// https://leetcode.com/problems/sum-of-subarray-ranges/
// Embedded subtemplate: see README.md for its invariant and tie policy.
// monotonicBoundaries settles right boundaries on pop and reads left on push.
// pop defines the strict/non-strict tie ownership for this specific problem.
func monotonicBoundaries(nums []int, pop func(previous, current int) bool) ([]int, []int) {
	n := len(nums)
	left, right := make([]int, n), make([]int, n)
	stack := []int{}
	for i, x := range nums {
		right[i] = n
		for len(stack) > 0 && pop(nums[stack[len(stack)-1]], x) {
			j := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			right[j] = i
		}
		left[i] = -1
		if len(stack) > 0 {
			left[i] = stack[len(stack)-1]
		}
		stack = append(stack, i)
	}
	return left, right
}

func subArrayRanges(nums []int) int64 {
	minLeft, minRight := monotonicBoundaries(nums, func(previous, current int) bool { return previous >= current })
	maxLeft, maxRight := monotonicBoundaries(nums, func(previous, current int) bool { return previous <= current })
	var result int64
	for i, value := range nums {
		minima := int64(i-minLeft[i]) * int64(minRight[i]-i)
		maxima := int64(i-maxLeft[i]) * int64(maxRight[i]-i)
		result += int64(value) * (maxima - minima)
	}
	return result
}
