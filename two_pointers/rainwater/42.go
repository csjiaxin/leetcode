package main

import "fmt"

func trap(height []int) int {
	n := len(height)
	l := 0
	r := n - 1
	lmax := 0
	rmax := 0
	ret := 0
	for l < r {
		if height[l] < height[r] {
			if height[l] < lmax {
				ret += lmax - height[l]
			} else {
				lmax = height[l]
			}
			l += 1
		} else {
			if height[r] < rmax {
				ret += rmax - height[r]
			} else {

				rmax = height[r]
			}
			r -= 1
		}
	}
	return ret
}

func main() {
	fmt.Println(trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
}
