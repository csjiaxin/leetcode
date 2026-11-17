# 不定长滑动窗口

[返回总目录](../algorithm_templates/README.md)

[扩展练习：迁移点与边界](#extended-practice) · [保留的原分类笔记](#legacy-notes)

## 同文件局部模板（先复制，再复制任一题适配器）

所有下文 Go 适配器共用本代码块；不需要任何 import。最长与最短的答案时机、定长与 atMost 的计数方式分别命名，不用一个反转含义的 valid 回调混用。

```go
// windowLongest maintains a valid suffix after removing invalid left endpoints.
func windowLongest(n int, add, remove func(int), invalid func() bool) int {
	left, best := 0, 0
	for right := 0; right < n; right++ {
		add(right)
		for left <= right && invalid() {
			remove(left)
			left++
		}
		if length := right - left + 1; length > best {
			best = length
		}
	}
	return best
}

// windowShortest records valid windows before shrinking; length n+1 means absent.
func windowShortest(n int, add, remove func(int), valid func() bool) (int, int) {
	left, start, length := 0, 0, n+1
	for right := 0; right < n; right++ {
		add(right)
		for left <= right && valid() {
			if right-left+1 < length {
				start, length = left, right-left+1
			}
			remove(left)
			left++
		}
	}
	return start, length
}

// windowFixed visits exactly the windows of length k.
func windowFixed(n, k int, add, remove func(int), visit func(int)) {
	if k <= 0 || k > n {
		return
	}
	left := 0
	for right := 0; right < n; right++ {
		add(right)
		if right-left+1 > k {
			remove(left)
			left++
		}
		if right-left+1 == k {
			visit(left)
		}
	}
}

// windowAtMost counts valid suffixes; removing the left endpoint preserves validity.
func windowAtMost(n int, add, remove func(int), invalid func() bool) int {
	left, count := 0, 0
	for right := 0; right < n; right++ {
		add(right)
		for left <= right && invalid() {
			remove(left)
			left++
		}
		count += right - left + 1
	}
	return count
}

// windowCoverCount counts all future right endpoints for each valid left endpoint.
func windowCoverCount(n int, add, remove func(int), valid func() bool) int {
	left, count := 0, 0
	for right := 0; right < n; right++ {
		add(right)
		for left <= right && valid() {
			count += n - right
			remove(left)
			left++
		}
	}
	return count
}

// windowMaximum uses a decreasing deque of indices with explicit expiration.
func windowMaximum(nums []int, k int) []int {
	result := []int{}
	if k <= 0 || k > len(nums) {
		return result
	}
	deque := []int{}
	for right, value := range nums {
		if len(deque) > 0 && deque[0] <= right-k {
			deque = deque[1:]
		}
		for len(deque) > 0 && nums[deque[len(deque)-1]] <= value {
			deque = deque[:len(deque)-1]
		}
		deque = append(deque, right)
		if right+1 >= k {
			result = append(result, nums[deque[0]])
		}
	}
	return result
}
```

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

每个元素最多加入、移除一次，内层while不使总复杂度变为平方。下列六个局部模板分别固定移动规则，逐题仅适配 add/remove 和条件。每份实际 Go 题解也在同一文件内嵌对应模板，可独立提交，不导入仓库本地包。

<a id="lc3"></a>

## 2. [3. 无重复字符的最长子串](https://leetcode.com/problems/longest-substring-without-repeating-characters/)

[练习](l3/variants/template/3.go) · [测试](l3/variants/template/3_test.go)

**题意**：求没有重复字符的最长连续子串长度。

**状态与修复**：窗口加入新字符前无重复，所以只可能新字符重复。只要它频次>1，就不断移除左端，恢复合法后更新长度。

```go
func lengthOfLongestSubstring(s string) int {
	var count [256]int
	duplicates := 0
	return windowLongest(len(s),
		func(i int) {
			count[s[i]]++
			if count[s[i]] == 2 {
				duplicates++
			}
		},
		func(i int) {
			if count[s[i]] == 2 {
				duplicates--
			}
			count[s[i]]--
		},
		func() bool { return duplicates > 0 })
}
```

**例子**："abba" 加入第二个b时，需要连续移除a、b；只写if会留下非法窗口。答案2。

**复杂度**：时间 O(n)，字节频次数组空间 O(1)。本题按ASCII字节处理；Unicode字符需转rune并统一下标含义。

<a id="lc1004"></a>

## 3. [1004. 最大连续1的个数 III](https://leetcode.com/problems/max-consecutive-ones-iii/)

[练习](l1004/variants/template/1004.go) · [测试](l1004/variants/template/1004_test.go)

**题意**：最多把k个0翻成1，求最长连续1长度。

**状态**：zeros为窗口0的数目。zeros<=k时整段可翻成1，违规时收缩，合法后更新。

```go
func longestOnes(nums []int, k int) int {
	zeros := 0
	return windowLongest(len(nums),
		func(i int) {
			if nums[i] == 0 {
				zeros++
			}
		},
		func(i int) {
			if nums[i] == 0 {
				zeros--
			}
		},
		func() bool { return zeros > k })
}
```

**例子**：[1,0,1,1]、k=0，零必须移出，后段[1,1]长度2。

**复杂度**：时间 O(n)，空间 O(1)。题设k>=0，空窗口也合法。

<a id="lc209"></a>

## 4. [209. 长度最小的子数组](https://leetcode.com/problems/minimum-size-subarray-sum/)

[练习](l209/variants/template/209.go) · [测试](l209/variants/template/209_test.go)

**题意**：正整数数组中找和>=target的最短连续子数组，不存在返回0。

**状态与不变量**：sum为窗口和；满足时先记录长度再收缩。正数保证移除左端只会减少和。

```go
func minSubArrayLen(target int, nums []int) int {
	var sum int64
	_, length := windowShortest(len(nums),
		func(i int) { sum += int64(nums[i]) },
		func(i int) { sum -= int64(nums[i]) },
		func() bool { return sum >= int64(target) })
	if length == len(nums)+1 {
		return 0
	}
	return length
}
```

**例子**：[2,3,1,2,4,3]、target=7，最短[4,3]长度2。题设target>0、元素为正。

**复杂度**：时间 O(n)，空间 O(1)。负数反例 [1,-1,5]、target=5，普通滑窗可能漏掉[5]。

<a id="lc76"></a>

## 5. [76. 最小覆盖子串](https://leetcode.com/problems/minimum-window-substring/)

[练习](l76/variants/template/76.go) · [测试](l76/variants/template/76_test.go)

**题意**：找最短子串，覆盖t的全部字符，重复字符必须满足次数。

**状态**：need、window记录频次；required是目标字符种数，formed是已满足次数要求的种数。只有跨过频次临界点才修改formed。

```go
func minWindow(s string, t string) string {
	if len(t) == 0 || len(t) > len(s) {
		return ""
	}
	var need, count [256]int
	missing := len(t)
	for i := range t {
		need[t[i]]++
	}
	start, length := windowShortest(len(s),
		func(i int) {
			c := s[i]
			if count[c] < need[c] {
				missing--
			}
			count[c]++
		},
		func(i int) {
			c := s[i]
			count[c]--
			if count[c] < need[c] {
				missing++
			}
		},
		func() bool { return missing == 0 })
	if length == len(s)+1 {
		return ""
	}
	return s[start : start+length]
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
- 含负数的和通常不满足单调收缩，考虑 [前缀状态](../prefix_sum/state_hash/readme.md) 等方法。
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

```text
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

```text
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

```text
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

```text
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

```text
answer += len(s) - right
```

然后移除左端并继续收缩。每次计入的是当前 left 对应的全部合法右端点，不会重复计数。另一种写法是收缩到不合法后，对当前 right 累加有效起点数量 left；两种计数方式不要混用。

## 6. 滑动窗口 + 单调队列

[239. Sliding Window Maximum](https://leetcode.com/problems/sliding-window-maximum/)：普通窗口不能在移除最大值后 O(1) 找到新最大值，因此维护值单调递减的下标队列。

1. 队首移除已经离开窗口的下标。
2. 队尾移除不大于当前值的旧下标，它们不可能比当前值更优且会更早过期。
3. 当前下标入队。
4. 窗口达到长度 k 后，队首就是最大值。

时间 O(n)，辅助空间 O(k)。实际实现见 [l239](l239)，使用本页独立的 windowMaximum 单调队列模板。

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

本轮核心题目均有 `l题号/题号.go` 和 `l题号/题号_test.go`，实际实现及各来源变体均使用本页命名模板，并运行各自原测试。

```text
l643 → l567 → l438             定长窗口
l3 → l1004                    最长合法窗口
l209 → l76                    最短满足窗口
l713 → l1358 → l992            子数组计数
```

原演示与练习现按题号归并，各来源保留独立包；567、438 的 permutation 来源在各题 variants/from_permutation 中，实际实现也使用同一 windowFixed 模板。

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

## 目录归并：一个题号入口，保留独立练习

本目录现在集中管理原来散落在 `two_pointers`、`permutation` 和 `at_most` 下的窗口题。重复源码与测试完整放入 `l题号/variants/来源` 独立包，不覆盖函数、不把未完成练习的测试改为调用另一份答案。

| 模板 | 必须维持的不变量与答案时机 | 题目入口 |
| --- | --- | --- |
| 最长合法窗口 | 加右端；`while invalid` 移左端；修复后更新最长 | [3](l3)、[1004](l1004)、[424](l424)、[904](l904) |
| 最短覆盖窗口 | 加右端；`while valid` 先记录答案再移左端 | [209](l209)、[76](l76) |
| 定长窗口 | 加右端；超过指定长度才移左端；长度恰好时统计 | [438](l438)、[567](l567)、[643](l643) |
| 向内收缩保持合法的区间计数 | 收缩到最小合法 left 后，固定 right 有 `right-left+1` 个合法起点 | [713](l713) |
| 满足覆盖的区间计数 | 满足时固定 left，并将当前及以后所有 right 计入，再继续移左端 | [1358](l1358) |
| 至多转恰好 | `exactly(k)=atMost(k)-atMost(k-1)`；atMost 的所有更短后缀也合法 | [992](l992)、[930](l930)、[1248](l1248) |
| 单调队列 | 队首是有效窗口内最大值；队首过期、队尾淘汰劣值是两种操作 | [239](l239) |

其中 1358 的常见实现也可在满足条件时累计后续右端点数量；不能不检查原实现的计数方向就改成统一回调。930 使用非负二进制和，1248 使用奇数数量，992 使用不同元素数量；它们共享 atMost 代数关系，但维护的状态不同。`atMost(k<0)` 应为 0。

```text
最长：add(right) → while 不合法：remove(left++) → 记录长度
最短：add(right) → while 合法：记录长度 → remove(left++)
定长：add(right) → 超长则 remove(left++) → 满长时记录
计数：恢复合法 → answer += right-left+1
```

这些循环分别成立的前提不同，不用一个含糊的 `valid` 回调同时驱动。209 依赖正数，713 依赖正整数并需考虑 k<=1；含负数的普通区间和不能直接套用。

当前变体入口：

| 题号 | 保留的独立来源 |
| --- | --- |
| 3、76、209、1004 | 各题的 `variants/template` 和 `variants/from_two_pointers` |
| 424、713 | 各题的 `variants/from_two_pointers` |
| 438、567 | 各题的 `variants/from_two_pointers`、`variants/from_permutation` |
| 992 | `variants/from_two_pointers`、`variants/from_at_most` |

[567/variants/from_permutation](l567/variants/from_permutation) 保留了一份已实现的定长频次窗口；[567 主练习](l567) 现同样使用 windowFixed；各变体与原测试保持独立。所有 `variants` 现按相同子模板独立实现，并保留各自原测试。[at_most/README.md](at_most/README.md) 继续作为公式专题笔记，但题目代码已归入统一题号目录。

单测按当前包或整个专题运行：`go test ./sliding_window/...`。目录归并与算法实现验证是两个独立步骤。迁移清单见 [reorganization.json](../two_pointers/reorganization.json)。


## 其他题的同模板适配器

以下适配器使用上方已经声明的同名模板；分别复制到对应题的同一文件即可。

### 239：maxSlidingWindow

```go
func maxSlidingWindow(nums []int, k int) []int { return windowMaximum(nums, k) }
```

### 424：characterReplacement

```go
func characterReplacement(s string, k int) int {
	var count [26]int
	length := 0
	return windowLongest(len(s),
		func(i int) { count[s[i]-'A']++; length++ },
		func(i int) { count[s[i]-'A']--; length-- },
		func() bool {
			highest := 0
			for _, c := range count {
				if c > highest {
					highest = c
				}
			}
			return length-highest > k
		})
}
```

### 438：findAnagrams

```go
func findAnagrams(s string, p string) []int {
	result := []int{}
	if len(p) == 0 {
		return result
	}
	var need, count [256]int
	for i := range p {
		need[p[i]]++
	}
	windowFixed(len(s), len(p),
		func(i int) { count[s[i]]++ },
		func(i int) { count[s[i]]-- },
		func(left int) {
			if count == need {
				result = append(result, left)
			}
		})
	return result
}
```

### 567：checkInclusion

```go
func checkInclusion(s1 string, s2 string) bool {
	if len(s1) == 0 {
		return true
	}
	var need, count [256]int
	for i := range s1 {
		need[s1[i]]++
	}
	found := false
	windowFixed(len(s2), len(s1),
		func(i int) { count[s2[i]]++ },
		func(i int) { count[s2[i]]-- },
		func(left int) {
			if count == need {
				found = true
			}
		})
	return found
}
```

### 643：findMaxAverage

```go
func findMaxAverage(nums []int, k int) float64 {
	if k <= 0 || k > len(nums) {
		return 0
	}
	var sum, best int64
	first := true
	windowFixed(len(nums), k,
		func(i int) { sum += int64(nums[i]) },
		func(i int) { sum -= int64(nums[i]) },
		func(left int) {
			if first || sum > best {
				best = sum
				first = false
			}
		})
	return float64(best) / float64(k)
}
```

### 713：numSubarrayProductLessThanK

```go
func numSubarrayProductLessThanK(nums []int, k int) int {
	if k <= 1 {
		return 0
	}
	product := int64(1)
	return windowAtMost(len(nums),
		func(i int) { product *= int64(nums[i]) },
		func(i int) { product /= int64(nums[i]) },
		func() bool { return product >= int64(k) })
}
```

### 904：totalFruit

```go
func totalFruit(nums []int) int {
	count := map[int]int{}
	return windowLongest(len(nums),
		func(i int) { count[nums[i]]++ },
		func(i int) {
			count[nums[i]]--
			if count[nums[i]] == 0 {
				delete(count, nums[i])
			}
		},
		func() bool { return len(count) > 2 })
}
```

### 930：numSubarraysWithSum

```go
func numSubarraysWithSum(nums []int, goal int) int {
	atMost := func(limit int) int {
		if limit < 0 {
			return 0
		}
		sum := 0
		return windowAtMost(len(nums),
			func(i int) { sum += nums[i] },
			func(i int) { sum -= nums[i] },
			func() bool { return sum > limit })
	}
	return atMost(goal) - atMost(goal-1)
}
```

### 992：subarraysWithKDistinct

```go
func subarraysWithKDistinct(nums []int, k int) int {
	atMost := func(limit int) int {
		if limit < 0 {
			return 0
		}
		count := map[int]int{}
		return windowAtMost(len(nums),
			func(i int) { count[nums[i]]++ },
			func(i int) {
				count[nums[i]]--
				if count[nums[i]] == 0 {
					delete(count, nums[i])
				}
			},
			func() bool { return len(count) > limit })
	}
	return atMost(k) - atMost(k-1)
}
```

### 1248：numberOfSubarrays

```go
func numberOfSubarrays(nums []int, k int) int {
	atMost := func(limit int) int {
		if limit < 0 {
			return 0
		}
		odd := 0
		return windowAtMost(len(nums),
			func(i int) { odd += nums[i] & 1 },
			func(i int) { odd -= nums[i] & 1 },
			func() bool { return odd > limit })
	}
	return atMost(k) - atMost(k-1)
}
```

### 1358：numberOfSubstrings

```go
func numberOfSubstrings(s string) int {
	var count [3]int
	return windowCoverCount(len(s),
		func(i int) { count[s[i]-'a']++ },
		func(i int) { count[s[i]-'a']-- },
		func() bool { return count[0] > 0 && count[1] > 0 && count[2] > 0 })
}
```
