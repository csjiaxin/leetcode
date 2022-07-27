package main

import "fmt"

type Ele struct {
	i int
	v int
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
func nextGreaterElements(nums []int) []int {
	s := New()
	n := len(nums)
	r := make([]int, n)
	for i := 0; i < n; i++ {
		r[i] = -1
	}
	for i := 0; i < 2*n; i++ {
		v := nums[i%n]
		for s.Len() > 0 && s.Peek().v < v {
			r[s.Pop().i] = v
		}
		s.Push(Ele{i % n, v})
	}
	return r
}
func main() {
	fmt.Println(nextGreaterElements([]int{1, 2, 1}))
}
