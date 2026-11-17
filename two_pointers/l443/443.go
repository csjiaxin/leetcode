package l443

import (
	"strconv"
)

// ForRuns 将相等连续段作为 [start,end) 交给 visit。
// visit 可改写已处理前缀，不能修改尚未扫描的 end 之后内容。
func ForRuns[T comparable](values []T, visit func(start, end int)) {
	for start := 0; start < len(values); {
		end := start + 1
		for end < len(values) && values[end] == values[start] {
			end++
		}
		visit(start, end)
		start = end
	}
}

// https://leetcode.com/problems/string-compression/
// 整段读取完成才写入字符和次数，写指针不会超过已读边界。
func compress(chars []byte) int {
	write := 0
	ForRuns(chars, func(start, end int) {
		value := chars[start]
		chars[write] = value
		write++
		if count := end - start; count > 1 {
			for _, digit := range strconv.Itoa(count) {
				chars[write] = byte(digit)
				write++
			}
		}
	})
	return write
}
