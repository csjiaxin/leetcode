package l962

// https://leetcode.com/problems/maximum-width-ramp/
// Embedded subtemplate: see README.md for its invariant and tie policy.
func widestRamp(nums []int) int {
	candidates := []int{}
	for i, value := range nums {
		if len(candidates) == 0 || value < nums[candidates[len(candidates)-1]] {
			candidates = append(candidates, i)
		}
	}
	best := 0
	for right := len(nums) - 1; right >= 0; right-- {
		for len(candidates) > 0 && nums[candidates[len(candidates)-1]] <= nums[right] {
			left := candidates[len(candidates)-1]
			candidates = candidates[:len(candidates)-1]
			if right-left > best {
				best = right - left
			}
		}
	}
	return best
}

func maxWidthRamp(nums []int) int { return widestRamp(nums) }
