package l202

// Floyd 用确定后继函数寻找环内相遇点；next 的 bool=false 表示轨迹结束。
func Floyd[S comparable](start S, next func(S) (S, bool)) (S, bool) {
	slow, fast := start, start
	for {
		var ok bool
		fast, ok = next(fast)
		if !ok {
			var zero S
			return zero, false
		}
		fast, ok = next(fast)
		if !ok {
			var zero S
			return zero, false
		}
		slow, ok = next(slow)
		if !ok {
			var zero S
			return zero, false
		}
		if slow == fast {
			return slow, true
		}
	}
}

// https://leetcode.com/problems/happy-number/
// 数字的确定后继为数位平方和；轨迹在 1 的自环中相遇才是快乐数。
func isHappy(n int) bool {
	next := func(value int) (int, bool) {
		sum := 0
		for value > 0 {
			digit := value % 10
			sum += digit * digit
			value /= 10
		}
		return sum, true
	}
	meeting, _ := Floyd(n, next)
	return meeting == 1
}
