// LeetCode 1035: https://leetcode.com/problems/uncrossed-lines/
package l1035

import prefixdp "leetcode/dp/strings"

// dp[i][j] 表示两个数组前缀之间最多能连接的不相交线数。
func maxUncrossedLines(nums1 []int, nums2 []int) int {
	table := prefixdp.Table(len(nums1), len(nums2),
		func(i, j int) int { return 0 },
		func(i, j int, up, left, diagonal int) int {
			if nums1[i-1] == nums2[j-1] {
				return diagonal + 1
			}
			if up > left {
				return up
			}
			return left
		})
	return table[len(nums1)][len(nums2)]
}
