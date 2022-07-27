package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
func maxArea(height []int) int {
	l := 0
	r := len(height) - 1

	ret := 0
	for l < r {
		minV := min(height[l], height[r])
		ret = max(ret, minV*(r-l))
		if height[l] == minV {
			l += 1
		} else {
			r -= 1
		}

	}
	return ret
}

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}
