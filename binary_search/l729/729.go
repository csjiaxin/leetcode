package l729

// 729：复用本文件的边界模板；特殊结构按对应排除不变量处理。
type interval struct{ start, end int }
type MyCalendar struct{ bookings []interval }

func Constructor() MyCalendar { return MyCalendar{} }
func (c *MyCalendar) Book(start, end int) bool {
	if start >= end {
		return false
	}
	position := FirstTrue(len(c.bookings), func(i int) bool { return c.bookings[i].start >= start })
	if position > 0 && c.bookings[position-1].end > start {
		return false
	}
	if position < len(c.bookings) && c.bookings[position].start < end {
		return false
	}
	c.bookings = append(c.bookings, interval{})
	copy(c.bookings[position+1:], c.bookings[position:])
	c.bookings[position] = interval{start, end}
	return true
}

func FirstTrue(n int, predicate func(int) bool) int {
	left, right := 0, n-1
	for left <= right {
		mid := left + (right-left)/2
		if predicate(mid) {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}
