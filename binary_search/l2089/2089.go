package l2089

import "sort"

// 2089：复用本文件的边界模板；特殊结构按对应排除不变量处理。
func targetIndices(nums []int, target int) []int {
	values := append([]int(nil), nums...)
	sort.Ints(values)
	left := FirstTrue(len(values), func(i int) bool { return values[i] >= target })
	right := FirstTrue(len(values), func(i int) bool { return values[i] > target })
	answer := make([]int, 0, right-left)
	for i := left; i < right; i++ {
		answer = append(answer, i)
	}
	return answer
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
