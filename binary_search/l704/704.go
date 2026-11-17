package l704

// 704：复用本文件的边界模板；特殊结构按对应排除不变量处理。
func search(nums []int, target int) int {
	index := FirstTrue(len(nums), func(i int) bool { return nums[i] >= target })
	if index < len(nums) && nums[index] == target {
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
