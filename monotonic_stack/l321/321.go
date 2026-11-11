package l321

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

// https://leetcode.com/problems/create-maximum-number/
func maxNumber(nums1 []int, nums2 []int, k int) []int {
	// Like 402, but remove smaller digits to maximize the subsequence.
	pickMax := func(nums []int, size int) []int {
		stack := New[int]()
		drop := len(nums) - size
		for _, v := range nums {
			for stack.Size() > 0 && stack.Peek() < v && drop > 0 {
				stack.Pop()
				drop--
			}
			stack.Push(v)
		}
		for ; drop > 0; drop-- {
			stack.Pop()
		}
		return stack.d
	}

	// Compare whole suffixes, not just their first digits.
	greater := func(a []int, i int, b []int, j int) bool {
		for i < len(a) && j < len(b) && a[i] == b[j] {
			i++
			j++
		}
		return j == len(b) || (i < len(a) && a[i] > b[j])
	}

	best := make([]int, 0, k)
	start := 0
	if k > len(nums2) {
		start = k - len(nums2)
	}
	for size := start; size <= k && size <= len(nums1); size++ {
		a := pickMax(nums1, size)
		b := pickMax(nums2, k-size)
		candidate := make([]int, 0, k)
		for i, j := 0, 0; i < len(a) || j < len(b); {
			if greater(a, i, b, j) {
				candidate = append(candidate, a[i])
				i++
			} else {
				candidate = append(candidate, b[j])
				j++
			}
		}
		if greater(candidate, 0, best, 0) {
			best = candidate
		}
	}
	return best
}
