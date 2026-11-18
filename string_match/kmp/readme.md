# KMP：前缀函数、全部匹配与状态扩展

[返回模板总目录](../../algorithm_templates/README.md)

PrefixFunction 维护当前前缀的最长相等真前后缀；失配沿 pi 回退。FindAll 扫描文本时不回退文本下标，匹配完整模式后回退状态以保留重叠匹配。整数相邻大小关系也可以先映射为 -1/0/1，再用相同泛型模板匹配。

周期题沿 border 链筛选合法移位；[1910](https://leetcode.com/problems/remove-all-occurrences-of-a-substring/) 保存匹配状态栈用于删除恢复；[1397](https://leetcode.com/problems/find-all-good-strings/) 把 KMP 前缀长度作为受上下界约束的 DP 状态；[2851](https://leetcode.com/problems/string-transformation/) 先数匹配旋转位置，再用二状态矩阵快速幂计算步数。模板统一匹配信息，额外约束分别建模。

## 复制约定

每题 `.go` 文件已内嵌所需模板，无需导入仓库的算法包。文档将固定模板只列一次，逐题展示适配器；每个适配器与它列出的模板组成一个独立代码组。不同题的同名 Constructor、类型或入口请分别使用。将列出的标准库 import 放在文件顶部，再放模板与适配器；复制提交时优先直接使用对应题目的完整 `.go`。

本组使用 Go 1.21+ 的 min/max 和 Go 1.18+ 的泛型。平台已提供的 ListNode 等类型沿用平台定义。

## 固定模板

### KMP

<!-- template: KMP -->
```go
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
```

## 题目索引

| 题目 | 模板 | 本地代码与测试 |
| --- | --- | --- |
| [28](https://leetcode.com/problems/find-the-index-of-the-first-occurrence-in-a-string/) | KMP | [代码](l28/28.go) · [测试](l28/28_test.go) |
| [214](https://leetcode.com/problems/shortest-palindrome/) | KMP | [代码](l214/214.go) · [测试](l214/214_test.go) |
| [459](https://leetcode.com/problems/repeated-substring-pattern/) | KMP | [代码](l459/459.go) · [测试](l459/459_test.go) |
| [686](https://leetcode.com/problems/repeated-string-match/) | KMP | [代码](l686/686.go) · [测试](l686/686_test.go) |
| [796](https://leetcode.com/problems/rotate-string/) | KMP | [代码](l796/796.go) · [测试](l796/796_test.go) |
| [1392](https://leetcode.com/problems/longest-happy-prefix/) | KMP | [代码](l1392/1392.go) · [测试](l1392/1392_test.go) |
| [1397](https://leetcode.com/problems/find-all-good-strings/) | KMP | [代码](l1397/1397.go) · [测试](l1397/1397_test.go) |
| [1408](https://leetcode.com/problems/string-matching-in-an-array/) | KMP | [代码](l1408/1408.go) · [测试](l1408/1408_test.go) |
| [1668](https://leetcode.com/problems/maximum-repeating-substring/) | KMP | [代码](l1668/1668.go) · [测试](l1668/1668_test.go) |
| [1764](https://leetcode.com/problems/form-array-by-concatenating-subarrays-of-another-array/) | KMP | [代码](l1764/1764.go) · [测试](l1764/1764_test.go) |
| [1910](https://leetcode.com/problems/remove-all-occurrences-of-a-substring/) | KMP | [代码](l1910/1910.go) · [测试](l1910/1910_test.go) |
| [2851](https://leetcode.com/problems/string-transformation/) | KMP | [代码](l2851/2851.go) · [测试](l2851/2851_test.go) |
| [3006](https://leetcode.com/problems/find-beautiful-indices-in-the-given-array-i/) | KMP | [代码](l3006/3006.go) · [测试](l3006/3006_test.go) |
| [3008](https://leetcode.com/problems/find-beautiful-indices-in-the-given-array-ii/) | KMP | [代码](l3008/3008.go) · [测试](l3008/3008_test.go) |
| [3029](https://leetcode.com/problems/minimum-time-to-revert-word-to-initial-state-i/) | KMP | [代码](l3029/3029.go) · [测试](l3029/3029_test.go) |
| [3031](https://leetcode.com/problems/minimum-time-to-revert-word-to-initial-state-ii/) | KMP | [代码](l3031/3031.go) · [测试](l3031/3031_test.go) |
| [3034](https://leetcode.com/problems/number-of-subarrays-that-match-a-pattern-i/) | KMP | [代码](l3034/3034.go) · [测试](l3034/3034_test.go) |
| [3036](https://leetcode.com/problems/number-of-subarrays-that-match-a-pattern-ii/) | KMP | [代码](l3036/3036.go) · [测试](l3036/3036_test.go) |

## 逐题适配器

### [28](https://leetcode.com/problems/find-the-index-of-the-first-occurrence-in-a-string/)

需要模板：`KMP`。

<!-- solution: string_match/kmp/l28/28.go; templates: KMP; imports: none -->
```go
func strStr(haystack string, pattern string) int {
	positions := FindAll([]byte(haystack), []byte(pattern))
	if len(positions) == 0 {
		return -1
	}
	return positions[0]
}
```

### [214](https://leetcode.com/problems/shortest-palindrome/)

需要模板：`KMP`。

<!-- solution: string_match/kmp/l214/214.go; templates: KMP; imports: none -->
```go
func shortestPalindrome(s string) string {
	values := make([]int, 0, 2*len(s)+1)
	for i := range s {
		values = append(values, int(s[i]))
	}
	values = append(values, -1)
	for i := len(s) - 1; i >= 0; i-- {
		values = append(values, int(s[i]))
	}
	pi := PrefixFunction(values)
	length := pi[len(pi)-1]
	prefix := []byte(s[length:])
	for l, r := 0, len(prefix)-1; l < r; l, r = l+1, r-1 {
		prefix[l], prefix[r] = prefix[r], prefix[l]
	}
	return string(prefix) + s
}
```

### [459](https://leetcode.com/problems/repeated-substring-pattern/)

需要模板：`KMP`。

<!-- solution: string_match/kmp/l459/459.go; templates: KMP; imports: none -->
```go
func repeatedSubstringPattern(s string) bool {
	if len(s) == 0 {
		return false
	}
	pi := PrefixFunction([]byte(s))
	border := pi[len(pi)-1]
	return border > 0 && len(s)%(len(s)-border) == 0
}
```

### [686](https://leetcode.com/problems/repeated-string-match/)

需要模板：`KMP`。

标准库：`strings`。

<!-- solution: string_match/kmp/l686/686.go; templates: KMP; imports: strings -->
```go
func repeatedStringMatch(a string, b string) int {
	if len(b) == 0 {
		return 0
	}
	if len(a) == 0 {
		return -1
	}
	count := (len(b) + len(a) - 1) / len(a)
	text := strings.Repeat(a, count)
	if len(FindAll([]byte(text), []byte(b))) > 0 {
		return count
	}
	if len(FindAll([]byte(text+a), []byte(b))) > 0 {
		return count + 1
	}
	return -1
}
```

### [796](https://leetcode.com/problems/rotate-string/)

需要模板：`KMP`。

<!-- solution: string_match/kmp/l796/796.go; templates: KMP; imports: none -->
```go
func rotateString(s string, goal string) bool {
	return len(s) == len(goal) && len(FindAll([]byte(s+s), []byte(goal))) > 0
}
```

### [1392](https://leetcode.com/problems/longest-happy-prefix/)

需要模板：`KMP`。

<!-- solution: string_match/kmp/l1392/1392.go; templates: KMP; imports: none -->
```go
func longestPrefix(s string) string {
	if len(s) == 0 {
		return ""
	}
	pi := PrefixFunction([]byte(s))
	return s[:pi[len(pi)-1]]
}
```

### [1397](https://leetcode.com/problems/find-all-good-strings/)

需要模板：`KMP`。

<!-- solution: string_match/kmp/l1397/1397.go; templates: KMP; imports: none -->
```go
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
```

### [1408](https://leetcode.com/problems/string-matching-in-an-array/)

需要模板：`KMP`。

<!-- solution: string_match/kmp/l1408/1408.go; templates: KMP; imports: none -->
```go
func stringMatching(words []string) []string {
	answer := []string{}
	for i, pattern := range words {
		for j, text := range words {
			if i != j && len(FindAll([]byte(text), []byte(pattern))) > 0 {
				answer = append(answer, pattern)
				break
			}
		}
	}
	return answer
}
```

### [1668](https://leetcode.com/problems/maximum-repeating-substring/)

需要模板：`KMP`。

<!-- solution: string_match/kmp/l1668/1668.go; templates: KMP; imports: none -->
```go
func maxRepeating(sequence string, word string) int {
	if len(word) == 0 {
		return 0
	}
	runs := make(map[int]int)
	answer := 0
	for _, start := range FindAll([]byte(sequence), []byte(word)) {
		end := start + len(word)
		runs[end] = runs[start] + 1
		answer = max(answer, runs[end])
	}
	return answer
}
```

### [1764](https://leetcode.com/problems/form-array-by-concatenating-subarrays-of-another-array/)

需要模板：`KMP`。

<!-- solution: string_match/kmp/l1764/1764.go; templates: KMP; imports: none -->
```go
func canChoose(groups [][]int, nums []int) bool {
	start := 0
	for _, pattern := range groups {
		positions := FindAll(nums[start:], pattern)
		if len(positions) == 0 {
			return false
		}
		start += positions[0] + len(pattern)
	}
	return true
}
```

### [1910](https://leetcode.com/problems/remove-all-occurrences-of-a-substring/)

需要模板：`KMP`。

<!-- solution: string_match/kmp/l1910/1910.go; templates: KMP; imports: none -->
```go
func removeOccurrences(s string, part string) string {
	if len(part) == 0 {
		return s
	}
	pi := PrefixFunction([]byte(part))
	stack := make([]byte, 0, len(s))
	states := make([]int, 0, len(s))
	matched := 0
	for i := range s {
		value := s[i]
		for matched > 0 && value != part[matched] {
			matched = pi[matched-1]
		}
		if value == part[matched] {
			matched++
		}
		stack = append(stack, value)
		states = append(states, matched)
		if matched == len(part) {
			stack = stack[:len(stack)-len(part)]
			states = states[:len(states)-len(part)]
			matched = 0
			if len(states) > 0 {
				matched = states[len(states)-1]
			}
		}
	}
	return string(stack)
}
```

### [2851](https://leetcode.com/problems/string-transformation/)

需要模板：`KMP`。

<!-- solution: string_match/kmp/l2851/2851.go; templates: KMP; imports: none -->
```go
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
```

### [3006](https://leetcode.com/problems/find-beautiful-indices-in-the-given-array-i/)

需要模板：`KMP`。

<!-- solution: string_match/kmp/l3006/3006.go; templates: KMP; imports: none -->
```go
func beautifulIndices(s, a, b string, k int) []int {
	first, second := FindAll([]byte(s), []byte(a)), FindAll([]byte(s), []byte(b))
	answer := []int{}
	j := 0
	for _, i := range first {
		for j < len(second) && second[j] < i-k {
			j++
		}
		if j < len(second) && second[j] <= i+k {
			answer = append(answer, i)
		}
	}
	return answer
}
```

### [3008](https://leetcode.com/problems/find-beautiful-indices-in-the-given-array-ii/)

需要模板：`KMP`。

<!-- solution: string_match/kmp/l3008/3008.go; templates: KMP; imports: none -->
```go
func beautifulIndices(s, a, b string, k int) []int {
	first, second := FindAll([]byte(s), []byte(a)), FindAll([]byte(s), []byte(b))
	answer := []int{}
	j := 0
	for _, i := range first {
		for j < len(second) && second[j] < i-k {
			j++
		}
		if j < len(second) && second[j] <= i+k {
			answer = append(answer, i)
		}
	}
	return answer
}
```

### [3029](https://leetcode.com/problems/minimum-time-to-revert-word-to-initial-state-i/)

需要模板：`KMP`。

<!-- solution: string_match/kmp/l3029/3029.go; templates: KMP; imports: none -->
```go
func minimumTimeToInitialState(word string, k int) int {
	if len(word) == 0 {
		return 0
	}
	pi := PrefixFunction([]byte(word))
	for border := pi[len(pi)-1]; border > 0; border = pi[border-1] {
		removed := len(word) - border
		if removed%k == 0 {
			return removed / k
		}
	}
	return (len(word) + k - 1) / k
}
```

### [3031](https://leetcode.com/problems/minimum-time-to-revert-word-to-initial-state-ii/)

需要模板：`KMP`。

<!-- solution: string_match/kmp/l3031/3031.go; templates: KMP; imports: none -->
```go
func minimumTimeToInitialState(word string, k int) int {
	if len(word) == 0 {
		return 0
	}
	pi := PrefixFunction([]byte(word))
	for border := pi[len(pi)-1]; border > 0; border = pi[border-1] {
		removed := len(word) - border
		if removed%k == 0 {
			return removed / k
		}
	}
	return (len(word) + k - 1) / k
}
```

### [3034](https://leetcode.com/problems/number-of-subarrays-that-match-a-pattern-i/)

需要模板：`KMP`。

<!-- solution: string_match/kmp/l3034/3034.go; templates: KMP; imports: none -->
```go
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
```

### [3036](https://leetcode.com/problems/number-of-subarrays-that-match-a-pattern-ii/)

需要模板：`KMP`。

<!-- solution: string_match/kmp/l3036/3036.go; templates: KMP; imports: none -->
```go
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
```

## 验证与原有推导

```sh
go test ./string_match/kmp/...
```

题解文件与文档使用相同声明和适配器；修改模板后需要同时验证每题包以及由文档重新组装的代码。原有推导、详细例子和扩展题单保存在 [整理前笔记](legacy/README.before-inline-templates.md)。
