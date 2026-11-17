package l1397

func findGoodStrings(n int, s1, s2, evil string) int {
	const mod = 1000000007
	m := len(evil)
	if m == 0 || s1 > s2 {
		return 0
	}
	pi := PrefixFunction([]byte(evil))
	transition := make([][26]int, m)
	for state := 0; state < m; state++ {
		for c := 0; c < 26; c++ {
			value := byte('a' + c)
			if value == evil[state] {
				transition[state][c] = state + 1
			} else if state > 0 {
				transition[state][c] = transition[pi[state-1]][c]
			}
		}
	}
	dp := make([][4]int, m)
	dp[0][3] = 1
	for position := 0; position < n; position++ {
		next := make([][4]int, m)
		for state, row := range dp {
			for flags, count := range row {
				if count == 0 {
					continue
				}
				lo, hi := byte('a'), byte('z')
				if flags&1 != 0 {
					lo = s1[position]
				}
				if flags&2 != 0 {
					hi = s2[position]
				}
				for value := lo; value <= hi; value++ {
					newState := transition[state][value-'a']
					if newState == m {
						continue
					}
					newFlags := 0
					if flags&1 != 0 && value == s1[position] {
						newFlags |= 1
					}
					if flags&2 != 0 && value == s2[position] {
						newFlags |= 2
					}
					next[newState][newFlags] = (next[newState][newFlags] + count) % mod
				}
			}
		}
		dp = next
	}
	answer := 0
	for _, row := range dp {
		for _, count := range row {
			answer = (answer + count) % mod
		}
	}
	return answer
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
