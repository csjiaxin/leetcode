package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func right_greater(nums []int, v int) int {
	l := 0
	h := len(nums) - 1

	for l <= h {
		m := (l + h) / 2
		if nums[m] == v {
			l = m + 1
		} else if nums[m] < v {
			h = m - 1
		} else {
			l = m + 1
		}
	}
	return h
}
func maxDistance(nums1 []int, nums2 []int) int {

	m := 0
	for i, v := range nums1 {

		j := right_greater(nums2, v)

		m = max(m, j-i)
	}
	return m

}


func main() {
	fmt.Println(maxDistance([]int{55, 30, 5, 4, 2}, []int{100, 20, 10, 10, 5}))
}
