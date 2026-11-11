package l31

import "slices"

// https://leetcode.com/problems/next-permutation/
// nextPermutation modifies nums in place to its next lexicographic permutation.
// If nums is the largest permutation, it becomes the smallest permutation.
func nextPermutation(nums []int) {
	N := len(nums)
	i := N - 1
	for ; i > 0; i-- {
		if nums[i-1] < nums[i] {
			// i -1 , i
			for j := N - 1; j >= i; j-- {
				if nums[j] > nums[i-1] {
					//swap i-1 with j
					nums[i-1], nums[j] = nums[j], nums[i-1]
					//reverse [i:]
					slices.Reverse(nums[i:])
					return
				}
			}
		}
	}
	if i == 0 {
		slices.Reverse(nums)
	}
}
