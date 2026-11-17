package l907

// https://leetcode.com/problems/sum-of-subarray-minimums/
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

func sumSubarrayMins(arr []int) int {
	left, right := monotonicBoundaries(arr, func(previous, current int) bool { return previous >= current })
	const mod int64 = 1_000_000_007
	var result int64
	for i, value := range arr {
		result = (result + int64(value)*int64(i-left[i])*int64(right[i]-i)) % mod
	}
	return int(result)
}
