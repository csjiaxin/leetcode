package l905

// Partition 把满足 keepLeft 的值放到左区，返回边界；不保证稳定。
func Partition(nums []int, keepLeft func(int) bool) int {
	left, right := 0, len(nums)-1
	for left <= right {
		if keepLeft(nums[left]) {
			left++
			continue
		}
		nums[left], nums[right] = nums[right], nums[left]
		right--
	}
	return left
}

// https://leetcode.com/problems/sort-array-by-parity/
// 偶数左分区，奇数右分区；右端交换来的元素尚未判定，需继续检查。
func sortArrayByParity(nums []int) []int {
	Partition(nums, func(value int) bool { return value%2 == 0 })
	return nums
}
