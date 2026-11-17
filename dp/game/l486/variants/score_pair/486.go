package l486

type E struct {
	firstMax  int
	secondMax int
}

// https://leetcode.com/problems/predict-the-winner/
func predictTheWinner(nums []int) bool {
	N := len(nums)
	dp := make([][]E, N)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]E, N)
		dp[i][i] = E{
			firstMax:  nums[i],
			secondMax: 0,
		}
	}

	for i := N - 2; i >= 0; i-- {
		for j := i + 1; j < N; j++ {
			if j == i+1 {
				dp[i][j] = E{
					firstMax:  max(nums[i], nums[j]),
					secondMax: min(nums[i], nums[j]),
				}
			} else {
				left := nums[i] + dp[i+1][j].secondMax
				right := nums[j] + dp[i][j-1].secondMax
				if left >= right {
					dp[i][j] = E{
						firstMax:  left,
						secondMax: dp[i+1][j].firstMax,
					}
				} else {
					dp[i][j] = E{
						firstMax:  right,
						secondMax: dp[i][j-1].firstMax,
					}
				}
			}
		}
	}
	return dp[0][N-1].firstMax >= dp[0][N-1].secondMax
}
