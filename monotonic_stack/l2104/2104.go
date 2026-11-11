package l2104

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

// https://leetcode.com/problems/sum-of-subarray-ranges/
func subArrayRanges(nums []int) int64 {
	sumExtremes := func(maximum bool) int64 {
		idxStack := New[E]()
		var ret int64

		handleCurrent := func(pop E, right int) {
			left := -1
			if idxStack.Size() > 0 {
				left = idxStack.Peek().idx
			}
			currentIdx := pop.idx
			currentValue := pop.value
			leftNums := currentIdx - 1 - left
			rightNums := right - 1 - currentIdx
			area := int64(leftNums+1) * int64(rightNums+1) * int64(currentValue)
			ret += area
		}
		for i, v := range nums {
			for idxStack.Size() > 0 && ((maximum && idxStack.Peek().value < v) ||
				(!maximum && idxStack.Peek().value > v)) {
				pop := idxStack.Pop()
				handleCurrent(pop, i)
			}
			idxStack.Push(E{idx: i, value: v})
		}

		for idxStack.Size() > 0 {
			pop := idxStack.Pop()
			handleCurrent(pop, len(nums))
		}
		return ret
	}

	return sumExtremes(true) - sumExtremes(false)
}
