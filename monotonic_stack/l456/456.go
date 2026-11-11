package l456

import "math"

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

// https://leetcode.com/problems/132-pattern/
func find132pattern(nums []int) bool {
	stack := New[int]()

	N := len(nums)
	second := math.MinInt
	for j := N - 1; j >= 0; j-- {
		v := nums[j]
		if stack.Size() > 0 && v < stack.Peek() && second != math.MinInt && v < second {
			return true
		}
		for stack.Size() > 0 && stack.Peek() < v {
			popped := stack.Pop()
			second = max(second, popped)
		}
		stack.Push(v)
	}
	return false
}
