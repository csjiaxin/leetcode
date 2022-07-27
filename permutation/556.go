package main

import (
	"fmt"
	"math"
)

func reverse(nums []int, i int) {
	j := len(nums) - 1
	for ; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}

func nextGreaterElement(n int) int {
	nums := make([]int, 0)
	for n > 0 {
		nums = append(nums, n%10)
		n /= 10
	}
	reverse(nums, 0)

	n = len(nums)
	i := n - 2
	for ; i >= 0 && nums[i] >= nums[i+1]; i-- {
	}
	if i < 0 {
		return -1
	} else {
		j := n - 1
		for ; j >= 0 && nums[j] <= nums[i]; j-- {
		}
		nums[i], nums[j] = nums[j], nums[i]
		reverse(nums, i+1)
		N := 0
		f := 1
		for j := n - 1; j >= 0; j-- {
			N += nums[j] * f
			f *= 10
		}
		if N > math.MaxInt32 {
			return -1
		}
		return N
	}
}

func main() {
	fmt.Println(nextGreaterElement(2147483486))
}
