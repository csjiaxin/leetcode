package l283

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

// https://leetcode.com/problems/move-zeroes/
// 先稳定保留非零，再把尾部补零。
func moveZeroes(nums []int) {
	length := Compact(nums, func(_ []int, value int) bool { return value != 0 })
	for i := length; i < len(nums); i++ {
		nums[i] = 0
	}
}
