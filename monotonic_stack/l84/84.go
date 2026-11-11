package l84

type Stack[T any] struct {
	d []T
}

func New[T any]() *Stack[T] {
	s := &Stack[T]{
		d: []T{},
	}
	return s
}
func (s *Stack[T]) Size() int {
	return len(s.d)
}

func (s *Stack[T]) Peek() T {
	return s.d[len(s.d)-1]
}

func (s *Stack[T]) Pop() T {
	v := s.d[len(s.d)-1]
	s.d = s.d[:len(s.d)-1]
	return v
}
func (s *Stack[T]) Push(v T) {
	s.d = append(s.d, v)
}

type E struct {
	idx   int
	value int
}

// https://leetcode.com/problems/largest-rectangle-in-histogram/
func largestRectangleArea(heights []int) int {
	idxStack := New[E]()
	ret := 0
	for i, v := range heights {
		for idxStack.Size() > 0 && idxStack.Peek().value > v {
			pop := idxStack.Pop()
			left := -1
			if idxStack.Size() > 0 {
				left = idxStack.Peek().idx
			}
			right := i
			width := right - left - 1 // (left,right)
			area := width * pop.value
			ret = max(ret, area)
		}
		idxStack.Push(E{idx: i, value: v})
	}

	for idxStack.Size() > 0 {
		pop := idxStack.Pop()
		left := -1
		if idxStack.Size() > 0 {
			left = idxStack.Peek().idx
		}
		right := len(heights)
		width := right - left - 1 // (left,right)
		area := width * pop.value
		ret = max(ret, area)
	}
	return ret
}
