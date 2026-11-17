package l1289

// Layers 的 transition 读取上一层，返回独立的新一层；不能覆盖 previous 的切片内容。
func Layers[S any](count int, initial S, transition func(layer int, previous S) S) S {
	state := initial
	for layer := 0; layer < count; layer++ {
		state = transition(layer, state)
	}
	return state
}

// 上一行必须不同列；维护来自不同下标的最小、次小值，重复最小值也算两个来源，O(mn)。
func minFallingPathSum(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	last := Layers(len(grid), []int(nil), func(r int, previous []int) []int {
		if r == 0 {
			return append([]int(nil), grid[0]...)
		}
		first, second, index := int(^uint(0)>>1), int(^uint(0)>>1), -1
		for c, value := range previous {
			if value < first {
				first, second, index = value, first, c
			} else if value < second {
				second = value
			}
		}
		row := make([]int, len(grid[r]))
		for c, value := range grid[r] {
			best := first
			if c == index {
				best = second
			}
			row[c] = best + value
		}
		return row
	})
	answer := last[0]
	for _, value := range last {
		answer = min(answer, value)
	}
	return answer
}
