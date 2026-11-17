package l888

import "sort"

// 888：复用本文件的边界模板；特殊结构按对应排除不变量处理。
func fairCandySwap(aliceSizes, bobSizes []int) []int {
	sumA, sumB := 0, 0
	for _, value := range aliceSizes {
		sumA += value
	}
	for _, value := range bobSizes {
		sumB += value
	}
	if (sumB-sumA)%2 != 0 {
		return []int{}
	}
	values := append([]int(nil), bobSizes...)
	sort.Ints(values)
	for _, a := range aliceSizes {
		b := a + (sumB-sumA)/2
		index := FirstTrue(len(values), func(i int) bool { return values[i] >= b })
		if index < len(values) && values[index] == b {
			return []int{a, b}
		}
	}
	return []int{}
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
