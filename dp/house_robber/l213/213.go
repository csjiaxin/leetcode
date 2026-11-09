package l213

// rob returns the maximum amount that can be robbed without robbing adjacent
// houses. The houses form a circle, so the first and last houses are adjacent.
func rob(nums []int) int {
	L := len(nums)
	dp0 := make([]int, L)
	dp1 := make([]int, L)

	dp0[0] = 0
	dp1[0] = nums[0]
	for i := 1; i < L; i++ {
		dp0[i] = max(dp0[i-1], dp1[i-1])
		dp1[i] = dp0[i-1] + nums[i]
	}
	return max(dp0[L-1], dp1[L-1])
}
