package l871

// https://leetcode.com/problems/minimum-number-of-refueling-stops/
func minRefuelStops(target int, startFuel int, stations [][]int) int {
	n := len(stations)
	dp := make([]int, n+1)
	dp[0] = startFuel

	for _, val := range stations {
		pos := val[0]
		fuel := val[1]
		for j := n; j > 0; j-- {
			if dp[j-1] >= pos {
				dp[j] = max(dp[j], dp[j-1]+fuel)
			}

		}
	}

	for i := 0; i <= n; i++ {
		if dp[i] >= target {
			return i
		}
	}
	return -1
}
