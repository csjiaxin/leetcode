package l11

// Opposite 维护闭区间 [left,right]；step 必须收缩区间且不扩大边界。
// 指针不得访问区间之外；返回收缩后的边界，必要时可再处理剩余中心。
func Opposite(left, right int, step func(l, r int) (int, int)) (int, int) {
	for left < right {
		left, right = step(left, right)
	}
	return left, right
}

// https://leetcode.com/problems/container-with-most-water/
// 面积受短板限制，移动短板才可能改善高度上界。
func maxArea(height []int) int {
	answer := 0
	Opposite(0, len(height)-1, func(l, r int) (int, int) {
		answer = max(answer, (r-l)*min(height[l], height[r]))
		if height[l] <= height[r] {
			return l + 1, r
		}
		return l, r - 1
	})
	return answer
}
