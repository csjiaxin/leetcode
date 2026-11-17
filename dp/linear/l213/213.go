package l213

// rob returns the maximum amount that can be robbed without robbing adjacent
// houses. The houses form a circle, so the first and last houses are adjacent.
func rob(nums []int) int {
	L := len(nums)
	dp0 := make([]int, L)
	dp1 := make([]int, L)

	dp0[0] = nums[0]
	if L == 1 {
		return nums[0]
	}
	dp0[1] = max(nums[0], nums[1])
	dp1[1] = nums[1]
	for i := 2; i < L; i++ {
		dp0[i] = max(dp0[i-1], dp0[i-2]+nums[i])
		dp1[i] = max(dp1[i-2]+nums[i], dp1[i-1])
	}
	return max(dp0[L-2], dp1[L-1])
}
