package l287

// https://leetcode.com/problems/find-the-duplicate-number/
func findDuplicate(nums []int) int {
	s, f := 0, 0

	for {
		s = nums[s]
		f = nums[nums[f]]

		if s == f {
			p := 0
			for p != s {
				p = nums[p]
				s = nums[s]
			}
			return p
		}
	}
	return -1
}
