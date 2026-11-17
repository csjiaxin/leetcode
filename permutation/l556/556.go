package l556

import "strconv"

// Standalone lexicographic/factorial subtemplate and problem adapter.
// nextLexicographic advances one lexicographic step, or resets to the minimum.
func nextLexicographic[T ~int | ~byte](values []T) bool {
	pivot := len(values) - 2
	for pivot >= 0 && values[pivot] >= values[pivot+1] {
		pivot--
	}
	advanced := pivot >= 0
	if advanced {
		successor := len(values) - 1
		for values[successor] <= values[pivot] {
			successor--
		}
		values[pivot], values[successor] = values[successor], values[pivot]
	}
	for left, right := pivot+1, len(values)-1; left < right; left, right = left+1, right-1 {
		values[left], values[right] = values[right], values[left]
	}
	return advanced
}

func nextGreaterElement(n int) int {
	digits := []byte(strconv.Itoa(n))
	if !nextLexicographic(digits) {
		return -1
	}
	var value int64
	for _, digit := range digits {
		value = value*10 + int64(digit-'0')
	}
	if value > 1<<31-1 {
		return -1
	}
	return int(value)
}
