package main

// Standalone template plus public LeetCode adapter; README.md describes the invariant.
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

func maxSumMinProduct(nums []int) int {
	prefix := make([]int64, len(nums)+1)
	for i, value := range nums {
		prefix[i+1] = prefix[i] + int64(value)
	}
	left, right := monotonicBoundaries(nums, func(previous, current int) bool { return previous >= current })
	var best int64
	for i, value := range nums {
		product := int64(value) * (prefix[right[i]] - prefix[left[i]+1])
		if product > best {
			best = product
		}
	}
	return int(best % 1_000_000_007)
}

func main() {}
