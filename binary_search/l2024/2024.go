package l2024

// 2024：复用本文件的边界模板；特殊结构按对应排除不变量处理。
func maxConsecutiveAnswers(answerKey string, k int) int {
	feasible := func(length int) bool {
		if length == 0 {
			return true
		}
		ts := 0
		for right := range answerKey {
			if answerKey[right] == 'T' {
				ts++
			}
			if right >= length && answerKey[right-length] == 'T' {
				ts--
			}
			if right+1 >= length && min(ts, length-ts) <= k {
				return true
			}
		}
		return false
	}
	return FirstTrue(len(answerKey)+1, func(length int) bool { return !feasible(length) }) - 1
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
