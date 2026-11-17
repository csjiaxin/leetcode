package l84

// https://leetcode.com/problems/largest-rectangle-in-histogram/
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

func histogramArea(heights []int) int {
	left, right := monotonicBoundaries(heights, func(previous, current int) bool { return previous > current })
	var best int64
	for i, h := range heights {
		area := int64(h) * int64(right[i]-left[i]-1)
		if area > best {
			best = area
		}
	}
	return int(best)
}

func largestRectangleArea(heights []int) int { return histogramArea(heights) }
