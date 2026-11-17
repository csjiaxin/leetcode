package l11

// https://leetcode.com/problems/container-with-most-water/
func maxArea(height []int) int {
	l := 0
	r := len(height) - 1
	ret := 0

	for l < r {
		lv := height[l]
		rv := height[r]
		area := min(lv, rv) * (r - l)
		ret = max(ret, area)
		if lv < rv {
			for ; l < r && height[l] <= lv; l++ {
			}
		} else {
			for ; l < r && height[r] <= rv; r-- {

			}
		}
	}
	return ret
}
