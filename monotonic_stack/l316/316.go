package l316

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

// https://leetcode.com/problems/remove-duplicate-letters/
func removeDuplicateLetters(s string) string {
	stack := New[rune]()
	remainning := map[rune]int{}
	instack := map[rune]bool{}
	for _, c := range s {
		remainning[c] += 1
	}
	for _, c := range s {
		for stack.Size() > 0 && !instack[c] && stack.Peek() > c && remainning[stack.Peek()] > 1 {
			popped := stack.Pop()
			instack[popped] = false
			remainning[popped] -= 1
		}
		if !instack[c] {
			stack.Push(c)
			instack[c] = true
		} else {
			remainning[c] -= 1
		}
	}

	ret := ""
	for stack.Size() > 0 {
		popped := stack.Pop()
		ret = string(popped) + ret
	}
	return ret
}
