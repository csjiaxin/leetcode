package l85

// https://leetcode.com/problems/maximal-rectangle/
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

func maximalRectangle(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}
	heights := make([]int, len(matrix[0]))
	best := 0
	for _, row := range matrix {
		for j, value := range row {
			if value == '1' {
				heights[j]++
			} else {
				heights[j] = 0
			}
		}
		if area := histogramArea(heights); area > best {
			best = area
		}
	}
	return best
}
