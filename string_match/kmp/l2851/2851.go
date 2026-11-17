package l2851

func numberOfWays(s, t string, k int64) int {
	const mod int64 = 1000000007
	if len(s) != len(t) {
		return 0
	}
	n := len(s)
	if n < 2 {
		if k == 0 && s == t {
			return 1
		}
		return 0
	}
	good := int64(0)
	for _, index := range FindAll([]byte(s+s), []byte(t)) {
		if index < n {
			good++
		}
	}
	type matrix [2][2]int64
	multiply := func(a, b matrix) matrix {
		var c matrix
		for i := 0; i < 2; i++ {
			for j := 0; j < 2; j++ {
				for p := 0; p < 2; p++ {
					c[i][j] = (c[i][j] + a[i][p]*b[p][j]) % mod
				}
			}
		}
		return c
	}
	base := matrix{{(good - 1 + mod) % mod, good}, {int64(n) - good, (int64(n) - good - 1 + mod) % mod}}
	result := matrix{{1, 0}, {0, 1}}
	for k > 0 {
		if k&1 != 0 {
			result = multiply(result, base)
		}
		base = multiply(base, base)
		k >>= 1
	}
	if s == t {
		return int(result[0][0])
	}
	return int(result[0][1])
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
