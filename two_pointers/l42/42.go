package l42

// Opposite 维护闭区间 [left,right]；step 必须收缩区间且不扩大边界。
// 指针不得访问区间之外；返回收缩后的边界，必要时可再处理剩余中心。
func Opposite(left, right int, step func(l, r int) (int, int)) (int, int) {
	for left < right {
		left, right = step(left, right)
	}
	return left, right
}

// https://leetcode.com/problems/trapping-rain-water/
// 先结算当前较低边界；该侧已扫描最大高度足以确定本格水量。
func trap(height []int) int {
	leftMax, rightMax, answer := 0, 0, 0
	Opposite(0, len(height)-1, func(l, r int) (int, int) {
		if height[l] <= height[r] {
			leftMax = max(leftMax, height[l])
			answer += leftMax - height[l]
			return l + 1, r
		}
		rightMax = max(rightMax, height[r])
		answer += rightMax - height[r]
		return l, r - 1
	})
	return answer
}
