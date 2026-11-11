package l496

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

// https://leetcode.com/problems/next-greater-element-i/
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	stack := New[int]()
	nextGreaterMap := map[int]int{}
	for _, v := range nums2 {

		for stack.Size() > 0 && stack.Peek() < v {
			pop := stack.Pop()
			nextGreaterMap[pop] = v
		}
		stack.Push(v)
	}
	result := make([]int, len(nums1))
	for i, v := range nums1 {
		result[i] = -1
		if exist, ok := nextGreaterMap[v]; ok {
			result[i] = exist
		}
	}
	return result
}
