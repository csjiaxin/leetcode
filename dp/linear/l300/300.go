package main

// PrefixTable 的 dp[0] 是空前缀；step(i, dp[:i]) 只能读取已完成状态。
func PrefixTable[T any](n int, empty T, step func(i int, previous []T) T) []T {
	dp := make([]T, n+1)
	dp[0] = empty
	for i := 1; i <= n; i++ {
		dp[i] = step(i, dp[:i])
	}
	return dp
}

// dp[i] 是以 nums[i-1] 结尾的长度，答案要汇总全部结尾；O(n²)。
func lengthOfLIS(nums []int) int {
	answer := 0
	PrefixTable(len(nums), 0, func(i int, old []int) int {
		best := 1
		for j := 1; j < i; j++ {
			if nums[j-1] < nums[i-1] {
				best = max(best, old[j]+1)
			}
		}
		answer = max(answer, best)
		return best
	})
	return answer
}
