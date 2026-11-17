package l875

// 875：复用本文件的边界模板；特殊结构按对应排除不变量处理。
func minEatingSpeed(piles []int, h int) int {
	highest := 0
	for _, pile := range piles {
		if pile > highest {
			highest = pile
		}
	}
	return 1 + FirstTrue(highest, func(offset int) bool {
		speed := offset + 1
		hours := int64(0)
		for _, pile := range piles {
			hours += int64((pile-1)/speed + 1)
			if hours > int64(h) {
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
