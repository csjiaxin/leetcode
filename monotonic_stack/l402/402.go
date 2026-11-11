package l402

import "strings"

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

// https://leetcode.com/problems/remove-k-digits/
func removeKdigits(num string, k int) string {
	stack := New[rune]()
	for _, c := range num {
		for stack.Size() > 0 && stack.Peek() > c && k > 0 {
			stack.Pop()
			k--
		}
		stack.Push(c)
	}

	for ; k > 0; k-- {
		stack.Pop()
	}
	ret := ""
	for stack.Size() > 0 {
		popped := stack.Pop()
		ret = string(popped) + ret
	}
	ret = strings.TrimLeft(ret, "0")
	if ret == "" {
		return "0"
	}
	return ret
}
