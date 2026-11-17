package l55

// RangeJumps 把连续可达区间作为 BFS 层；无法推进时返回 -1。
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

// https://leetcode.com/problems/jump-game/
// 按连续可达区间推进，只关心是否存在到达最后位置的层。
func canJump(nums []int) bool { return RangeJumps(nums) >= 0 }
