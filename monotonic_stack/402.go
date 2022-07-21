package main

type Ele struct {
	i int
	v byte
}
type (
	Stack struct {
		top    *node
		length int
	}
	node struct {
		value Ele
		prev  *node
	}
)

func New() *Stack {
	return &Stack{nil, 0}
}
func (this *Stack) Len() int {
	return this.length
}
func (this *Stack) Peek() Ele {
	return this.top.value
}
func (this *Stack) Pop() Ele {
	n := this.top
	this.top = n.prev
	this.length--
	return n.value
}
func (this *Stack) Push(value Ele) {
	n := &node{value, this.top}
	this.top = n
	this.length++
}
func initArr(arr []int, n int, v int) {
	for i := 0; i < n; i++ {
		arr[i] = v
	}
}
func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}

func removeKdigits(num string, k int) string {
	n := len(num)
	s := New()
	for i := 0; i < n; i++ {
		// v := num[i] - '0'
		v := num[i]
		for s.Len() > 0 && s.Peek().v > v && k > 0 {
			s.Pop()
			k -= 1
		}
		if v != '0' || s.Len() > 0 {
			s.Push(Ele{i, v})
		}
	}
	for ; k > 0 && s.Len() > 0; k-- {
		s.Pop()
	}
	r := ""
	for s.Len() > 0 {
		v := string(s.Pop().v)
		r = v + r
	}
	if r == "" {
		r = "0"
	}
	return r
}
func main() {
	println(removeKdigits("10", 2))
}
