package l904

// https://leetcode.com/problems/fruit-into-baskets/
// totalFruit returns the longest contiguous subarray containing at most two
// distinct fruit types.
func totalFruit(nums []int) int {
	N := len(nums)
	left := 0
	count := map[int]int{}
	ret := 0
	for right := 0; right < N; right++ {
		c := nums[right]
		count[c] += 1

		for len(count) > 2 {
			count[nums[left]] -= 1
			if count[nums[left]] == 0 {
				delete(count, nums[left])
			}
			left += 1
		}
		ret = max(ret, right-left+1)

	}
	return ret
}
