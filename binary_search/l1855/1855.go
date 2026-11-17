package l1855

// 1855：复用本文件的边界模板；特殊结构按对应排除不变量处理。
func maxDistance(nums1 []int, nums2 []int) int {
	answer := 0
	for i, value := range nums1 {
		if i >= len(nums2) {
			break
		}
		end := i + FirstTrue(len(nums2)-i, func(offset int) bool { return nums2[i+offset] < value })
		if end-1-i > answer {
			answer = end - 1 - i
		}
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
