package l974

func subarraysDivByK(nums []int, k int) int {
	return Count(nums, 0, func(remainder, value int) int { return ((remainder+value)%k + k) % k }, func(remainder int) int { return remainder })
}

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
