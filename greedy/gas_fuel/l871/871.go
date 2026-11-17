package l871

// RefuelByCount[j] 是恰好停靠 j 次可达到的最远位置；每站倒序避免重复使用。
func RefuelByCount(startFuel int, stations [][]int) []int64 {
	reach := make([]int64, len(stations)+1)
	for j := 1; j < len(reach); j++ {
		reach[j] = -1
	}
	reach[0] = int64(startFuel)
	for i, station := range stations {
		for stops := i + 1; stops >= 1; stops-- {
			if reach[stops-1] >= int64(station[0]) {
				reach[stops] = max(reach[stops], reach[stops-1]+int64(station[1]))
			}
		}
	}
	return reach
}

// https://leetcode.com/problems/minimum-number-of-refueling-stops/
// 保留原主版本的按次数 DP 方法，倒序更新后选第一个能覆盖目标的停靠数。
func minRefuelStops(target int, startFuel int, stations [][]int) int {
	for stops, reach := range RefuelByCount(startFuel, stations) {
		if reach >= int64(target) {
			return stops
		}
	}
	return -1
}
