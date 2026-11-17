package l410

// 410：复用本文件的边界模板；特殊结构按对应排除不变量处理。
func splitArray(nums []int, k int) int {
	low, high := 0, 0
	for _, value := range nums {
		if value > low {
			low = value
		}
		high += value
	}
	return low + FirstTrue(high-low+1, func(offset int) bool {
		limit := low + offset
		parts, total := 1, 0
		for _, value := range nums {
			if total+value > limit {
				parts++
				total = 0
			}
			total += value
		}
		return parts <= k
	})
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
