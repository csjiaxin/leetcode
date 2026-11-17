package l287

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

// CycleEntry 需要同时复制 Floyd；相遇后，一端回到起点，双方每次走一步。
func CycleEntry[S comparable](start S, next func(S) (S, bool)) (S, bool) {
	meeting, ok := Floyd(start, next)
	if !ok {
		var zero S
		return zero, false
	}
	for start != meeting {
		start, _ = next(start)
		meeting, _ = next(meeting)
	}
	return start, true
}

// https://leetcode.com/problems/find-the-duplicate-number/
// 题设 n+1 个数均在 1..n；把下标到数值看作函数图，入口即重复值。
func findDuplicate(nums []int) int {
	entry, _ := CycleEntry(0, func(i int) (int, bool) { return nums[i], true })
	return entry
}
