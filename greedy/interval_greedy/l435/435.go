package l435

import (
	"sort"
)

// SelectByEnd 按结束点排序，separate 决定何时需要选择新的结束点。
func SelectByEnd(intervals [][]int, separate func(start, end int) bool) int {
	sort.Slice(intervals, func(i, j int) bool { return intervals[i][1] < intervals[j][1] })
	chosen, end := 0, 0
	for _, interval := range intervals {
		if chosen == 0 || separate(interval[0], end) {
			chosen++
			end = interval[1]
		}
	}
	return chosen
}

// https://leetcode.com/problems/non-overlapping-intervals/
// 相接允许共存；最少删除数等于总数减最多保留数。
func eraseOverlapIntervals(intervals [][]int) int {
	keep := SelectByEnd(intervals, func(start, end int) bool { return start >= end })
	return len(intervals) - keep
}
