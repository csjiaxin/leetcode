package l713

// https://leetcode.com/problems/subarray-product-less-than-k/
// numSubarrayProductLessThanK counts nonempty contiguous subarrays whose product
// is strictly less than k. All elements are positive integers.
func numSubarrayProductLessThanK(nums []int, k int) int {
	window := 1
	N := len(nums)
	left := 0
	ret := 0
	for right := 0; right < N; right++ {
		window *= nums[right]

		for left <= right && window >= k {
			window /= nums[left]
			left += 1
		}
		ret += right - left + 1
	}
	return ret
}
