package l1608

import "sort"

// 1608：复用本文件的边界模板；特殊结构按对应排除不变量处理。
func specialArray(nums []int) int {
	values := append([]int(nil), nums...)
	sort.Ints(values)
	count := func(x int) int {
		return len(values) - FirstTrue(len(values), func(i int) bool { return values[i] >= x })
	}
	x := FirstTrue(len(values)+1, func(x int) bool { return count(x) <= x })
	if count(x) == x {
		return x
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
