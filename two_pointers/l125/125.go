package l125

// Opposite 维护闭区间 [left,right]；step 必须收缩区间且不扩大边界。
// 指针不得访问区间之外；返回收缩后的边界，必要时可再处理剩余中心。
func Opposite(left, right int, step func(l, r int) (int, int)) (int, int) {
	for left < right {
		left, right = step(left, right)
	}
	return left, right
}

func Alnum(c byte) bool { return c >= '0' && c <= '9' || c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z' }

func Lower(c byte) byte {
	if c >= 'A' && c <= 'Z' {
		return c + ('a' - 'A')
	}
	return c
}

// https://leetcode.com/problems/valid-palindrome/
// 按题目 ASCII 条件，跳过非字母数字并忽略大小写。
func isPalindrome(s string) bool {
	answer := true
	Opposite(0, len(s)-1, func(l, r int) (int, int) {
		if !Alnum(s[l]) {
			return l + 1, r
		}
		if !Alnum(s[r]) {
			return l, r - 1
		}
		if Lower(s[l]) != Lower(s[r]) {
			answer = false
			return r, r
		}
		return l + 1, r - 1
	})
	return answer
}
