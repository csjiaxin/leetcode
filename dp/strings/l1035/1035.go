// LeetCode 1035: https://leetcode.com/problems/uncrossed-lines/
package l1035

// dp[i][j] 表示两个数组前缀之间最多能连接的不相交线数。
func maxUncrossedLines(nums1 []int, nums2 []int) int {
	table := Table(len(nums1), len(nums2),
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

// Table 构建 (m+1) × (n+1) 的前缀状态表。
//
// dp[i][j] 对应第一个序列的前 i 个元素与第二个序列的前 j 个元素。
// base 只负责第 0 行、第 0 列；step 负责其余格子。
// step 的 up、left、diagonal 分别是 dp[i-1][j]、dp[i][j-1]、dp[i-1][j-1]。
//
// m、n 必须非负，回调必须在需要调用时有效。回调通过闭包读取原序列。
// 当回调为 O(1) 时，时间、空间均为 O((m+1)*(n+1))。
func Table[T any](
	m, n int,
	base func(i, j int) T,
	step func(i, j int, up, left, diagonal T) T,
) [][]T {
	dp := make([][]T, m+1)
	for i := range dp {
		dp[i] = make([]T, n+1)
	}

	for i := 0; i <= m; i++ {
		dp[i][0] = base(i, 0)
	}
	for j := 1; j <= n; j++ {
		dp[0][j] = base(0, j)
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			dp[i][j] = step(i, j, dp[i-1][j], dp[i][j-1], dp[i-1][j-1])
		}
	}
	return dp
}
