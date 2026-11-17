package l871

import (
	"container/heap"
)

type FuelHeap []int

func (h FuelHeap) Len() int           { return len(h) }
func (h FuelHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h FuelHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *FuelHeap) Push(value any)    { *h = append(*h, value.(int)) }
func (h *FuelHeap) Pop() any {
	values := *h
	last := values[len(values)-1]
	*h = values[:len(values)-1]
	return last
}

// GreedyRefuel 从已可达站点选择最大油量，延迟决定在哪个经过的站停靠。
// 需要 FuelHeap 与标准库 container/heap；stations 按位置升序。
func GreedyRefuel(target, startFuel int, stations [][]int) int {
	reach := int64(startFuel)
	index, stops := 0, 0
	candidates := &FuelHeap{}
	heap.Init(candidates)
	for reach < int64(target) {
		for index < len(stations) && int64(stations[index][0]) <= reach {
			heap.Push(candidates, stations[index][1])
			index++
		}
		if candidates.Len() == 0 {
			return -1
		}
		reach += int64(heap.Pop(candidates).(int))
		stops++
	}
	return stops
}

// https://leetcode.com/problems/minimum-number-of-refueling-stops/
// 练习变体使用最大堆贪心 O(n log n)，与主版本 DP 独立保留。
func minRefuelStops(target int, startFuel int, stations [][]int) int {
	return GreedyRefuel(target, startFuel, stations)
}
