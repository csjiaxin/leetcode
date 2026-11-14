package l80

// https://leetcode.com/problems/remove-duplicates-from-sorted-array-ii/
func removeDuplicates(nums []int) int {
	write := 0
	for read, v := range nums {
		if write >= 2 && nums[write-1] == nums[write-2] && v == nums[write-1] {
			continue
		}
		nums[read], nums[write] = nums[write], nums[read]
		write++
	}
	return write
}
