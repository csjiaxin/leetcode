package l1937

// Layers 的 transition 读取上一层，返回独立的新一层；不能覆盖 previous 的切片内容。
func Layers[S any](count int, initial S, transition func(layer int, previous S) S) S {
	state := initial
	for layer := 0; layer < count; layer++ {
		state = transition(layer, state)
	}
	return state
}

// 把 max(previous[k]-abs(k-c)) 拆为左右扫描，保留 O(mn) 优化与 int64 收益。
func maxPoints(points [][]int) int64 {
	if len(points) == 0 || len(points[0]) == 0 {
		return 0
	}
	last := Layers(len(points), []int64(nil), func(r int, previous []int64) []int64 {
		n := len(points[r])
		row := make([]int64, n)
		if r == 0 {
			for c, v := range points[r] {
				row[c] = int64(v)
			}
			return row
		}
		left, right := make([]int64, n), make([]int64, n)
		left[0], right[n-1] = previous[0], previous[n-1]
		for c := 1; c < n; c++ {
			left[c] = max(previous[c], left[c-1]-1)
		}
		for c := n - 2; c >= 0; c-- {
			right[c] = max(previous[c], right[c+1]-1)
		}
		for c, v := range points[r] {
			row[c] = int64(v) + max(left[c], right[c])
		}
		return row
	})
	answer := last[0]
	for _, v := range last {
		answer = max(answer, v)
	}
	return answer
}
