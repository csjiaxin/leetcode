package l198

// Independent 维护旧前缀的“不选末项 / 选末项”收益，允许全部不选。
func Independent(nums []int) int {
	skip, take := 0, 0
	for _, value := range nums {
		skip, take = max(skip, take), skip+value
	}
	return max(skip, take)
}

// 相邻房屋不能同时选。
func rob(nums []int) int { return Independent(nums) }
