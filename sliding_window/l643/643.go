package l643

import "math"

// https://leetcode.com/problems/maximum-average-subarray-i/
// findMaxAverage returns the maximum average of any contiguous subarray of length k.
func findMaxAverage(nums []int, k int) float64 {

	left := 0
	N := len(nums)
	right := 0
	total := 0.0
	ret := float64(math.MinInt64)

	for ; right < N; right++ {
		total += float64(nums[right])
		if right-left+1 > k {
			total -= float64(nums[left])
			left++
		}
		if right-left+1 == k {
			ret = max(ret, total/float64(k))
		}

	}
	return ret
}
