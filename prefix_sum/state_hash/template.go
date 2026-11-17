// Package prefixstate 提供“当前前缀寻找历史前缀”的两种聚合模板。
package prefixstate

// Count 统计满足状态关系的非空连续区间数量。
// advance 推进前缀状态；needed 返回当前状态应匹配的历史状态。
// 有效区间必须能表示为 previousState == needed(currentState) 的单键匹配。
// 计数必须先查询、后记录，避免把当前前缀与自身配成空区间。
func Count[T any, S comparable](values []T, initial S, advance func(S, T) S, needed func(S) S) int {
	freq := map[S]int{initial: 1}
	state, answer := initial, 0
	for _, value := range values {
		state = advance(state, value)
		answer += freq[needed(state)]
		freq[state]++
	}
	return answer
}

// Longest 返回满足同类状态关系的最长连续区间长度。
// earliest 保存前缀第一次出现的位置，空前缀位于 -1；不得覆盖较早位置。
// 有效区间同样必须对应一个历史键，而非任意区间关系或多个可接受键。
func Longest[T any, S comparable](values []T, initial S, advance func(S, T) S, needed func(S) S) int {
	earliest := map[S]int{initial: -1}
	state, answer := initial, 0
	for i, value := range values {
		state = advance(state, value)
		if position, ok := earliest[needed(state)]; ok && i-position > answer {
			answer = i - position
		}
		if _, exists := earliest[state]; !exists {
			earliest[state] = i
		}
	}
	return answer
}
