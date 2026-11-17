package l392

// AdvanceTwo 扫描两个独立序列；step 每次至少推进一个指针，不能后退。
func AdvanceTwo(n, m int, step func(i, j int) (int, int)) (int, int) {
	i, j := 0, 0
	for i < n && j < m {
		i, j = step(i, j)
	}
	return i, j
}

// https://leetcode.com/problems/is-subsequence/
// 长串每步推进，只有匹配成功时才推进短串。
func isSubsequence(s string, t string) bool {
	matched, _ := AdvanceTwo(len(s), len(t), func(i, j int) (int, int) {
		if s[i] == t[j] {
			return i + 1, j + 1
		}
		return i, j + 1
	})
	return matched == len(s)
}
