package l1710

import (
	"sort"
)

// TakeByValue：所有物品占用相同容量单位，类型为 [数量,单位收益]。
func TakeByValue(items [][]int, capacity int) int {
	sort.Slice(items, func(i, j int) bool { return items[i][1] > items[j][1] })
	total := 0
	for _, item := range items {
		take := min(capacity, item[0])
		total += take * item[1]
		capacity -= take
		if capacity == 0 {
			break
		}
	}
	return total
}

// https://leetcode.com/problems/maximum-units-on-a-truck/
// 每箱占一个车位，按单位收益优先取，允许只取某种箱子的一部分。
func maximumUnits(boxTypes [][]int, truckSize int) int { return TakeByValue(boxTypes, truckSize) }
