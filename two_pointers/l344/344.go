package l344

// Opposite 维护闭区间 [left,right]；step 必须收缩区间且不扩大边界。
// 指针不得访问区间之外；返回收缩后的边界，必要时可再处理剩余中心。
func Opposite(left, right int, step func(l, r int) (int, int)) (int, int) {
	for left < right {
		left, right = step(left, right)
	}
	return left, right
}

// https://leetcode.com/problems/reverse-string/
// 交换两端，再同时收缩。
func reverseString(s []byte) {
	Opposite(0, len(s)-1, func(l, r int) (int, int) { s[l], s[r] = s[r], s[l]; return l + 1, r - 1 })
}
