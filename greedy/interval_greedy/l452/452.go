package l452

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

// https://leetcode.com/problems/minimum-number-of-arrows-to-burst-balloons/
// 闭区间相接时可共用一箭；严格分离才需要新箭。
func findMinArrowShots(points [][]int) int {
	return SelectByEnd(points, func(start, end int) bool { return start > end })
}
