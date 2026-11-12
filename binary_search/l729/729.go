package l729

type interval struct{ start, end int }
type MyCalendar struct{ bookings []interval }

func Constructor() MyCalendar {
}

func (c *MyCalendar) Book(start, end int) bool {
}
