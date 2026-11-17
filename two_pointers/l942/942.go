package l942

// ChooseEnds 从尚未使用的连续整数 [low,high] 中取当前需要的极值。
func ChooseEnds(n int, chooseLow func(i int) bool) []int {
	answer := make([]int, n+1)
	low, high := 0, n
	for i := 0; i < n; i++ {
		if chooseLow(i) {
			answer[i] = low
			low++
		} else {
			answer[i] = high
			high--
		}
	}
	answer[n] = low
	return answer
}

// https://leetcode.com/problems/di-string-match/
// I 取剩余最小值，D 取剩余最大值，后续剩余值都能满足当前不等式。
func diStringMatch(s string) []int {
	return ChooseEnds(len(s), func(i int) bool { return s[i] == 'I' })
}
