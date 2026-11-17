package l18

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

// https://leetcode.com/problems/4sum/
// 固定两个不同下标后做相向查找，所有层次跳过重复值。
func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	answer := [][]int{}
	for i := 0; i+3 < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j+2 < len(nums); j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			Opposite(j+1, len(nums)-1, func(l, r int) (int, int) {
				sum := int64(nums[i]) + int64(nums[j]) + int64(nums[l]) + int64(nums[r])
				if sum < int64(target) {
					return l + 1, r
				}
				if sum > int64(target) {
					return l, r - 1
				}
				answer = append(answer, []int{nums[i], nums[j], nums[l], nums[r]})
				l++
				r--
				for l < r && nums[l] == nums[l-1] {
					l++
				}
				for l < r && nums[r] == nums[r+1] {
					r--
				}
				return l, r
			})
		}
	}
	return answer
}
