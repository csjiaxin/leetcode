package l15

import "slices"

// https://leetcode.com/problems/3sum/
// threeSum sorts nums in place and returns unique triplets whose sum is zero.
func threeSum(nums []int) [][]int {
	slices.Sort(nums)
	ans := make([][]int, 0)

	for i := 0; i+2 < len(nums); i++ {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		l, r := i+1, len(nums)-1
		for l < r {
			sum := int64(nums[i]) + int64(nums[l]) + int64(nums[r])
			switch {
			case sum < 0:
				l++
			case sum > 0:
				r--
			default:
				ans = append(ans, []int{nums[i], nums[l], nums[r]})
				l++
				r--
				// Skip repeated values so this fixed nums[i] cannot emit duplicates.
				for l < r && nums[l] == nums[l-1] {
					l++
				}
				for l < r && nums[r] == nums[r+1] {
					r--
				}
			}
		}
	}
	return ans
}
