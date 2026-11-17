package l213

// Independent 维护旧前缀的“不选末项 / 选末项”收益，允许全部不选。
func Independent(nums []int) int {
	skip, take := 0, 0
	for _, value := range nums {
		skip, take = max(skip, take), skip+value
	}
	return max(skip, take)
}

// 环拆成不选尾项与不选首项，取两者最大值。
func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	return max(Independent(nums[:len(nums)-1]), Independent(nums[1:]))
}
