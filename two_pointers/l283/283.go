package l283

// https://leetcode.com/problems/move-zeroes/
// moveZeroes moves zeroes to the end in place, preserving nonzero order.
func moveZeroes(nums []int) {
	write := 0
	for read, v := range nums {
		if v == 0 {
			continue
		}
		if read != write {
			nums[read], nums[write] = nums[write], nums[read]
		}
		write++
	}
}
