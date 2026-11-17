package l518

// https://leetcode.com/problems/coin-change-ii/

func change(amount int, coins []int) int {
	dp := make([]int, amount+1)

	dp[0] = 1
	for _, c := range coins {
		for v := 0; v+c <= amount; v++ {
			newV := v + c
			dp[newV] += dp[v]
		}
	}
	return dp[amount]
}
