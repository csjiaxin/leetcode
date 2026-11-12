# 不定长滑动窗口

[返回总目录](../README.md)

## 目录

- [1. 核心思路](#core)
- [2. 3：最长无重复子串](#lc3)
- [3. 1004：最长连续1](#lc1004)
- [4. 209：最短正数和窗口](#lc209)
- [5. 76：最小覆盖子串](#lc76)
- [6. 对比与易错点](#pitfalls)

<a id="core"></a>

## 1. 核心思路

左右端点只向右移动，复用窗口状态。前提是收缩能单调修复限制，而不是所有连续区间题都能滑窗。

```text
最长合法：加入右端 → while 违规则移除左端 → 更新最大长度
最短满足：加入右端 → while 满足则先记录答案、再移除左端
```

每个元素最多加入、移除一次，内层while不使总复杂度变为平方。下面不再封装 add/remove/valid，每题直接写循环；各代码块独立，练习文件不改动。

<a id="lc3"></a>

## 2. [3. 无重复字符的最长子串](https://leetcode.com/problems/longest-substring-without-repeating-characters/)

[练习](l3/3.go) · [测试](l3/3_test.go)

**题意**：求没有重复字符的最长连续子串长度。

**状态与修复**：窗口加入新字符前无重复，所以只可能新字符重复。只要它频次>1，就不断移除左端，恢复合法后更新长度。

```go
package template

func lengthOfLongestSubstring(s string) int {
	var freq [256]int
	left, answer := 0, 0
	for right := 0; right < len(s); right++ {
		freq[s[right]]++
		for freq[s[right]] > 1 {
			freq[s[left]]--
			left++
		}
		answer = max(answer, right-left+1)
	}
	return answer
}
```

**例子**："abba" 加入第二个b时，需要连续移除a、b；只写if会留下非法窗口。答案2。

**复杂度**：时间 O(n)，字节频次数组空间 O(1)。本题按ASCII字节处理；Unicode字符需转rune并统一下标含义。

<a id="lc1004"></a>

## 3. [1004. 最大连续1的个数 III](https://leetcode.com/problems/max-consecutive-ones-iii/)

[练习](l1004/1004.go) · [测试](l1004/1004_test.go)

**题意**：最多把k个0翻成1，求最长连续1长度。

**状态**：zeros为窗口0的数目。zeros<=k时整段可翻成1，违规时收缩，合法后更新。

```go
package template

func longestOnes(nums []int, k int) int {
	left, zeros, answer := 0, 0, 0
	for right, x := range nums {
		if x == 0 {
			zeros++
		}
		for zeros > k {
			if nums[left] == 0 {
				zeros--
			}
			left++
		}
		answer = max(answer, right-left+1)
	}
	return answer
}
```

**例子**：[1,0,1,1]、k=0，零必须移出，后段[1,1]长度2。

**复杂度**：时间 O(n)，空间 O(1)。题设k>=0，空窗口也合法。

<a id="lc209"></a>

## 4. [209. 长度最小的子数组](https://leetcode.com/problems/minimum-size-subarray-sum/)

[练习](l209/209.go) · [测试](l209/209_test.go)

**题意**：正整数数组中找和>=target的最短连续子数组，不存在返回0。

**状态与不变量**：sum为窗口和；满足时先记录长度再收缩。正数保证移除左端只会减少和。

```go
package template

func minSubArrayLen(target int, nums []int) int {
	left, answer := 0, len(nums)+1
	var sum int64
	for right, x := range nums {
		sum += int64(x)
		for sum >= int64(target) {
			answer = min(answer, right-left+1)
			sum -= int64(nums[left])
			left++
		}
	}
	if answer == len(nums)+1 {
		return 0
	}
	return answer
}
```

**例子**：[2,3,1,2,4,3]、target=7，最短[4,3]长度2。题设target>0、元素为正。

**复杂度**：时间 O(n)，空间 O(1)。负数反例 [1,-1,5]、target=5，普通滑窗可能漏掉[5]。

<a id="lc76"></a>

## 5. [76. 最小覆盖子串](https://leetcode.com/problems/minimum-window-substring/)

[练习](l76/76.go) · [测试](l76/76_test.go)

**题意**：找最短子串，覆盖t的全部字符，重复字符必须满足次数。

**状态**：need、window记录频次；required是目标字符种数，formed是已满足次数要求的种数。只有跨过频次临界点才修改formed。

```go
package template

func minWindow(s string, t string) string {
	if len(t) == 0 || len(t) > len(s) {
		return ""
	}
	var need, window [256]int
	required := 0
	for i := 0; i < len(t); i++ {
		if need[t[i]] == 0 {
			required++
		}
		need[t[i]]++
	}
	left, formed := 0, 0
	bestStart, bestLength := 0, len(s)+1
	for right := 0; right < len(s); right++ {
		c := s[right]
		window[c]++
		if need[c] > 0 && window[c] == need[c] {
			formed++
		}
		for formed == required {
			if right-left+1 < bestLength {
				bestStart, bestLength = left, right-left+1
			}
			c = s[left]
			if need[c] > 0 && window[c] == need[c] {
				formed--
			}
			window[c]--
			left++
		}
	}
	if bestLength == len(s)+1 {
		return ""
	}
	return s[bestStart : bestStart+bestLength]
}
```

**例子**：s="ADOBECODEBANC"、t="ABC"，答案"BANC"。t="AA"时必须有两个A，不是见到A就认为满足。

**复杂度**：时间 O(len(s)+len(t))，字节数组空间 O(1)。移除前检查是否刚好达到需求；多余副本被删时不能减少formed。

<a id="pitfalls"></a>

## 6. 对比与易错点

| 题目 | 窗口状态 | while条件 | 更新答案 |
| --- | --- | --- | --- |
| 3 | 字符频次 | 新字符重复 | 修复后求最长 |
| 1004 | zeros | zeros>k | 修复后求最长 |
| 209 | sum | sum>=target | 移除前求最短 |
| 76 | formed、required | formed==required | 移除前求最短 |

- 不要只用一次if收缩；一次移除未必够。
- 最长合法与最短满足不是同一个while条件，更新时机也不同。
- 含负数的和通常不满足单调收缩，考虑 [前缀状态](../prefix_state_hash/README.md) 等方法。
- 字节下标不等于任意Unicode字符下标；若转rune，返回字符串时也要处理位置转换。
- 若要统计合法区间数量，见 [atMost](../at_most/README.md)，不是每个右端点只加1。

已有分类：[sliding_window](../../sliding_window/)。
