package l167

// Opposite 维护闭区间 [left,right]；step 必须收缩区间且不扩大边界。
// 指针不得访问区间之外；返回收缩后的边界，必要时可再处理剩余中心。
func Opposite(left, right int, step func(l, r int) (int, int)) (int, int) {
	for left < right {
		left, right = step(left, right)
	}
	return left, right
}

// https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/
// 和偏小排除左端，和偏大排除右端；答案下标从 1 开始。
func twoSum(numbers []int, target int) []int {
	result := []int{}
	Opposite(0, len(numbers)-1, func(l, r int) (int, int) {
		sum := int64(numbers[l]) + int64(numbers[r])
		if sum < int64(target) {
			return l + 1, r
		}
		if sum > int64(target) {
			return l, r - 1
		}
		result = []int{l + 1, r + 1}
		return r, r
	})
	return result
}
