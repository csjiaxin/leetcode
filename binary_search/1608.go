package main

import (
	"fmt"
	"sort"
)

func specialArray(nums []int) int {
	n := len(nums)
	sort.Slice(nums, func(i, j int) bool { return nums[i] > nums[j] })

	l := 0
	r := n - 1
	for l <= r {
		m := l + (r-l)/2
		if nums[m] < m {
			if m-1 >= 0 && nums[m-1] >= m {
				return m
			}
			l = m + 1
		} else {
			r = m - 1
		}
	}
	if l == n {
		return n
	}
	return -1
}

func main() {
	fmt.Println(specialArray([]int{3, 5})) //2
}
