package l1109

// 支持有符号整数、浮点数，以及以它们为底层类型的自定义类型。
// 不包含无符号整数，避免负增量和取消事件发生无符号下溢。
type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~float32 | ~float64
}

type RangeAdd[T Number] struct {
	Left, Right int
	Delta       T
}

// 返回新数组，不修改 nums 或 updates。
func applyRangeAdds[T Number](nums []T, updates []RangeAdd[T]) []T {
	diff := make([]T, len(nums)+1) // 末尾多一个取消事件的位置
	for _, u := range updates {
		diff[u.Left] += u.Delta
		diff[u.Right+1] -= u.Delta
	}

	answer := make([]T, len(nums))
	var extra T
	for i, value := range nums {
		extra += diff[i]
		answer[i] = value + extra
	}
	return answer
}

// https://leetcode.com/problems/corporate-flight-bookings/
func corpFlightBookings(bookings [][]int, n int) []int {

	input := make([]int, n)
	updates := make([]RangeAdd[int], len(bookings))
	for i, booking := range bookings {
		updates[i] = RangeAdd[int]{
			Left:  booking[0] - 1,
			Right: booking[1] - 1,
			Delta: booking[2],
		}
	}
	ret := applyRangeAdds(input, updates)
	return ret
}
