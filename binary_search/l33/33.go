package l33

// 33：复用本文件的边界模板；特殊结构按对应排除不变量处理。
func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	n := len(nums)
	pivot := FirstTrue(n, func(i int) bool { return nums[i] <= nums[n-1] })
	left, right := 0, pivot
	if target <= nums[n-1] {
		left, right = pivot, n
	}
	index := left + FirstTrue(right-left, func(i int) bool { return nums[left+i] >= target })
	if index < right && nums[index] == target {
		return index
	}
	return -1
}

func FirstTrue(n int, predicate func(int) bool) int {
	left, right := 0, n-1
	for left <= right {
		mid := left + (right-left)/2
		if predicate(mid) {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}
