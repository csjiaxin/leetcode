package l120

// Layers 的 transition 读取上一层，返回独立的新一层；不能覆盖 previous 的切片内容。
func Layers[S any](count int, initial S, transition func(layer int, previous S) S) S {
	state := initial
	for layer := 0; layer < count; layer++ {
		state = transition(layer, state)
	}
	return state
}

// 每行从上一行相邻两个合法列转移；最终答案取末行最小值。
func minimumTotal(triangle [][]int) int {
	if len(triangle) == 0 {
		return 0
	}
	last := Layers(len(triangle), []int(nil), func(r int, previous []int) []int {
		row := make([]int, len(triangle[r]))
		for c, value := range triangle[r] {
			if r == 0 {
				row[c] = value
				continue
			}
			best := int(^uint(0) >> 1)
			if c < len(previous) {
				best = min(best, previous[c])
			}
			if c > 0 {
				best = min(best, previous[c-1])
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
