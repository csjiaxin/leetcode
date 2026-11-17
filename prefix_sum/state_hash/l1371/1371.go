package l1371

func findTheLongestSubstring(s string) int {
	return Longest([]byte(s), 0, func(mask int, value byte) int {
		for i, vowel := range []byte("aeiou") {
			if value == vowel {
				return mask ^ (1 << i)
			}
		}
		return mask
	}, func(mask int) int { return mask })
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
