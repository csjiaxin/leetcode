package l60

// Standalone lexicographic/factorial subtemplate and problem adapter.
// factorialSelection chooses a zero-based block among the remaining digits.
func factorialSelection(n, k int) string {
	factorial := make([]int, n+1)
	factorial[0] = 1
	available := make([]int, n)
	for i := 1; i <= n; i++ {
		factorial[i] = factorial[i-1] * i
		available[i-1] = i
	}
	result := make([]byte, 0, n)
	k--
	for remaining := n; remaining > 0; remaining-- {
		block := factorial[remaining-1]
		index := k / block
		k %= block
		result = append(result, byte('0'+available[index]))
		available = append(available[:index], available[index+1:]...)
	}
	return string(result)
}

func getPermutation(n int, k int) string { return factorialSelection(n, k) }
