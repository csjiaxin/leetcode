package l1031

// https://leetcode.com/problems/maximum-sum-of-two-non-overlapping-subarrays/
// Time: O(n). Extra space: O(n).
func maxSumTwoNoOverlap(nums []int, firstLen int, secondLen int) int {
	n := len(nums)
	sum := make([]int, n)
	sum[0] = nums[0]
	for i := 1; i < n; i++ {
		sum[i] = sum[i-1] + nums[i]
	}
	//[i:j]
	getWindowSum := func(i, j int) int {
		ret := sum[j]
		if i > 0 {
			ret -= sum[i-1]
		}
		return ret
	}
	maxSumTwo := func(leftLen, rightLen int) int {
		maxLeftWindowSum := getWindowSum(0, leftLen-1)
		ret := getWindowSum(0, rightLen+leftLen-1)
		for j := leftLen + rightLen; j < n; j++ {

			rightWindowStart := j + 1 - rightLen
			newLeftWindowEnd := rightWindowStart - 1
			newLeftWindowStart := newLeftWindowEnd + 1 - leftLen
			newLeftWindowSum := getWindowSum(newLeftWindowStart, newLeftWindowEnd)
			maxLeftWindowSum = max(maxLeftWindowSum, newLeftWindowSum)
			rightWindowSum := getWindowSum(rightWindowStart, j)
			ret = max(ret, maxLeftWindowSum+rightWindowSum)
		}
		return ret
	}
	return max(maxSumTwo(firstLen, secondLen), maxSumTwo(secondLen, firstLen))
}
