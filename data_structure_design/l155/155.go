package l155

type entry struct{ value, minimum int }

// MinStack records the prefix minimum at every level, including duplicates.
// Push takes amortized O(1); other operations take O(1). Space is O(n).
// Pop, Top and GetMin require a nonempty stack. Not safe for concurrent use.
type MinStack struct{ entries []entry }

func Constructor() MinStack { return MinStack{} }

func (s *MinStack) Push(val int) {
	minimum := val
	if len(s.entries) > 0 && s.entries[len(s.entries)-1].minimum < minimum {
		minimum = s.entries[len(s.entries)-1].minimum
	}
	s.entries = append(s.entries, entry{val, minimum})
}

func (s *MinStack) Pop() {
	s.entries = s.entries[:len(s.entries)-1]
}

func (s *MinStack) Top() int {
	return s.entries[len(s.entries)-1].value
}

func (s *MinStack) GetMin() int {
	return s.entries[len(s.entries)-1].minimum
}
