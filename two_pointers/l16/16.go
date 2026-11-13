package l16

import "slices"

// https://leetcode.com/problems/3sum-closest/
// threeSumClosest sorts nums in place. The input must contain at least three values.
func threeSumClosest(nums []int, target int) int {
	slices.Sort(nums)
	best := nums[0] + nums[1] + nums[2]
	bestDiff := abs(best - target)

	for i := 0; i+2 < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		l, r := i+1, len(nums)-1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			if diff := abs(sum - target); diff < bestDiff {
				best, bestDiff = sum, diff
			}

			switch {
			case sum < target:
				l++
			case sum > target:
				r--
			default:
				return target
			}
		}
	}
	return best
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
