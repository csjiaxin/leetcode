package l1094

// 支持有符号整数、浮点数及以它们为底层类型的自定义类型。
// 不包含无符号整数，避免负增量和取消事件发生无符号下溢。
type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~float32 | ~float64
}

// 一次闭区间更新：[Left, Right] 内的每个元素增加 Delta。
type RangeAdd[T Number] struct {
	Left, Right int
	Delta       T
}

// 返回新数组，不修改 nums 或 updates。
func applyRangeAdds[T Number](nums []T, updates []RangeAdd[T]) []T {
	n := len(nums)
	diff := make([]T, n+1) // diff[n]用于记录末尾的取消事件
	for _, u := range updates {
		diff[u.Left] += u.Delta
		diff[u.Right+1] -= u.Delta
	}

	answer := make([]T, n)
	var extra T // T 的零值，即数值0
	for i, value := range nums {
		extra += diff[i]
		answer[i] = value + extra
	}
	return answer
}

// https://leetcode.com/problems/car-pooling/
func carPooling(trips [][]int, capacity int) bool {
	maxPosition := 0
	updates := make([]RangeAdd[int], len(trips))
	for i, t := range trips {
		maxPosition = max(maxPosition, t[2])
		updates[i] = RangeAdd[int]{
			Left:  t[1],
			Right: t[2] - 1,
			Delta: t[0],
		}
	}
	input := make([]int, maxPosition+1)

	result := applyRangeAdds(input, updates)

	for _, t := range result {
		if t > capacity {
			return false
		}
	}
	return true
}
