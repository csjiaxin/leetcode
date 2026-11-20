package l232

// MyQueue keeps newly pushed elements in input and the oldest at output's top.
// Each element transfers at most once: operations take amortized O(1) time,
// while one Pop or Peek can take O(n). Space is O(n).
// Pop and Peek require a nonempty queue. Not safe for concurrent use.
type MyQueue struct{ input, output []int }

func Constructor() MyQueue { return MyQueue{} }

func (q *MyQueue) Push(x int) {
	q.input = append(q.input, x)
}

func (q *MyQueue) refill() {
	if len(q.output) != 0 {
		return
	}
	for len(q.input) > 0 {
		last := len(q.input) - 1
		q.output = append(q.output, q.input[last])
		q.input = q.input[:last]
	}
}

func (q *MyQueue) Pop() int {
	q.refill()
	last := len(q.output) - 1
	value := q.output[last]
	q.output = q.output[:last]
	return value
}

func (q *MyQueue) Peek() int {
	q.refill()
	return q.output[len(q.output)-1]
}

func (q *MyQueue) Empty() bool {
	return len(q.input) == 0 && len(q.output) == 0
}
