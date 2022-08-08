package main

import "fmt"

type E struct {
	v int
	i int
}
type stack []E

func New() *stack {
	s := make(stack, 0)
	return &s
}
func (s *stack) Len() int {
	return len(*s)
}

func (s *stack) Push(v E) {
	*s = append(*s, v)
}

func (s *stack) Pop() E {
	l := len(*s)
	v := (*s)[l-1]
	*s = (*s)[:l-1]
	return v
}
func (s *stack) Peek() E {
	l := len(*s)
	v := (*s)[l-1]
	return v
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func findLengthOfShortestSubarray(arr []int) int {
	n := len(arr)
	s := New()

	L := 0
	for _, v := range arr {
		for s.Len() > 0 && s.Peek().v > v {
			s.Pop()
		}
		l := 1
		if s.Len() > 0 {
			l += s.Peek().i
		}
		L = max(L, l)
		s.Push(E{v, l})
	}
	return n - L

}
func main() {
	fmt.Println(findLengthOfShortestSubarray([]int{1, 2, 3, 10, 4, 2, 3, 5}))
}
