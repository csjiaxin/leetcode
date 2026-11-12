# 前缀状态＋哈希表

[返回总目录](../README.md)

## 目录

- [1. 核心思路](#core)
- [2. 560：子数组和为 K](#lc560)
- [3. 974：子数组和可整除](#lc974)
- [4. 525：0 和 1 数量相等](#lc525)
- [5. 1371：元音出现偶数次](#lc1371)
- [6. 1248：恰好 K 个奇数](#lc1248)
- [7. 对比与易错点](#pitfalls)

<a id="core"></a>

## 1. 核心思路

区间 [l,r] 可用两个前缀状态的关系表示。对每个右端点，查询能和当前前缀配对的历史前缀。

```text
初始化空前缀
遍历：更新状态 → 查询历史 → 记录当前
求数量：保存出现次数
求最长：保存第一次出现的位置
```

**不变量**：查询时只使用当前之前的前缀（含空前缀）。允许负数，不依赖窗口和单调性。先推导等式，再写直接循环，无需 step/need 回调。

以下实现分别完整独立，使用标准题目签名；仅展示于文档，不填练习函数。

<a id="lc560"></a>

## 2. [560. 和为 K 的子数组](https://leetcode.com/problems/subarray-sum-equals-k/)

[练习](l560/560.go) · [测试](l560/560_test.go)

**题意**：求非空连续子数组和恰好等于 k 的数量。

**推导**：当前前缀和-历史前缀和=k，所以查询 prefix-k。freq[0]=1 表示空前缀。

```go
package template

func subarraySum(nums []int, k int) int {
	freq := map[int64]int64{0: 1}
	var prefix, answer int64
	for _, x := range nums {
		prefix += int64(x)
		answer += freq[prefix-int64(k)]
		freq[prefix]++
	}
	return int(answer)
}
```

**例子**：[1,-1,1]、k=1。依次查询0、-1、0，命中次数1、0、2，答案3。最后一次历史前缀0已出现两次。

**复杂度**：平均时间 O(n)，空间 O(n)。k=0 时也先查后记，避免把空区间计入。

<a id="lc974"></a>

## 3. [974. 和可被 K 整除的子数组](https://leetcode.com/problems/subarray-sums-divisible-by-k/)

[练习](l974/974.go) · [测试](l974/974_test.go)

**题意**：统计和能被 k 整除的非空子数组，题设 k>0。

**推导**：两前缀相减能被 k 整除，等价于规范化余数相同。Go 的负余数要加 k，统一到 [0,k)。

```go
package template

func subarraysDivByK(nums []int, k int) int {
	freq := map[int64]int64{0: 1}
	var prefix, answer int64
	for _, x := range nums {
		prefix += int64(x)
		remainder := prefix % int64(k)
		if remainder < 0 {
			remainder += int64(k)
		}
		answer += freq[remainder]
		freq[remainder]++
	}
	return int(answer)
}
```

**例子**：[-1,1]、k=2。两个前缀余数为1、0；后者与空前缀匹配，答案1。

**复杂度**：平均时间 O(n)，空间 O(min(n+1,k))。不必构造复杂的模加法，int64 前缀和足够覆盖本题约束。

<a id="lc525"></a>

## 4. [525. 连续数组](https://leetcode.com/problems/contiguous-array/)

[练习](l525/525.go) · [测试](l525/525_test.go)

**题意**：0/1 数组中，求0与1数量相等的最长连续区间。

**状态**：0记-1、1记+1，balance相同的两前缀之间数量平衡。first[0]=-1；只保留最早位置，才能得到最长距离。

```go
package template

func findMaxLength(nums []int) int {
	first := map[int]int{0: -1}
	balance, answer := 0, 0
	for i, x := range nums {
		if x == 0 {
			balance--
		} else {
			balance++
		}
		if j, ok := first[balance]; ok {
			answer = max(answer, i-j)
		} else {
			first[balance] = i
		}
	}
	return answer
}
```

**例子**：[0,1,0,1]。balance=0 首次位置始终为-1，最后得到3-(-1)=4；覆盖最早位置会漏掉这个答案。

**复杂度**：平均时间 O(n)，空间 O(n)。

<a id="lc1371"></a>

## 5. [1371. 每个元音包含偶数次的最长子字符串](https://leetcode.com/problems/find-the-longest-substring-containing-vowels-in-even-counts/)

[练习](l1371/1371.go) · [测试](l1371/1371_test.go)

**题意**：求最长子串，使 a/e/i/o/u 各自出现偶数次，0次也合法。

**状态**：用五位掩码记录奇偶性，遇元音翻转对应位。两个前缀掩码相同，区间内每个元音都出现偶数次。与525一样保存最早下标。

```go
package template

func findTheLongestSubstring(s string) int {
	first := map[int]int{0: -1}
	mask, answer := 0, 0
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case 'a':
			mask ^= 1 << 0
		case 'e':
			mask ^= 1 << 1
		case 'i':
			mask ^= 1 << 2
		case 'o':
			mask ^= 1 << 3
		case 'u':
			mask ^= 1 << 4
		}
		if j, ok := first[mask]; ok {
			answer = max(answer, i-j)
		} else {
			first[mask] = i
		}
	}
	return answer
}
```

**例子**："aba" 的掩码依次为1、1、0，最后与空前缀0匹配，长度3。辅音不改掩码，但仍计入长度。

**复杂度**：时间 O(n)，空间 O(32)=O(1)。题设为小写英文字母，按字节下标即为字符位置。

<a id="lc1248"></a>

## 6. [1248. 统计优美子数组](https://leetcode.com/problems/count-number-of-nice-subarrays/)

[练习](l1248/1248.go) · [测试](l1248/1248_test.go)

**题意**：统计恰好有 k 个奇数的连续子数组。

**状态**：odd 为前缀奇数数目，查询 odd-k 的历史次数。它是560的“每个元素先变成0/1”的版本。

```go
package template

func numberOfSubarrays(nums []int, k int) int {
	freq := map[int]int64{0: 1}
	odd := 0
	var answer int64
	for _, x := range nums {
		if x%2 != 0 {
			odd++
		}
		answer += freq[odd-k]
		freq[odd]++
	}
	return int(answer)
}
```

**例子**：[2,1,2]、k=1，答案4：左右偶数都可带或不带。也可用 [atMost相减](../at_most/README.md#lc1248)，本节只展示前缀版本。

**复杂度**：平均时间 O(n)，空间 O(n)。

<a id="pitfalls"></a>

## 7. 对比与易错点

| 题目 | 状态 | 查询历史 | 保存内容 |
| --- | --- | --- | --- |
| 560 | 前缀和 | prefix-k | 次数 |
| 974 | 前缀余数 | 相同余数 | 次数 |
| 525 | 1数量-0数量 | 相同balance | 最早下标 |
| 1371 | 元音奇偶掩码 | 相同mask | 最早下标 |
| 1248 | 前缀奇数数目 | odd-k | 次数 |

- 求数量初始化 freq[0]=1；求最长初始化 first[0]=-1，不要混淆。
- 数量问题先查询再记录，否则可能计空区间；最长问题不要覆盖最早下标。
- map 中的0可能是合法下标，判断存在须用 ok。
- 计数和前缀和使用 int64 中间值；本题答案范围允许返回 int。若通用场景答案更大，应连返回类型也改成 int64。
- 前缀状态可以处理负数；不能据此推论含负数的和也能用普通滑窗。
- 区间最大值一般不能由两个前缀最大值相减得到。

已有分类：[滑动窗口](../../sliding_window/)、[子数组](../../Subarray/)。
