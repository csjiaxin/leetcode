package l27

// https://leetcode.com/problems/remove-element/
func removeElement(nums []int, val int) int {
	write := 0
	for read, v := range nums {
		if v == val {
			continue
		}
		if read != write {
			nums[read], nums[write] = nums[write], nums[read]
		}
		write++
	}
	return write
}
