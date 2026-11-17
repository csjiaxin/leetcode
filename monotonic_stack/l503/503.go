package l503

// https://leetcode.com/problems/next-greater-element-ii/
// Embedded subtemplate: see README.md for its invariant and tie policy.
// nextGreaterPositions returns the first strictly greater position, or -1.
// A circular second pass resolves pending originals without pushing copies.
func nextGreaterPositions(nums []int, circular bool) []int {
	n := len(nums)
	next := make([]int, n)
	for i := range next {
		next[i] = -1
	}
	stack := []int{}
	rounds := n
	if circular {
		rounds *= 2
	}
	for step := 0; step < rounds; step++ {
		current := step % n
		for len(stack) > 0 && nums[stack[len(stack)-1]] < nums[current] {
			previous := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			next[previous] = current
		}
		if step < n {
			stack = append(stack, current)
		}
	}
	return next
}

func nextGreaterElements(nums []int) []int {
	next := nextGreaterPositions(nums, true)
	result := make([]int, len(next))
	for i, j := range next {
		result[i] = -1
		if j >= 0 {
			result[i] = nums[j]
		}
	}
	return result
}
