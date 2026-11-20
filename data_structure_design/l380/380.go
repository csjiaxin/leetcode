package l380

import "math/rand"

// RandomizedSet keeps a bijection between values and dense slice indices.
// Insert takes average amortized O(1), Remove average O(1), and GetRandom O(1).
// Space is O(n). Construct before use; do not copy after use.
// GetRandom requires a nonempty set. Not safe for concurrent use.
type RandomizedSet struct {
	values []int
	index  map[int]int
}

func Constructor() RandomizedSet {
	return RandomizedSet{index: make(map[int]int)}
}

func (s *RandomizedSet) Insert(val int) bool {
	if _, exists := s.index[val]; exists {
		return false
	}
	s.index[val] = len(s.values)
	s.values = append(s.values, val)
	return true
}

func (s *RandomizedSet) Remove(val int) bool {
	i, exists := s.index[val]
	if !exists {
		return false
	}
	last := len(s.values) - 1
	moved := s.values[last]
	s.values[i] = moved
	s.index[moved] = i // Also valid when removing the last element itself.
	s.values = s.values[:last]
	delete(s.index, val)
	return true
}

func (s *RandomizedSet) GetRandom() int {
	// Every remaining value occupies exactly one index, so sampling is uniform.
	return s.values[rand.Intn(len(s.values))]
}
