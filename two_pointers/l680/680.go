package l680

// Opposite 维护闭区间 [left,right]；step 必须收缩区间且不扩大边界。
// 指针不得访问区间之外；返回收缩后的边界，必要时可再处理剩余中心。
func Opposite(left, right int, step func(l, r int) (int, int)) (int, int) {
	for left < right {
		left, right = step(left, right)
	}
	return left, right
}

// PalindromeRange 检查闭区间；需要同时复制 Opposite 模板。
func PalindromeRange(s string, l, r int) bool {
	answer := true
	Opposite(l, r, func(l, r int) (int, int) {
		if s[l] != s[r] {
			answer = false
			return r, r
		}
		return l + 1, r - 1
	})
	return answer
}

// https://leetcode.com/problems/valid-palindrome-ii/
// 首次失配只能消耗一次删除机会，分别验证跳左或跳右后的完整回文。
func validPalindrome(s string) bool {
	answer := true
	Opposite(0, len(s)-1, func(l, r int) (int, int) {
		if s[l] != s[r] {
			answer = PalindromeRange(s, l+1, r) || PalindromeRange(s, l, r-1)
			return r, r
		}
		return l + 1, r - 1
	})
	return answer
}
