package l1854

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

// 统一转换为0-based闭区间，再复用差分模板。
func maximumPopulation(logs [][]int) int {
	const firstYear, lastYear = 1950, 2050
	updates := make([]RangeAdd[int], 0, len(logs))
	for _, entry := range logs {
		updates = append(updates, RangeAdd[int]{entry[0] - firstYear, entry[1] - firstYear - 1, 1})
	}
	population := applyRangeAdds(make([]int, lastYear-firstYear+1), updates)
	best, year := 0, firstYear
	for i, count := range population {
		if count > best {
			best, year = count, firstYear+i
		}
	}
	return year
}
