# 前缀状态＋哈希：频次与最早位置两个模板

[返回模板总目录](../../algorithm_templates/README.md)

有效区间必须能表示为“历史状态 == needed(当前状态)”的单键关系。Count 先查频次再记录当前前缀；Longest 先用最早位置更新长度，再仅在首次出现时记录。空前缀分别初始化为频次 1、位置 -1。

[560](https://leetcode.com/problems/subarray-sum-equals-k/) 把和转换为 sum-k；[974](https://leetcode.com/problems/subarray-sums-divisible-by-k/) 规范化负数余数；[1248](https://leetcode.com/problems/count-number-of-nice-subarrays/) 数奇数；[525](https://leetcode.com/problems/contiguous-array/) 保存 1 与 0 的数量差；[1371](https://leetcode.com/problems/find-the-longest-substring-containing-vowels-in-even-counts/) 保存元音奇偶 mask。需要多个历史键、大小关系或区间最值时，需换聚合方式。

## 复制约定

每题 `.go` 文件已内嵌所需模板，无需导入仓库的算法包。文档将固定模板只列一次，逐题展示适配器；每个适配器与它列出的模板组成一个独立代码组。不同题的同名 Constructor、类型或入口请分别使用。将列出的标准库 import 放在文件顶部，再放模板与适配器；复制提交时优先直接使用对应题目的完整 `.go`。

本组使用 Go 1.21+ 的 min/max 和 Go 1.18+ 的泛型。平台已提供的 ListNode 等类型沿用平台定义。

## 固定模板

### Count

<!-- template: Count -->
```go
// Count 统计满足状态关系的非空连续区间数量。
// advance 推进前缀状态；needed 返回当前状态应匹配的历史状态。
// 有效区间必须能表示为 previousState == needed(currentState) 的单键匹配。
// 计数必须先查询、后记录，避免把当前前缀与自身配成空区间。
func Count[T any, S comparable](values []T, initial S, advance func(S, T) S, needed func(S) S) int {
	freq := map[S]int{initial: 1}
	state, answer := initial, 0
	for _, value := range values {
		state = advance(state, value)
		answer += freq[needed(state)]
		freq[state]++
	}
	return answer
}
```

### Longest

<!-- template: Longest -->
```go
// Longest 返回满足同类状态关系的最长连续区间长度。
// earliest 保存前缀第一次出现的位置，空前缀位于 -1；不得覆盖较早位置。
// 有效区间同样必须对应一个历史键，而非任意区间关系或多个可接受键。
func Longest[T any, S comparable](values []T, initial S, advance func(S, T) S, needed func(S) S) int {
	earliest := map[S]int{initial: -1}
	state, answer := initial, 0
	for i, value := range values {
		state = advance(state, value)
		if position, ok := earliest[needed(state)]; ok && i-position > answer {
			answer = i - position
		}
		if _, exists := earliest[state]; !exists {
			earliest[state] = i
		}
	}
	return answer
}
```

## 题目索引

| 题目 | 模板 | 本地代码与测试 |
| --- | --- | --- |
| [525](https://leetcode.com/problems/contiguous-array/) | Longest | [代码](l525/525.go) · [测试](l525/525_test.go) |
| [560](https://leetcode.com/problems/subarray-sum-equals-k/) | Count | [代码](l560/560.go) · [测试](l560/560_test.go) |
| [974](https://leetcode.com/problems/subarray-sums-divisible-by-k/) | Count | [代码](l974/974.go) · [测试](l974/974_test.go) |
| [1248](https://leetcode.com/problems/count-number-of-nice-subarrays/) | Count | [代码](l1248/1248.go) · [测试](l1248/1248_test.go) |
| [1371](https://leetcode.com/problems/find-the-longest-substring-containing-vowels-in-even-counts/) | Longest | [代码](l1371/1371.go) · [测试](l1371/1371_test.go) |

## 逐题适配器

### [525](https://leetcode.com/problems/contiguous-array/)

需要模板：`Longest`。

<!-- solution: prefix_sum/state_hash/l525/525.go; templates: Longest; imports: none -->
```go
func findMaxLength(nums []int) int {
	return Longest(nums, 0, func(balance, value int) int { return balance + 2*value - 1 }, func(balance int) int { return balance })
}
```

### [560](https://leetcode.com/problems/subarray-sum-equals-k/)

需要模板：`Count`。

<!-- solution: prefix_sum/state_hash/l560/560.go; templates: Count; imports: none -->
```go
func subarraySum(nums []int, k int) int {
	return Count(nums, 0, func(sum, value int) int { return sum + value }, func(sum int) int { return sum - k })
}
```

### [974](https://leetcode.com/problems/subarray-sums-divisible-by-k/)

需要模板：`Count`。

<!-- solution: prefix_sum/state_hash/l974/974.go; templates: Count; imports: none -->
```go
func subarraysDivByK(nums []int, k int) int {
	return Count(nums, 0, func(remainder, value int) int { return ((remainder+value)%k + k) % k }, func(remainder int) int { return remainder })
}
```

### [1248](https://leetcode.com/problems/count-number-of-nice-subarrays/)

需要模板：`Count`。

<!-- solution: prefix_sum/state_hash/l1248/1248.go; templates: Count; imports: none -->
```go
func numberOfSubarrays(nums []int, k int) int {
	return Count(nums, 0, func(odds, value int) int { return odds + value%2 }, func(odds int) int { return odds - k })
}
```

### [1371](https://leetcode.com/problems/find-the-longest-substring-containing-vowels-in-even-counts/)

需要模板：`Longest`。

<!-- solution: prefix_sum/state_hash/l1371/1371.go; templates: Longest; imports: none -->
```go
func findTheLongestSubstring(s string) int {
	return Longest([]byte(s), 0, func(mask int, value byte) int {
		for i, vowel := range []byte("aeiou") {
			if value == vowel {
				return mask ^ (1 << i)
			}
		}
		return mask
	}, func(mask int) int { return mask })
}
```

## 验证与原有推导

```sh
go test ./prefix_sum/state_hash/...
```

题解文件与文档使用相同声明和适配器；修改模板后需要同时验证每题包以及由文档重新组装的代码。原有推导、详细例子和扩展题单保存在 [整理前笔记](legacy/README.before-inline-templates.md)。
