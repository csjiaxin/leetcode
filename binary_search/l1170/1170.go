package l1170

import "sort"

// 1170：复用本文件的边界模板；特殊结构按对应排除不变量处理。
func frequency(s string) int {
	if len(s) == 0 {
		return 0
	}
	smallest, count := s[0], 0
	for i := range s {
		if s[i] < smallest {
			smallest, count = s[i], 1
		} else if s[i] == smallest {
			count++
		}
	}
	return count
}
func numSmallerByFrequency(queries, words []string) []int {
	values := make([]int, len(words))
	for i, word := range words {
		values[i] = frequency(word)
	}
	sort.Ints(values)
	answer := make([]int, len(queries))
	for i, query := range queries {
		value := frequency(query)
		answer[i] = len(values) - FirstTrue(len(values), func(j int) bool { return values[j] > value })
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
