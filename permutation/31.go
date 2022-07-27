package main

import "fmt"

func reverse(nums []int, i int) {
	j := len(nums) - 1
	for ; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}

func nextPermutation(nums []int) {
	n := len(nums)
	i := n - 2
	for ; i >= 0 && nums[i] >= nums[i+1]; i-- {
	}
	if i < 0 {
		reverse(nums, 0)
	} else {
		j := n - 1
		for ; j >= 0 && nums[j] <= nums[i]; j-- {
		}
		nums[i], nums[j] = nums[j], nums[i]
		reverse(nums, i+1)
	}
}
func main() {
	s := []int{1, 2, 3}
	nextPermutation(s)
	fmt.Println(s)
}
