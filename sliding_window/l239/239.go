package l239

// https://leetcode.com/problems/sliding-window-maximum/
// The standalone helper below matches the named subtemplate in README.md.
// windowMaximum uses a decreasing deque of indices with explicit expiration.
func windowMaximum(nums []int, k int) []int {
	result := []int{}
	if k <= 0 || k > len(nums) {
		return result
	}
	deque := []int{}
	for right, value := range nums {
		if len(deque) > 0 && deque[0] <= right-k {
			deque = deque[1:]
		}
		for len(deque) > 0 && nums[deque[len(deque)-1]] <= value {
			deque = deque[:len(deque)-1]
		}
		deque = append(deque, right)
		if right+1 >= k {
			result = append(result, nums[deque[0]])
		}
	}
	return result
}

func maxSlidingWindow(nums []int, k int) []int { return windowMaximum(nums, k) }
