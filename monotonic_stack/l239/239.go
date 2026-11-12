package l239

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

// https://leetcode.com/problems/sliding-window-maximum/
func maxSlidingWindow(nums []int, k int) []int {
	if k <= 0 || k > len(nums) {
		return nil
	}

	queue := make([]int, 0, k) // Indices; corresponding values strictly decrease.
	ret := make([]int, 0, len(nums)-k+1)
	for right, value := range nums {
		// Remove indices outside the current window [right-k+1, right].
		if len(queue) > 0 && queue[0] <= right-k {
			queue = queue[1:]
		}

		// The new index has an equal or greater value and expires later.
		for len(queue) > 0 && nums[queue[len(queue)-1]] <= value {
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, right)

		if right >= k-1 {
			ret = append(ret, nums[queue[0]])
		}
	}
	return ret
}
