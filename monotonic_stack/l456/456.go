package l456

// https://leetcode.com/problems/132-pattern/
// Embedded subtemplate: see README.md for its invariant and tie policy.
func has132(nums []int) bool {
	stack := []int{}
	third := -int(^uint(0)>>1) - 1
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] < third {
			return true
		}
		for len(stack) > 0 && stack[len(stack)-1] < nums[i] {
			third = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, nums[i])
	}
	return false
}

func find132pattern(nums []int) bool { return has132(nums) }
