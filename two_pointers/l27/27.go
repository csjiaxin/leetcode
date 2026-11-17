package l27

// Compact 稳定保留满足 keep 的元素，nums[:write] 是已经完成的结果。
// keep 只能读取 out，不得修改它或尚未扫描的元素；返回有效前缀长度。
func Compact(nums []int, keep func(out []int, value int) bool) int {
	write := 0
	for _, value := range nums {
		if keep(nums[:write], value) {
			nums[write] = value
			write++
		}
	}
	return write
}

// https://leetcode.com/problems/remove-element/
// 保留不等于 val 的元素。
func removeElement(nums []int, val int) int {
	return Compact(nums, func(_ []int, value int) bool { return value != val })
}
