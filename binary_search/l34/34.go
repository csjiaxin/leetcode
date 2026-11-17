package l34

// 34：复用本文件的边界模板；特殊结构按对应排除不变量处理。
func searchRange(nums []int, target int) []int {
	left := FirstTrue(len(nums), func(i int) bool { return nums[i] >= target })
	right := FirstTrue(len(nums), func(i int) bool { return nums[i] > target })
	if left == right {
		return []int{-1, -1}
	}
	return []int{left, right - 1}
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
