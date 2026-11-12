# 不定长滑动窗口

[返回总目录](../algorithm_templates/README.md)

[扩展练习：迁移点与边界](#extended-practice) · [保留的原分类笔记](#legacy-notes)

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

[练习](l3_template/3.go) · [测试](l3_template/3_test.go)

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

[练习](l1004_template/1004.go) · [测试](l1004_template/1004_test.go)

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

[练习](l209_template/209.go) · [测试](l209_template/209_test.go)

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

[练习](l76_template/76.go) · [测试](l76_template/76_test.go)

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
- 含负数的和通常不满足单调收缩，考虑 [前缀状态](../Subarray/readme.md) 等方法。
- 字节下标不等于任意Unicode字符下标；若转rune，返回字符串时也要处理位置转换。
- 若要统计合法区间数量，见 [atMost](at_most/README.md)，不是每个右端点只加1。

已有分类：[sliding_window](.)。


<a id="legacy-notes"></a>

## 原分类笔记（保留）

# 滑动窗口（Sliding Window）

## 1. 什么时候使用？

滑动窗口维护一个连续区间，通过右端加入、左端移除来增量更新状态，避免反复扫描整个区间。

常见题目特征：连续子数组或子串、固定长度、最多 k 次操作、最短覆盖、满足条件的区间数量。

使用前要判断：移动边界后，条件是否具有足够的单调性？例如正数数组中扩展窗口使和不减，收缩使和不增；包含负数时，209 的普通窗口逻辑不再适用。

本文统一使用闭区间 `[left, right]`，窗口长度是 `right-left+1`。每个元素最多进入、离开一次；若状态更新和检查都是 O(1)，总时间为 O(n)。

## 2. 定长窗口

适用：长度固定为 k 的区间和、字符频次匹配。

```go
left := 0
for right := 0; right < len(nums); right++ {
    // 加入 nums[right]。
    if right-left+1 > k {
        // 移除 nums[left]。
        left++
    }
    if right-left+1 == k {
        // 更新答案或检查是否匹配。
    }
}
```

| 题目 | 状态与关键点 |
| --- | --- |
| [643. Maximum Average Subarray I](https://leetcode.com/problems/maximum-average-subarray-i/) | 维护长度 k 的窗口和，取最大值后除以 k；最大和不能初始化为 0，因为可能全为负数 |
| [567. Permutation in String](https://leetcode.com/problems/permutation-in-string/) | 窗口长度为 len(s1)，比较 26 个小写字母的频次；只需判断是否存在 |
| [438. Find All Anagrams in a String](https://leetcode.com/problems/find-all-anagrams-in-a-string/) | 与 567 相同，但记录所有起点；重叠窗口也要保留 |

对于 567、438，比较 26 项频次是 O(26n)，在固定字母表下为 O(n)，辅助空间 O(26)。不需要生成字符串的全部排列。

## 3. 不定长窗口：最长合法区间

模式：先扩展，条件不合法时持续收缩，恢复合法后更新最大长度。

```go
left, best := 0, 0
for right := 0; right < len(nums); right++ {
    // 加入 nums[right]。
    for /* 窗口不合法 */ {
        // 移除 nums[left]。
        left++
    }
    best = max(best, right-left+1)
}
```

| 题目 | 合法条件 | 常用状态 |
| --- | --- | --- |
| [3. Longest Substring Without Repeating Characters](https://leetcode.com/problems/longest-substring-without-repeating-characters/) | 没有重复字符 | 字符频次或最后出现位置 |
| [1004. Max Consecutive Ones III](https://leetcode.com/problems/max-consecutive-ones-iii/) | 窗口中的 0 不超过 k 个 | 零的数量 |

相关进阶：[904. Fruit Into Baskets](https://leetcode.com/problems/fruit-into-baskets/) 维护最多两种元素；[424. Longest Repeating Character Replacement](https://leetcode.com/problems/longest-repeating-character-replacement/) 使用“窗口长度减去最高字符频次不超过 k”的条件。424 中使用历史最大频次的优化需要单独理解其不变量。

## 4. 不定长窗口：最短满足区间

模式：扩展到满足条件后，持续收缩，直到不满足。答案在移除左端之前更新。

```go
left := 0
for right := 0; right < len(nums); right++ {
    // 加入 nums[right]。
    for /* 窗口满足要求 */ {
        // 更新最短答案。
        // 移除 nums[left]。
        left++
    }
}
```

| 题目 | 满足条件 | 注意点 |
| --- | --- | --- |
| [209. Minimum Size Subarray Sum](https://leetcode.com/problems/minimum-size-subarray-sum/) | 正数窗口和 >= target | 无解返回 0；依赖所有元素为正 |
| [76. Minimum Window Substring](https://leetcode.com/problems/minimum-window-substring/) | 覆盖 t 的全部字符及其出现次数 | 重复字符不能只按种类判断；区分大小写；无解返回空字符串 |

76 可以维护每种字符的需求和窗口计数，再维护已匹配的字符数量或已达标的种类数，避免每次重新扫描全部需求。

## 5. 统计满足条件的子数组数量

### 5.1 最多 / 小于：固定右端点计数

收缩后，若 `[left,right]` 合法，且删除左侧元素不会破坏合法性，则所有以 right 结尾、起点在 left..right 之间的窗口都合法：

```go
answer += right - left + 1
```

- [713. Subarray Product Less Than K](https://leetcode.com/problems/subarray-product-less-than-k/)：元素均为正整数，维护乘积并在 `product >= k` 时收缩；`k <= 1` 时直接返回 0。
- [992. Subarrays with K Different Integers](https://leetcode.com/problems/subarrays-with-k-different-integers/)：使用频次表统计不同元素数量，并转换为两个“最多”计数之差。

### 5.2 恰好 k：两个最多相减

```text
exactly(k) = atMost(k) - atMost(k-1)
```

`atMost(k)` 内部使用最长合法窗口的收缩方式，但累加 `right-left+1`，而不是只更新最大长度。边界 `atMost(0)` 应正确返回 0；通用实现可对负 k 直接返回 0。

相关题：[930. Binary Subarrays With Sum](https://leetcode.com/problems/binary-subarrays-with-sum/)、[1248. Count Number of Nice Subarrays](https://leetcode.com/problems/count-number-of-nice-subarrays/)。这类差分要求统计量是整数且窗口条件具有相应单调性，不能直接套用于任意含负数的数组和问题。

### 5.3 至少覆盖：固定左端点计数

- [1358. Number of Substrings Containing All Three Characters](https://leetcode.com/problems/number-of-substrings-containing-all-three-characters/)：字符串只含 a、b、c，要求三者都出现。

当 `[left,right]` 已覆盖 a、b、c，保持 left 不变，右端点继续向右都有效，所以一次加入：

```go
answer += len(s) - right
```

然后移除左端并继续收缩。每次计入的是当前 left 对应的全部合法右端点，不会重复计数。另一种写法是收缩到不合法后，对当前 right 累加有效起点数量 left；两种计数方式不要混用。

## 6. 滑动窗口 + 单调队列

[239. Sliding Window Maximum](https://leetcode.com/problems/sliding-window-maximum/)：普通窗口不能在移除最大值后 O(1) 找到新最大值，因此维护值单调递减的下标队列。

1. 队首移除已经离开窗口的下标。
2. 队尾移除不大于当前值的旧下标，它们不可能比当前值更优且会更早过期。
3. 当前下标入队。
4. 窗口达到长度 k 后，队首就是最大值。

时间 O(n)，辅助空间 O(k)。已有实现见 `../monotonic_stack/l239/`，本轮不重复生成模板。

## 7. 常见错误

- 混淆“满足时收缩”和“不合法时收缩”：最短覆盖通常是前者，最长合法窗口通常是后者。
- 需要持续修复窗口时误用 if 而非 for；某些单步收缩优化有额外不变量，不要直接混用。
- 移除字符时忘记减少计数，或频次降为 0 时忘记减少不同字符数。
- 76 只检查字符种类，忽略重复需求；567、438 忘记固定窗口长度。
- 数组全负时把最大窗口和初始化成 0。
- 把正数窗口和的算法直接用于含负数的输入。
- Go 的字符串下标是字节；这些字符串题的题目字符集允许按字节处理，泛化到 Unicode 时需先明确 rune 与返回下标的语义。
- 子数组总数可能达到 n(n+1)/2；按题目约束选择计数类型，通用大输入建议使用 int64。

## 8. 目录与练习路线

本轮核心练习均有 `l题号/题号.go` 和 `l题号/题号_test.go`，函数模板使用 `panic("not implemented")`，测试留待实现后运行。

```text
l643 → l567 → l438             定长窗口
l3 → l1004                    最长合法窗口
l209 → l76                    最短满足窗口
l713 → l1358 → l992            子数组计数
```

原有 `1004.go`、`1358.go`、`567.go` 保留不动。它们各有独立的 main 及重复辅助函数，不适合直接一起编译；练习时单独测试 `l题号/` 目录。`567` 和 `438` 在 permutation 目录也有模板，这里按滑动窗口专题独立保留。

无根目录 go.mod 时，可从仓库根目录按文件运行，例如：

```bash
go test sliding_window/l643/*.go
# 仅检查模板和测试是否可编译，不执行测试：
go test -run '^$' sliding_window/l643/*.go
```

## 9. 参考资料

- [Sliding Window for Beginners: Problems, Template, Sample Solutions](https://leetcode.com/discuss/general-discussion/657507/Sliding-Window-for-Beginners-Problems-or-Template-or-Sample-Solutions)
- [How to Solve Sliding Window Problems](https://medium.com/outco/how-to-solve-sliding-window-problems-28d67601a66)

<a id="extended-practice"></a>

## 扩展练习：迁移点与边界

新增题目只补充学习路线与解法总结，不自动生成未请求的实现或测试。

| 题目（原题） | 如何迁移当前模板 | 必须确认的条件 / 边界 |
| --- | --- | --- |
| [904. 水果成篮](https://leetcode.com/problems/fruit-into-baskets/) | 频次表维护最多两种水果，修复后更新最长长度 | 频次归零才减少种数 |
| [424. 替换后的最长重复字符](https://leetcode.com/problems/longest-repeating-character-replacement/) | 窗口长度减最大频次<=k，求最长 | 历史最高频次优化有额外不变量，不能把历史值解释为当前准确频次 |
| [1208. 尽可能使字符串相等](https://leetcode.com/problems/get-equal-substrings-within-budget/) | 每个位置字符差绝对值作为非负费用，维护总费用预算 | 对应位置不可重排；费用非负保证单调 |
| [1658. 将X减到0的最小操作数](https://leetcode.com/problems/minimum-operations-to-reduce-x-to-zero/) | 转为保留和为sum-x的最长正数连续段，答案n-长度 | sum-x<0无解；为0时允许保留空段 |
| [1838. 最高频元素的频数](https://leetcode.com/problems/frequency-of-the-most-frequent-element/) | 排序后把窗口抬到右端最大值，代价=最大值×长度-窗口和 | 排序改变原顺序但本题允许选任意元素；乘法用int64 |
| [2516. 每种字符至少取K个](https://leetcode.com/problems/take-k-of-each-character-from-left-and-right/) | 两端取走等价于保留中间窗口，每字符保留数<=总数-k | 任一字符总数不足k返回-1；k=0返回0 |

**复习方法**：先独立写出状态、合法性或松弛条件，再指出相较本篇基础题改变的那一行；若前提不满足，主动换算法，不为统一模板牺牲正确性。
