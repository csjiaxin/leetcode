package l209

import "math"

// https://leetcode.com/problems/minimum-size-subarray-sum/
// minSubArrayLen returns the shortest contiguous subarray length whose sum is
// at least target, or 0 if none exists. All elements and target are positive.
func minSubArrayLen(target int, nums []int) int {

	window := 0
	N := len(nums)
	left := 0
	ret := math.MaxInt
	for right := 0; right < N; right++ {
		window += nums[right]

		for window >= target {
			ret = min(ret, right-left+1)
			window -= nums[left]
			left += 1
		}
	}
	if ret == math.MaxInt {
		return 0
	}
	return ret
}
