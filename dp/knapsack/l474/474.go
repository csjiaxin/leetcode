package main

// TwoResources：每项消耗两种资源，两个容量都倒序，每项只选一次。
func TwoResources(costs [][2]int, zeroLimit, oneLimit int) [][]int {
	dp := make([][]int, zeroLimit+1)
	for z := range dp {
		dp[z] = make([]int, oneLimit+1)
	}
	for _, cost := range costs {
		for z := zeroLimit; z >= cost[0]; z-- {
			for o := oneLimit; o >= cost[1]; o-- {
				dp[z][o] = max(dp[z][o], dp[z-cost[0]][o-cost[1]]+1)
			}
		}
	}
	return dp
}

// 每个字符串消耗零和一两种资源，不是单容量背包。
func findMaxForm(strs []string, m int, n int) int {
	costs := make([][2]int, len(strs))
	for i, s := range strs {
		for _, char := range s {
			if char == '0' {
				costs[i][0]++
			} else {
				costs[i][1]++
			}
		}
	}
	return TwoResources(costs, m, n)[m][n]
}
