package l1011

// 1011：复用本文件的边界模板；特殊结构按对应排除不变量处理。
func shipWithinDays(weights []int, days int) int {
	low, high := 0, 0
	for _, weight := range weights {
		if weight > low {
			low = weight
		}
		high += weight
	}
	return low + FirstTrue(high-low+1, func(offset int) bool {
		capacity := low + offset
		needed, load := 1, 0
		for _, weight := range weights {
			if load+weight > capacity {
				needed++
				load = 0
			}
			load += weight
		}
		return needed <= days
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
