package l496

// https://leetcode.com/problems/next-greater-element-i/
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

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	next := nextGreaterPositions(nums2, false)
	values := map[int]int{}
	for i, j := range next {
		values[nums2[i]] = -1
		if j >= 0 {
			values[nums2[i]] = nums2[j]
		}
	}
	result := make([]int, len(nums1))
	for i, x := range nums1 {
		result[i] = values[x]
	}
	return result
}
