package main

import "fmt"

func canPartition(nums []int) bool {
	s := 0
	for _, v := range nums {
		s += v
	}
	if s%2 != 0 {
		return false
	}
	s /= 2
	dp := make([]bool, s+1)
	dp[0] = true
	for i := 0; i < len(nums); i++ {
		for j := s; j >= nums[i]; j-- {
			dp[j] = dp[j] || dp[j-nums[i]]
		}
	}
	return dp[s]
}
func main() {

	fmt.Println(canPartition([]int{1, 5, 11, 5}))
}
