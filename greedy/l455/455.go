package l455

import (
	"sort"
)

// MatchSorted 已排序的需求与资源，每次用最小够用资源满足最小需求。
func MatchSorted(requirements, resources []int) int {
	matched := 0
	for _, resource := range resources {
		if matched < len(requirements) && resource >= requirements[matched] {
			matched++
		}
	}
	return matched
}

// https://leetcode.com/problems/assign-cookies/
// 饼干太小则跳过；够用时不浪费更大的饼干。
func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	return MatchSorted(g, s)
}
