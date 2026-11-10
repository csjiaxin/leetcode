package l1306

// https://leetcode.com/problems/jump-game-iii/
func canReach(arr []int, start int) bool {

	n := len(arr)
	dp := make([]int, n)

	var dfs func(i int)
	dfs = func(i int) {
		if i < 0 || i >= n {
			return
		}

		if dp[i] == 0 {
			dp[i] = 1

			dfs(i + arr[i])
			dfs(i - arr[i])
		}
	}
	dfs(start)

	for i := 0; i < n; i++ {
		if arr[i] == 0 && dp[i] == 1 {
			return true
		}
	}
	return false
}
