package l3034

func countMatchingSubarrays(nums []int, pattern []int) int {
	if len(nums) == 0 {
		return 0
	}
	comparisons := make([]int, len(nums)-1)
	for i := range comparisons {
		if nums[i+1] > nums[i] {
			comparisons[i] = 1
		} else if nums[i+1] < nums[i] {
			comparisons[i] = -1
		}
	}
	return len(FindAll(comparisons, pattern))
}

// PrefixFunction 记录每个前缀的最长相等真前后缀长度。
func PrefixFunction[T comparable](pattern []T) []int {
	pi := make([]int, len(pattern))
	for i := 1; i < len(pattern); i++ {
		j := pi[i-1]
		for j > 0 && pattern[i] != pattern[j] {
			j = pi[j-1]
		}
		if pattern[i] == pattern[j] {
			j++
		}
		pi[i] = j
	}
	return pi
}

// FindAll 在文本中寻找模式，匹配成功后回退以允许重叠。
func FindAll[T comparable](text, pattern []T) []int {
	answer := []int{}
	if len(pattern) == 0 {
		for i := 0; i <= len(text); i++ {
			answer = append(answer, i)
		}
		return answer
	}
	pi := PrefixFunction(pattern)
	matched := 0
	for i, value := range text {
		for matched > 0 && value != pattern[matched] {
			matched = pi[matched-1]
		}
		if value == pattern[matched] {
			matched++
		}
		if matched == len(pattern) {
			answer = append(answer, i-len(pattern)+1)
			matched = pi[matched-1]
		}
	}
	return answer
}
