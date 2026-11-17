// Package binarysearch 提供单调谓词的边界查找模板。
package binarysearch

// FirstTrue 在 [0,n) 上寻找第一个使 predicate 为真的下标，无解返回 n。
// 前提：n >= 0，predicate 的结果按下标呈 false...false,true...true。
// 不会调用 predicate(n)，空范围不调用 predicate。
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
