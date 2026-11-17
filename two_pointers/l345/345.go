package l345

// Opposite 维护闭区间 [left,right]；step 必须收缩区间且不扩大边界。
// 指针不得访问区间之外；返回收缩后的边界，必要时可再处理剩余中心。
func Opposite(left, right int, step func(l, r int) (int, int)) (int, int) {
	for left < right {
		left, right = step(left, right)
	}
	return left, right
}

func Vowel(c byte) bool {
	switch c {
	case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
		return true
	}
	return false
}

// https://leetcode.com/problems/reverse-vowels-of-a-string/
// 两端分别跳过非元音，只交换元音。
func reverseVowels(s string) string {
	data := []byte(s)
	Opposite(0, len(data)-1, func(l, r int) (int, int) {
		if !Vowel(data[l]) {
			return l + 1, r
		}
		if !Vowel(data[r]) {
			return l, r - 1
		}
		data[l], data[r] = data[r], data[l]
		return l + 1, r - 1
	})
	return string(data)
}
