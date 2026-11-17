package l45

// RangeJumps 将连续可达区间作为 BFS 层；边界无法推进时返回 -1。
func RangeJumps(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}
	end, farthest, steps := 0, 0, 0
	for i := 0; i < len(nums)-1; i++ {
		if i > farthest {
			return -1
		}
		farthest = max(farthest, i+nums[i])
		if i == end {
			if farthest == end {
				return -1
			}
			end = farthest
			steps++
			if end >= len(nums)-1 {
				return steps
			}
		}
	}
	return -1
}

// https://leetcode.com/problems/jump-game-ii/
// 扫描完整可达层后才增加一次跳跃。
func jump(nums []int) int { return RangeJumps(nums) }
