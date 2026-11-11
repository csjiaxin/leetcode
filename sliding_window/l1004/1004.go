package l1004

// https://leetcode.com/problems/max-consecutive-ones-iii/
// longestOnes returns the longest run of ones after flipping at most k zeros
// in the binary array nums.
func longestOnes(nums []int, k int) int {
	N := len(nums)
	left := 0
	count := [2]int{}
	ret := 0
	for right := 0; right < N; right++ {
		c := nums[right]
		count[c] += 1

		for count[0] > k {
			count[nums[left]] -= 1
			left += 1
		}
		ret = max(ret, right-left+1)

	}
	return ret
}
