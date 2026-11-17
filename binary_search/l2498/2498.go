package l2498

// 2498：复用本文件的边界模板；特殊结构按对应排除不变量处理。
func maxJump(stones []int) int {
	n := len(stones)
	if n < 2 {
		return 0
	}
	return FirstTrue(stones[n-1]-stones[0]+1, func(limit int) bool {
		if stones[1]-stones[0] > limit || stones[n-1]-stones[n-2] > limit {
			return false
		}
		for i := 2; i < n; i++ {
			if stones[i]-stones[i-2] > limit {
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
