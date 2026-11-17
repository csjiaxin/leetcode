package l977

// Opposite 维护闭区间 [left,right]；step 必须收缩区间且不扩大边界。
// 指针不得访问区间之外；返回收缩后的边界，必要时可再处理剩余中心。
func Opposite(left, right int, step func(l, r int) (int, int)) (int, int) {
	for left < right {
		left, right = step(left, right)
	}
	return left, right
}

// https://leetcode.com/problems/squares-of-a-sorted-array/
// 比较两端平方，从结果末尾填最大值，最后补剩余中心。
func sortedSquares(nums []int) []int {
	answer := make([]int, len(nums))
	write := len(nums) - 1
	l, _ := Opposite(0, len(nums)-1, func(l, r int) (int, int) {
		a, b := nums[l]*nums[l], nums[r]*nums[r]
		if a > b {
			answer[write] = a
			write--
			return l + 1, r
		}
		answer[write] = b
		write--
		return l, r - 1
	})
	if len(nums) > 0 {
		answer[0] = nums[l] * nums[l]
	}
	return answer
}
