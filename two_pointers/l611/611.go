package l611

import (
	"sort"
)

// Opposite 维护闭区间 [left,right]；step 必须收缩区间且不扩大边界。
// 指针不得访问区间之外；返回收缩后的边界，必要时可再处理剩余中心。
func Opposite(left, right int, step func(l, r int) (int, int)) (int, int) {
	for left < right {
		left, right = step(left, right)
	}
	return left, right
}

// https://leetcode.com/problems/valid-triangle-number/
// 固定最大边，条件成立时一次计入 l..r-1 这一整段第一条边。
func triangleNumber(nums []int) int {
	sort.Ints(nums)
	answer := 0
	for largest := 2; largest < len(nums); largest++ {
		Opposite(0, largest-1, func(l, r int) (int, int) {
			if int64(nums[l])+int64(nums[r]) > int64(nums[largest]) {
				answer += r - l
				return l, r - 1
			}
			return l + 1, r
		})
	}
	return answer
}
