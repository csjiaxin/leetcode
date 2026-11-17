package l1760

// 1760：复用本文件的边界模板；特殊结构按对应排除不变量处理。
func minimumSize(nums []int, maxOperations int) int {
	highest := 0
	for _, value := range nums {
		if value > highest {
			highest = value
		}
	}
	return 1 + FirstTrue(highest, func(offset int) bool {
		size := offset + 1
		operations := int64(0)
		for _, value := range nums {
			operations += int64((value - 1) / size)
			if operations > int64(maxOperations) {
				return false
			}
		}
		return true
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
