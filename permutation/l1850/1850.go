package l1850

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

func getMinSwaps(num string, k int) int {
	target := []byte(num)
	for step := 0; step < k; step++ {
		nextLexicographic(target)
	}
	current := []byte(num)
	swaps := 0
	for i, digit := range target {
		j := i
		for current[j] != digit {
			j++
		}
		for j > i {
			current[j], current[j-1] = current[j-1], current[j]
			j--
			swaps++
		}
	}
	return swaps
}
