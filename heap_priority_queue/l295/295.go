package l295

import "leetcode/heap_priority_queue/internal/pqueue"

// MedianFinder partitions values into lower(max heap) and upper(min heap).
// Lower has either the same size as upper or one additional element.
type MedianFinder struct{ lower, upper *pqueue.Queue[int] }

func Constructor() MedianFinder {
	return MedianFinder{pqueue.New(func(a, b int) bool { return a > b }), pqueue.New(func(a, b int) bool { return a < b })}
}

// AddNum takes O(log n) time. All inserted values occupy O(n) space.
func (m *MedianFinder) AddNum(num int) {
	if m.lower.Len() == 0 || num <= m.lower.Peek() {
		m.lower.Push(num)
	} else {
		m.upper.Push(num)
	}
	if m.lower.Len() > m.upper.Len()+1 {
		m.upper.Push(m.lower.Pop())
	}
	if m.upper.Len() > m.lower.Len() {
		m.lower.Push(m.upper.Pop())
	}
}

// FindMedian takes O(1) time; at least one number must have been inserted.
func (m *MedianFinder) FindMedian() float64 {
	if m.lower.Len() > m.upper.Len() {
		return float64(m.lower.Peek())
	}
	return float64(m.lower.Peek())/2 + float64(m.upper.Peek())/2
}
