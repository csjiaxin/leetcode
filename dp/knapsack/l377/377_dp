package main

import "fmt"

func Init(m, n, dvalue int) [][]int {
	dp := make([][]int, m)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, n)
		for j := 0; j < len(dp[i]); j++ {
			dp[i][j] = dvalue
		}
	}
	return dp
}

func combinationSum4(nums []int, target int) int {
	n := len(nums)
	dp := make([]int, target+1)
	dp[0] = 1
	for t := 1; t <= target; t++ {
		for i := 0; i < n; i++ {
			// for k := 1; k*nums[i] <= t; k++ {
			if nums[i] <= t {
				dp[t] += dp[t-nums[i]]
			}
			// }
		}
	}
	return dp[target]
}

func main() {
	fmt.Println(combinationSum4([]int{1, 2, 3}, 4))
}
