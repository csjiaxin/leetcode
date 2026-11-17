package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func col_max(nums [][]int, col int) ([]int, int) {
	m := 0
	r := []int{}
	for i := 0; i < len(nums); i++ {
		if nums[i][col] > m {
			m = nums[i][col]
			r = []int{i, col}
		}
	}
	return r, m
}
func findPeakGrid(nums [][]int) []int {

	r := []int{}
	l := 0
	h := len(nums[0]) - 1
	for l < h {
		m := (l + h) / 2
		ma, mv := col_max(nums, m)
		m1a, m1v := col_max(nums, m+1)
		if mv <= m1v {
			l = m + 1
			r = m1a
		} else {
			r = ma
			h = m
		}
	}
	return r
}

func main() {
	fmt.Println(findPeakGrid([][]int{{10, 20, 15}, {21, 30, 14}, {7, 16, 32}}))
}
