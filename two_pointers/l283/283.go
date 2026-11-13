package l283

// https://leetcode.com/problems/move-zeroes/
func moveZeroes(nums []int) {
	l, r := 0, 0
	N := len(nums)

	for l < N && r < N {
		for ; l < N && nums[l] == 0; l++ {

		}
		for ; r < N && nums[r] != 0; r++ {

		}
		if l < r {
			l++
		} else if l < N && r < N {
			nums[l], nums[r] = nums[r], nums[l]
			l++
			r++
		}
	}
}
