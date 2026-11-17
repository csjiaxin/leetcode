package l844

// PreviousVisible 是含退格字符串的反向迭代器，返回下一个可见字符下标或 -1。
func PreviousVisible(s string, i int) int {
	skipped := 0
	for i >= 0 {
		if s[i] == '#' {
			skipped++
			i--
		} else if skipped > 0 {
			skipped--
			i--
		} else {
			break
		}
	}
	return i
}

// https://leetcode.com/problems/backspace-string-compare/
// 两个反向可见字符迭代器同步比较，额外空间 O(1)。
func backspaceCompare(s string, t string) bool {
	i, j := len(s)-1, len(t)-1
	for i >= 0 || j >= 0 {
		i, j = PreviousVisible(s, i), PreviousVisible(t, j)
		if i < 0 || j < 0 {
			return i == j
		}
		if s[i] != t[j] {
			return false
		}
		i--
		j--
	}
	return true
}
