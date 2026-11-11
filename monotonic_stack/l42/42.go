package l42

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

// https://leetcode.com/problems/trapping-rain-water/
func trap(height []int) int {
	idxStack := New[E]()
	ret := 0

	handleCurrent := func(pop E, right int) {
		left := -1
		if idxStack.Size() > 0 {
			left = idxStack.Peek().idx
		}
		currentIdx := pop.idx
		currentValue := pop.value
		leftNums := currentIdx - 1 - left
		rightNums := right - 1 - currentIdx
		area := (leftNums + 1) * (rightNums + 1) * currentValue
		ret += area
		ret = ret % (1_000_000_007)

	}
	for i, v := range heights {
		for idxStack.Size() > 0 && idxStack.Peek().value > v {
			pop := idxStack.Pop()
			handleCurrent(pop, i)
		}
		idxStack.Push(E{idx: i, value: v})
	}

	for idxStack.Size() > 0 {
		pop := idxStack.Pop()
		handleCurrent(pop, len(heights))
	}
	return ret
}
