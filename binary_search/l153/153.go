package l153

// 153：复用本文件的边界模板；特殊结构按对应排除不变量处理。
func findMin(nums []int) int {
	last := nums[len(nums)-1]
	index := FirstTrue(len(nums), func(i int) bool { return nums[i] <= last })
	return nums[index]
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
