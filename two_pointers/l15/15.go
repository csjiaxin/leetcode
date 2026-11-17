package l15

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

// https://leetcode.com/problems/3sum/
// 排序固定第一个数，剩余区间寻找有序数对；固定数和左右端都去重。
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	answer := [][]int{}
	for i := 0; i+2 < len(nums); i++ {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		Opposite(i+1, len(nums)-1, func(l, r int) (int, int) {
			sum := int64(nums[i]) + int64(nums[l]) + int64(nums[r])
			if sum < 0 {
				return l + 1, r
			}
			if sum > 0 {
				return l, r - 1
			}
			answer = append(answer, []int{nums[i], nums[l], nums[r]})
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
	return answer
}
