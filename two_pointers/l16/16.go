package l16

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

func Distance(a, b int64) int64 {
	if a > b {
		return a - b
	}
	return b - a
}

// https://leetcode.com/problems/3sum-closest/
// 候选和先转 int64；每次访问数对都比较到目标的距离。
func threeSumClosest(nums []int, target int) int {
	if len(nums) < 3 {
		return 0
	}
	sort.Ints(nums)
	desired := int64(target)
	best := int64(nums[0]) + int64(nums[1]) + int64(nums[2])
	for i := 0; i+2 < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		Opposite(i+1, len(nums)-1, func(l, r int) (int, int) {
			sum := int64(nums[i]) + int64(nums[l]) + int64(nums[r])
			if Distance(sum, desired) < Distance(best, desired) {
				best = sum
			}
			if sum < desired {
				return l + 1, r
			}
			if sum > desired {
				return l, r - 1
			}
			return r, r
		})
		if best == desired {
			break
		}
	}
	return int(best)
}
