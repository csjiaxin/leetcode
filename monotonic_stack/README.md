# 单调栈：最近边界与贡献结算

## 同文件命名模板与适配器

先复制本代码块，再复制任一题的适配器；实际题解文件同样内嵌 helper，保留官方函数签名。只使用标准库，不导入仓库本地模板包。本文所有 Go 代码块可按顺序组合成一个包；不要重复复制 helper。

```go
import "sort"

// nextGreaterPositions returns the first strictly greater position, or -1.
// A circular second pass resolves pending originals without pushing copies.
func nextGreaterPositions(nums []int, circular bool) []int {
	n := len(nums)
	next := make([]int, n)
	for i := range next {
		next[i] = -1
	}
	stack := []int{}
	rounds := n
	if circular {
		rounds *= 2
	}
	for step := 0; step < rounds; step++ {
		current := step % n
		for len(stack) > 0 && nums[stack[len(stack)-1]] < nums[current] {
			previous := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			next[previous] = current
		}
		if step < n {
			stack = append(stack, current)
		}
	}
	return next
}

// monotonicBoundaries settles right boundaries on pop and reads left on push.
// pop defines the strict/non-strict tie ownership for this specific problem.
func monotonicBoundaries(nums []int, pop func(previous, current int) bool) ([]int, []int) {
	n := len(nums)
	left, right := make([]int, n), make([]int, n)
	stack := []int{}
	for i, x := range nums {
		right[i] = n
		for len(stack) > 0 && pop(nums[stack[len(stack)-1]], x) {
			j := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			right[j] = i
		}
		left[i] = -1
		if len(stack) > 0 {
			left[i] = stack[len(stack)-1]
		}
		stack = append(stack, i)
	}
	return left, right
}

func histogramArea(heights []int) int {
	left, right := monotonicBoundaries(heights, func(previous, current int) bool { return previous > current })
	var best int64
	for i, h := range heights {
		area := int64(h) * int64(right[i]-left[i]-1)
		if area > best {
			best = area
		}
	}
	return int(best)
}

// greedySubsequence keeps k elements; better may replace a worse previous value.
func greedySubsequence[T any](values []T, k int, better func(current, previous T) bool) []T {
	if k < 0 {
		k = 0
	}
	if k > len(values) {
		k = len(values)
	}
	drop := len(values) - k
	stack := make([]T, 0, len(values))
	for _, value := range values {
		for drop > 0 && len(stack) > 0 && better(value, stack[len(stack)-1]) {
			stack = stack[:len(stack)-1]
			drop--
		}
		stack = append(stack, value)
	}
	return stack[:k]
}

// trappedLayers settles a basin bottom only when both bounding walls exist.
func trappedLayers(heights []int) int {
	stack := []int{}
	water := 0
	for right, h := range heights {
		for len(stack) > 0 && heights[stack[len(stack)-1]] < h {
			bottom := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				break
			}
			left := stack[len(stack)-1]
			level := heights[left]
			if h < level {
				level = h
			}
			water += (right - left - 1) * (level - heights[bottom])
		}
		stack = append(stack, right)
	}
	return water
}

// distinctSubsequence pops a larger symbol only when another copy remains.
func distinctSubsequence(s string) string {
	var remaining [256]int
	var used [256]bool
	for i := range s {
		remaining[s[i]]++
	}
	stack := []byte{}
	for i := range s {
		value := s[i]
		remaining[value]--
		if used[value] {
			continue
		}
		for len(stack) > 0 && stack[len(stack)-1] > value && remaining[stack[len(stack)-1]] > 0 {
			used[stack[len(stack)-1]] = false
			stack = stack[:len(stack)-1]
		}
		used[value] = true
		stack = append(stack, value)
	}
	return string(stack)
}

func suffixGreater(a, b []int) bool {
	i := 0
	for i < len(a) && i < len(b) && a[i] == b[i] {
		i++
	}
	return i < len(a) && (i == len(b) || a[i] > b[i])
}
func mergeLargest(a, b []int) []int {
	result := make([]int, 0, len(a)+len(b))
	for len(a) > 0 || len(b) > 0 {
		if suffixGreater(a, b) {
			result = append(result, a[0])
			a = a[1:]
		} else {
			result = append(result, b[0])
			b = b[1:]
		}
	}
	return result
}

func has132(nums []int) bool {
	stack := []int{}
	third := -int(^uint(0)>>1) - 1
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] < third {
			return true
		}
		for len(stack) > 0 && stack[len(stack)-1] < nums[i] {
			third = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, nums[i])
	}
	return false
}

func sortedChunks(arr []int) int {
	maxima := []int{}
	for _, value := range arr {
		if len(maxima) == 0 || maxima[len(maxima)-1] <= value {
			maxima = append(maxima, value)
			continue
		}
		mergedMax := maxima[len(maxima)-1]
		for len(maxima) > 0 && maxima[len(maxima)-1] > value {
			maxima = maxima[:len(maxima)-1]
		}
		maxima = append(maxima, mergedMax)
	}
	return len(maxima)
}

func minimumLeafTreeCost(arr []int) int {
	stack := []int{}
	cost := 0
	for _, value := range arr {
		for len(stack) > 0 && stack[len(stack)-1] <= value {
			middle := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			neighbor := value
			if len(stack) > 0 && stack[len(stack)-1] < neighbor {
				neighbor = stack[len(stack)-1]
			}
			cost += middle * neighbor
		}
		stack = append(stack, value)
	}
	for len(stack) > 1 {
		last := len(stack) - 1
		cost += stack[last] * stack[last-1]
		stack = stack[:last]
	}
	return cost
}

type fleetCar struct{ position, speed int }

func countFleets(target int, position, speed []int) int {
	cars := make([]fleetCar, len(position))
	for i, p := range position {
		cars[i] = fleetCar{p, speed[i]}
	}
	sort.Slice(cars, func(i, j int) bool { return cars[i].position > cars[j].position })
	fleets := 0
	lastArrival := -1.0
	for _, car := range cars {
		arrival := float64(target-car.position) / float64(car.speed)
		if arrival > lastArrival {
			fleets++
			lastArrival = arrival
		}
	}
	return fleets
}

type priceSpan struct{ price, span int }
type StockSpanner struct{ pending []priceSpan }

func collapseSpan(pending []priceSpan, price int) ([]priceSpan, int) {
	span := 1
	for len(pending) > 0 && pending[len(pending)-1].price <= price {
		span += pending[len(pending)-1].span
		pending = pending[:len(pending)-1]
	}
	return append(pending, priceSpan{price, span}), span
}

func widestRamp(nums []int) int {
	candidates := []int{}
	for i, value := range nums {
		if len(candidates) == 0 || value < nums[candidates[len(candidates)-1]] {
			candidates = append(candidates, i)
		}
	}
	best := 0
	for right := len(nums) - 1; right >= 0; right-- {
		for len(candidates) > 0 && nums[candidates[len(candidates)-1]] <= nums[right] {
			left := candidates[len(candidates)-1]
			candidates = candidates[:len(candidates)-1]
			if right-left > best {
				best = right - left
			}
		}
	}
	return best
}
```

[返回总目录](../algorithm_templates/README.md)

[扩展练习：迁移点与边界](#extended-practice) · [保留的原分类笔记](#legacy-notes)

## 目录

- [1. 核心思路](#core)
- [2. 739：每日温度](#lc739)
- [3. 496：下一个更大元素](#lc496)
- [4. 84：柱状图最大矩形](#lc84)
- [5. 907：子数组最小值贡献](#lc907)
- [6. 对比与易错点](#pitfalls)

<a id="core"></a>

## 1. 核心思路

栈保存尚未确定答案的下标。遍历当前元素，遇到能确定栈顶边界的值，就弹栈结算，然后入栈当前下标。

```text
遍历i → while栈顶满足弹出条件：弹出j并结算 → i入栈
```

每个下标最多入栈、出栈一次，所以总时间O(n)。比较符号取决于“严格更大/更小”和重复值归属，不能脱离题目机械背符号。下列适配器调用本页命名局部模板；每份实际 Go 题解也内嵌同一模板，不依赖其他题或仓库本地算法包。

<a id="lc739"></a>

## 2. [739. 每日温度](https://leetcode.com/problems/daily-temperatures/)

[练习](l739/739.go) · [测试](l739/739_test.go)

**题意**：对每一天求还需等几天才有更高温度，不存在为0。

**栈含义**：等待更高温度的下标，温度非递增。当前温度严格更大，才为旧下标确定答案。

```go
func dailyTemperatures(temperatures []int) []int {
	next := nextGreaterPositions(temperatures, false)
	result := make([]int, len(next))
	for i, j := range next {
		if j >= 0 {
			result[i] = j - i
		}
	}
	return result
}
```

**例子**：[30,30,40]，i=2同时解决前两天下标，答案[2,1,0]。相等不是更暖，不能弹。

**复杂度**：时间O(n)，空间O(n)，包括输出和栈。

<a id="lc496"></a>

## 3. [496. 下一个更大元素 I](https://leetcode.com/problems/next-greater-element-i/)

[练习](l496/variants/template/496.go) · [测试](l496/variants/template/496_test.go)

**题意**：nums1是nums2的子集，两数组各自无重复。为nums1每个值查找它在nums2右侧的第一个更大值。

**变化**：仍弹出等待右侧更大值的下标，但结果存为“旧值→下一个更大值”的map，最后按nums1顺序输出。

```go
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	next := nextGreaterPositions(nums2, false)
	values := map[int]int{}
	for i, j := range next {
		values[nums2[i]] = -1
		if j >= 0 {
			values[nums2[i]] = nums2[j]
		}
	}
	result := make([]int, len(nums1))
	for i, x := range nums1 {
		result[i] = values[x]
	}
	return result
}
```

**例子**：nums1=[4,1,2]、nums2=[1,3,4,2]，答案[-1,3,-1]。不是寻找右侧最大值。

**复杂度**：时间O(n+m)，空间O(n+m)，n、m为两数组长度。

<a id="lc84"></a>

## 4. [84. 柱状图中最大的矩形](https://leetcode.com/problems/largest-rectangle-in-histogram/)

[练习](l84/variants/template/84.go) · [测试](l84/variants/template/84_test.go)

**题意**：每根柱宽1，求能组成的最大矩形面积。

**栈含义**：柱高非递减，遇更矮柱时结算旧柱j。弹出后新栈顶left限制左侧扩展，宽为i-left-1。相等柱保留，较早的柱最终覆盖完整宽度。

```go
func largestRectangleArea(heights []int) int { return histogramArea(heights) }
```

**例子**：[2,1,2]，末尾清栈时高度1覆盖宽3，答案3。i=n强制结算，短路条件避免访问heights[n]。

**复杂度**：时间O(n)，空间O(n)。相等保留时left可为等高柱，不能误称每次都是“严格更小左边界”。

<a id="lc907"></a>

## 5. [907. 子数组的最小值之和](https://leetcode.com/problems/sum-of-subarray-minimums/)

[练习](l907/variants/template/907.go) · [测试](l907/variants/template/907_test.go)

**题意**：所有非空子数组的最小值求和，对10^9+7取模。

**贡献法**：为i找左侧最近严格小于L、右侧最近小于等于R。以i为指定最小值的区间数是(i-L)×(R-i)，贡献再乘arr[i]。

```go
func sumSubarrayMins(arr []int) int {
	left, right := monotonicBoundaries(arr, func(previous, current int) bool { return previous >= current })
	const mod int64 = 1_000_000_007
	var result int64
	for i, value := range arr {
		result = (result + int64(value)*int64(i-left[i])*int64(right[i]-i)) % mod
	}
	return int(result)
}
```

**例子**：[2,2]，left=[-1,-1]、right=[1,2]，贡献2和4，总和6。双元素区间归给右侧的2，防止重复。

**复杂度**：时间O(n)，空间O(n)。先转int64再乘；本题约束下乘积可表示。

<a id="pitfalls"></a>

## 6. 对比与易错点

| 题目 | 结算内容 | 相等值处理 |
| --- | --- | --- |
| 739 | 等待天数i-j | 相等不弹 |
| 496 | 下一个更大值 | 题目值唯一 |
| 84 | 高×宽 | 本实现相等保留，由早柱覆盖完整范围 |
| 907 | 最小值×左右选择数 | 左严格、右非严格，统一归属 |

- 栈存下标才能算距离、宽度；不要混淆当前i和被弹出的j。
- 左不存在用-1，右不存在用n，但宽度公式与贡献公式不同。
- 907两侧都严格会重复计数，两侧都非严格会遗漏。
- 遍历方向和弹栈含义改变时，比较符号也要重新推导。
- 窗口最大值的过期处理需要单调队列，不是这个栈骨架。

已有分类：[monotonic_stack](.)。


<a id="legacy-notes"></a>

## 原分类笔记（保留）

# 单调栈题目分类与学习路线

本清单将原有重复题目合并，共 **18 道题**。单调栈不是目的：先明确要寻找的边界、要淘汰的候选或要维护的不变量，再决定栈的方向。

## 一、先理解通用机制

单调栈通常用切片实现，维护栈中元素对应值的单调性。多数题存**下标**，方便计算距离和区间长度；有些题存 `(值, 合并数量)`、字符或分块最大值。

从左到右寻找下一个严格更大元素的基本过程：

```text
遍历下标 i：
    当栈非空且 nums[栈顶] < nums[i]：
        j = 弹出栈顶
        i 是 j 右侧第一个严格更大元素的位置
    将 i 入栈
```

- 栈中下标递增，对应值从栈底到栈顶非递增；相等值在此模板中不弹出。
- 循环里的 while 不意味着 O(n²)：每个下标最多入栈、出栈各一次，因此总时间 O(n)。
- 弹出 j 后的新栈顶，经常表示 j 的另一个方向的边界。
- `>`、`>=`、`<`、`<=` 不能随意替换，尤其涉及重复值和贡献去重时。
- 哨兵或末尾清栈可以统一处理没有找到右边界的元素，但边界必须和状态定义一致。

## 二、题型总览

点击类型或题号跳转到本页对应章节，点击题号旁的 [leetcode] 打开 LeetCode 题目。

| 类型（本页章节） | 题目（本页讲解 / LeetCode） | 主要训练目标 |
| --- | --- | --- |
| [A. 最近更大元素与跨度](#a-最近更大元素与跨度入门) | [496](#496-下一个更大元素-i--next-greater-element-i--简单) [leetcode](https://leetcode.com/problems/next-greater-element-i/)、[503](#503-下一个更大元素-ii--next-greater-element-ii--中等) [leetcode](https://leetcode.com/problems/next-greater-element-ii/)、[901](#901-股票价格跨度--online-stock-span--中等) [leetcode](https://leetcode.com/problems/online-stock-span/) | 理解一个元素何时完成任务、何时能被永久淘汰 |
| [B. 区间边界与贡献计数](#b-区间边界与贡献核心进阶) | [84](#84-柱状图中最大的矩形--largest-rectangle-in-histogram--困难) [leetcode](https://leetcode.com/problems/largest-rectangle-in-histogram/)、[85](#85-最大矩形--maximal-rectangle--困难) [leetcode](https://leetcode.com/problems/maximal-rectangle/)、[907](#907-子数组的最小值之和--sum-of-subarray-minimums--中等) [leetcode](https://leetcode.com/problems/sum-of-subarray-minimums/)、[2104](#2104-子数组范围和--sum-of-subarray-ranges--中等) [leetcode](https://leetcode.com/problems/sum-of-subarray-ranges/) | 弹栈时确定左右边界，处理相等元素 |
| [C. 凹槽与按层计算](#c-凹槽与按层计算) | [42](#42-接雨水--trapping-rain-water--困难) [leetcode](https://leetcode.com/problems/trapping-rain-water/) | 区分凹槽底、左右墙和实际高度 |
| [D. 贪心删选与字典序](#d-贪心删选构造最优子序列) | [402](#402-移掉-k-位数字--remove-k-digits--中等) [leetcode](https://leetcode.com/problems/remove-k-digits/)、[316](#316-去除重复字母--remove-duplicate-letters--中等) [leetcode](https://leetcode.com/problems/remove-duplicate-letters/)、[321](#321-拼接最大数--create-maximum-number--困难) [leetcode](https://leetcode.com/problems/create-maximum-number/) | 在剩余预算/可行性约束下弹栈，构造最优子序列 |
| [E. 特殊不变量与局部合并](#e-特殊不变量局部淘汰合并与模式识别) | [456](#456-132-模式--132-pattern--中等) [leetcode](https://leetcode.com/problems/132-pattern/)、[768](#768-最多能完成排序的块-ii--max-chunks-to-make-sorted-ii--困难) [leetcode](https://leetcode.com/problems/max-chunks-to-make-sorted-ii/)、[1130](#1130-叶值的最小代价生成树--minimum-cost-tree-from-leaf-values--中等) [leetcode](https://leetcode.com/problems/minimum-cost-tree-from-leaf-values/) | 不只盯栈顶；理解弹出值、分块最大值或局部成本 |
| [F. 最远合法配对](#f-最远合法配对先筛左端点) | [962](#962-最大宽度坡--maximum-width-ramp--中等) [leetcode](https://leetcode.com/problems/maximum-width-ramp/) | 维护可能成为最优左端点的候选 |
| [G. 单调队列](#g-单调队列额外处理过期) | [239](#239-滑动窗口最大值--sliding-window-maximum--困难) [leetcode](https://leetcode.com/problems/sliding-window-maximum/) | 同时处理优劣淘汰和窗口过期 |
| [H. 普通栈与排序扫描对照](#h-普通栈与排序扫描不要强行套模板) | [856](#856-括号的分数--score-of-parentheses--中等) [leetcode](https://leetcode.com/problems/score-of-parentheses/)、[853](#853-车队--car-fleet--中等) [leetcode](https://leetcode.com/problems/car-fleet/) | 分清嵌套结构、车队合并与单调栈的差异 |

## A. 最近更大元素与跨度：入门

### [496. 下一个更大元素 I / Next Greater Element I](https://leetcode.com/problems/next-greater-element-i/) · 简单

- **任务**：对查询数组中的每个值，找到它在另一数组中右侧第一个更大的值。
- **状态与操作**：扫描 `nums2`，用非递增栈保存未解决的值或下标。当前值更大时，弹出的元素找到答案；用映射记录答案，再回答 `nums1` 的查询。题目保证值唯一，按值建映射才不会歧义。
- **为什么能弹出**：当前值是被弹出元素遇到的第一个更大值，它的查询任务已经结束。
- **复杂度**：O(len(nums1)+len(nums2)) 时间，O(len(nums2)) 空间。

### [503. 下一个更大元素 II / Next Greater Element II](https://leetcode.com/problems/next-greater-element-ii/) · 中等

- **任务**：496 的循环数组版本，允许从末尾绕回开头。
- **状态与操作**：扫描 `2*n` 个位置，通过 `i%n` 读取值；栈存原数组下标，仅第一轮入栈，第二轮只帮助剩余元素找到答案。找不到的初始化为 -1。
- **关键区别**：需要严格更大，相等值不能作为答案；不要无条件重复入栈，让同一任务出现多份。
- **复杂度**：O(n) 时间、O(n) 空间。先做 496，再做本题。

### [901. 股票价格跨度 / Online Stock Span](https://leetcode.com/problems/online-stock-span/) · 中等

- **任务**：每天返回以当天结束、价格均不高于当天的连续天数。
- **状态与操作**：栈存 `(price, span)`，价格从栈底到顶严格递减。新价格到来，弹出所有 `price <= 当前价格` 的记录，将跨度加到今天的 1 上。
- **为什么能合并**：被弹出记录代表的整段连续历史都不高于当前价格，无须逐天保留。
- **复杂度**：单次最坏 O(n)，每次操作摊还 O(1)，累计空间 O(n)。相等价格也要计入跨度。

## B. 区间边界与贡献：核心进阶

### [84. 柱状图中最大的矩形 / Largest Rectangle in Histogram](https://leetcode.com/problems/largest-rectangle-in-histogram/) · 困难

- **任务**：选择连续柱子，以其中最低柱为高，求最大矩形面积。
- **状态与操作**：可采用严格递增高度栈（遇到 `当前高度 <= 栈顶高度` 时弹出）。弹出下标 j 后，新栈顶 L 是左侧严格更矮边界，当前 i 是右侧第一个小于等于它的边界。
- **计算**：`area = heights[j] * (i-L-1)`，空栈时 L=-1。末尾用虚拟零高度或显式清栈结算。
- **重复值**：上述相等时弹出让后面的相同高度接管更宽区间；也可以保留相等值，但要一致地解释边界。
- **复杂度**：O(n) 时间、O(n) 空间。这里是“取最大面积”，不是“累加每个区间贡献”。

### [85. 最大矩形 / Maximal Rectangle](https://leetcode.com/problems/maximal-rectangle/) · 困难

- **任务**：在二进制矩阵中找全为 1 的最大矩形。
- **转化**：逐行维护直方图高度：当前位置为 1 则高度加一，为 0 则清零。每行调用 84 的直方图算法，表示以本行为底的矩形。
- **关键点**：矩形可以跨多行，但不能跨越同一列上的零；零会重置连续高度。
- **复杂度**：r 行 c 列，O(rc) 时间、O(c) 额外空间。建议做完 84 后立即做本题。

### [907. 子数组的最小值之和 / Sum of Subarray Minimums](https://leetcode.com/problems/sum-of-subarray-minimums/) · 中等

- **任务**：累加每个非空连续子数组的最小值。
- **贡献法**：固定 a[i]，找到左侧最近严格更小的位置 L，右侧最近小于等于的位置 R；它负责的子数组数为 `(i-L)*(R-i)`。
- **贡献**：`a[i] * (i-L) * (R-i)`。可以用两个方向扫描求边界，或在遇到 `当前值 <= 栈顶值` 时弹栈一次结算。
- **为什么两边不对称**：`[2,2]` 中跨越两个 2 的子数组只能归属其中一个。两边都严格会重复计数，两边都非严格会漏计；反向的“一侧严格、一侧非严格”也可行。
- **复杂度**：O(n) 时间、O(n) 空间。Go 用 `int64` 计算乘积，按题意对 `1_000_000_007` 取模。

### [2104. 子数组范围和 / Sum of Subarray Ranges](https://leetcode.com/problems/sum-of-subarray-ranges/) · 中等

- **转化**：每个子数组的范围是 `max-min`，答案等于“所有子数组最大值之和”减去“所有子数组最小值之和”。
- **状态与操作**：最小值贡献使用 907 的递增栈，最大值贡献使用对称的递减栈。两种贡献都要独立处理相等值归属。
- **边界示例**：最小值可用左 `<`、右 `<=`；最大值可用左 `>`、右 `>=`。
- **复杂度**：O(n) 时间、O(n) 空间。返回 `int64`，本题不取模；不要沿用 907 的模运算。

## C. 凹槽与按层计算

### [42. 接雨水 / Trapping Rain Water](https://leetcode.com/problems/trapping-rain-water/) · 困难

- **任务**：计算柱状图可以接住的雨水总量。
- **状态与操作**：维护非递增高度的下标栈。当当前柱高于栈顶，弹出的柱子是凹槽底 bottom；新栈顶是左墙 L，当前 i 是右墙。若弹出后栈空，没有左墙，不能结算。
- **计算**：本层水量为 `(min(height[L],height[i])-height[bottom]) * (i-L-1)`。循环弹栈可逐层填平凹槽。
- **与 84 区别**：84 的弹出柱决定矩形高度；本题的弹出柱是槽底，要用两侧墙的较低高度减去槽底，而不是直接乘槽底高度。
- **复杂度**：O(n) 时间、O(n) 栈空间。双指针还有 O(n) 时间、O(1) 空间解法，适合学完栈后对比。

## D. 贪心删选：构造最优子序列

### [402. 移掉 K 位数字 / Remove K Digits](https://leetcode.com/problems/remove-k-digits/) · 中等

- **任务**：删除恰好 k 位，让剩下的数字最小。
- **状态与操作**：栈存保留的数字。只要还有删除额度，且栈顶大于当前数字，就删除栈顶，让更小数字尽早出现在高位。
- **关键点**：删除额度耗尽后不再保证整体单调；如果遍历结束仍有额度，从末尾删除。最后去掉前导零，空结果返回 `"0"`。
- **复杂度**：O(n) 时间、O(n) 空间。

### [316. 去除重复字母 / Remove Duplicate Letters](https://leetcode.com/problems/remove-duplicate-letters/) · 中等

- **任务**：每个不同字符恰好保留一次，且结果是原串的字典序最小子序列。
- **状态与操作**：栈存选中的字符，另记录已选集合与字符剩余次数。当前字符还未选中时，若栈顶更大且之后还会出现，就可以弹出它。
- **关键约束**：栈顶即使更大，若后面再也不会出现，也不能弹出，否则丢失必选字符。已在栈中的字符跳过，但仍要正确更新剩余次数。
- **复杂度**：O(n) 时间；题目仅有小写字母，辅助状态和栈最多 O(26)，输出也最多 26 个字符。不是简单地排序去重。

### [321. 拼接最大数 / Create Maximum Number](https://leetcode.com/problems/create-maximum-number/) · 困难

- **任务**：从两个数组分别取子序列，保持各自相对顺序，合并得到长度 k 的最大数字序列。
- **三步**：枚举从第一个数组取 x 个；用 402 的反向贪心选取各数组指定长度的最大子序列；再按字典序合并，更新最佳答案。
- **关键点**：合并时不能只比较当前数字。若相同，要比较两个**剩余后缀**，从字典序较大的后缀取下一位，例如 `[6,7]` 与 `[6,0,4]`。
- **复杂度**：设长度为 m、n，可行拆分数为 S（S<=k+1）。朴素后缀比较实现为 O(S*(m+n+k²)) 时间、O(m+n+k) 工作空间；不能把后缀比较当成 O(1)。

## E. 特殊不变量：局部淘汰、合并与模式识别

### [456. 132 模式 / 132 Pattern](https://leetcode.com/problems/132-pattern/) · 中等

- **任务**：寻找 `i<j<k`，使 `nums[i] < nums[k] < nums[j]`。
- **状态与操作**：从右向左扫描。栈维护从栈底到顶非递增的候选值；另用 `middle` 记录已经形成某个“3 > 2”配对的最大“2”，初始为负无穷。
- **弹栈含义**：若当前值大于栈顶，当前值可以充当“3”，被弹出的较小值可以充当“2”；连续弹出的值递增，更新 middle。若之后遇到某值 `< middle`，它就是“1”，三者的下标顺序也因反向扫描而成立。
- **关键点**：答案由**已弹出的值**提供，不只是栈中剩余值；三个大小关系均为严格不等。
- **复杂度**：O(n) 时间、O(n) 空间。

### [768. 最多能完成排序的块 II / Max Chunks To Make Sorted II](https://leetcode.com/problems/max-chunks-to-make-sorted-ii/) · 困难

- **任务**：将数组切成最多的连续块，每块排序后拼接，得到整体有序数组；允许重复值。
- **状态与操作**：栈存各块最大值，非递减。若当前 x >= 栈顶，可以新建块；否则 x 与之前块发生顺序冲突，必须合并。
- **合并方式**：先保存最右块的最大值，再弹出所有最大值大于 x 的块，最后将保存的最大值作为合并块的最大值放回。合并后不能把块最大值错误改成 x。
- **关键点**：相等值允许分块，因此比较是 `>` 而非 `>=`。答案为最终块数。
- **复杂度**：O(n) 时间、O(n) 空间。不要与 [769. 最多能完成排序的块](https://leetcode.com/problems/max-chunks-to-make-sorted/) 混淆：769 限定输入是 0 到 n-1 的排列。

### [1130. 叶值的最小代价生成树 / Minimum Cost Tree From Leaf Values](https://leetcode.com/problems/minimum-cost-tree-from-leaf-values/) · 中等

- **任务**：叶子中序顺序固定，每个非叶节点的值是左右子树最大叶值的乘积，求非叶节点总和最小值。
- **状态与操作**：维护严格递减栈。遇到 `当前值 >= 栈顶值` 时弹出 mid，它已遇到右侧不小于自己的值；左侧候选为弹出后的栈顶。把 mid 与两侧候选中较小者配对，计入 `mid * min(left,right)`；左侧不存在时只能与 right 配对。
- **直觉**：较小叶值应尽早与较小的可用邻接最大值结合，之后它不再影响该子树的最大叶值。较大的值留下来参与更高层组合。
- **清栈**：扫描结束后，剩余递减值按相邻值配对，直到只剩一个。不要漏掉这部分成本。
- **复杂度**：O(n) 时间、O(n) 空间。也可用区间 DP，时间 O(n³)、空间 O(n²)；先理解 DP，再理解此处的贪心消元更稳妥。

## F. 最远合法配对：先筛左端点

### [962. 最大宽度坡 / Maximum Width Ramp](https://leetcode.com/problems/maximum-width-ramp/) · 中等

- **任务**：找 `i<j` 且 `nums[i]<=nums[j]` 的最大 `j-i`。
- **状态与操作**：第一遍从左向右，只把比栈顶值严格更小的下标入栈，维护严格递减的左端点候选。更晚且更大的左端点不可能优于更早且更小的左端点。
- **第二遍**：从右向左扫描 j，当前值足够大时弹出满足条件的左端点 i，并更新宽度。此时 j 已经是该候选能匹配的最靠右位置，之后更小的 j 无法产生更宽答案，所以可以永久弹出。
- **复杂度**：O(n) 时间、O(n) 空间。这不是“寻找最近更大值”，而是利用右侧反向遍历寻找最远合法配对。

## G. 单调队列：额外处理过期

### [239. 滑动窗口最大值 / Sliding Window Maximum](https://leetcode.com/problems/sliding-window-maximum/) · 困难

- **任务**：返回每个长度 k 的窗口最大值。
- **状态与操作**：双端队列存下标，值从队首到队尾递减。新元素到来，从队尾弹出所有不大于它的值：新值更大或相等、下标又更新，旧候选不再有优势。
- **过期处理**：从队首删除所有 `index<=i-k` 的元素。形成完整窗口后，队首对应值就是最大值。
- **为什么不是普通栈**：需要从队尾淘汰劣势候选，同时从队首移除过期候选，必须允许两端操作。
- **复杂度**：O(n) 时间、O(k) 辅助空间（不含输出；Go 切片实现需复用/压缩空间或用环形队列，单纯保留历史前缀可能占 O(n)）。堆也能做，但简单懒删除堆最坏 O(n log n) 时间、O(n) 空间。

## H. 普通栈与排序扫描：不要强行套模板

### [856. 括号的分数 / Score of Parentheses](https://leetcode.com/problems/score-of-parentheses/) · 中等

- **分类**：主要是普通栈/嵌套结构解析，**不要求栈值单调**。
- **规则**：`()` 得 1 分，`AB` 得分相加，`(A)` 得分翻倍。
- **状态与操作**：记录当前层得分；遇 `(` 保存外层得分并清零，遇 `)` 将内层变为 `max(1,2*内层得分)`，再加回外层。
- **复杂度**：O(n) 时间、O(n) 栈空间。也可用深度计数，仅在出现 `()` 时加入 `1 << (depth-1)`，实现 O(1) 额外空间。

### [853. 车队 / Car Fleet](https://leetcode.com/problems/car-fleet/) · 中等

- **分类**：排序 + 到达时间的单调扫描；可以写成栈，但通常不需要显式栈。
- **状态与操作**：按位置从靠近终点到远离终点排序，计算每辆车的独立到达时间 `time=(target-position)/speed`。维护前方最近车队的到达时间 fleetTime。
- **合并规则**：若当前 time <= fleetTime，说明它能在终点前或恰好终点追上前方车队，合并；否则它形成新车队，并更新 fleetTime。扫描中形成的新车队时间严格递增。
- **关键点**：不能超车，追上后只能按前车队速度行驶；不能仅按速度排序；Go 除法前要转 `float64`，否则发生整数截断。
- **复杂度**：O(n log n) 时间；排序副本通常 O(n) 额外空间，排序后的扫描本身只需 O(1) 状态。

## 三、建议学习顺序

1. **最近更大与跨度**：496 → 503 → 901。
2. **边界与重复值**：84 → 85 → 907 → 2104，重点手算 `[2,2]` 的边界归属。
3. **几何与构造**：42 → 402 → 316。
4. **特殊不变量**：962 → 768 → 456 → 1130。
5. **扩展与综合**：239 → 321。
6. **对照学习**：856（普通栈）、853（排序扫描），辨别何时不需要单调栈。

## 四、做题时的检查清单

- 栈保存的是下标、值，还是带数量的聚合信息？
- 栈底到栈顶究竟是严格递增、非递减、严格递减还是非递增？
- 弹栈代表“答案已确定”“候选被支配”“需要合并”，还是“修复字典序”？
- 弹出元素后，新栈顶有什么含义？
- 相等值应保留还是弹出？贡献类有没有重复计数？
- 没找到边界时，用 -1、n、哨兵还是末尾清栈？
- 答案是最大值、总贡献、配对费用还是子序列？乘法是否需要 `int64`？
- 复杂度是否按所有入栈/出栈次数合计，而不是只看 while 的嵌套？

## 五、参考资料

- [LeetCode 单调栈题目](https://leetcode.com/tag/monotonic-stack/)
- [907：单调栈逐步解释](https://leetcode.com/problems/sum-of-subarray-minimums/discuss/178876/stack-solution-with-very-detailed-explanation-step-by-step)
- [42：O(n) 时间、O(1) 空间双指针解释](https://leetcode.com/problems/trapping-rain-water/discuss/153992/Java-O(n)-time-and-O(1)-space-(with-explanations).)
- [456：从右向左扫描及弹出值的意义](https://leetcode.com/problems/132-pattern/discuss/94071/Single-pass-C%2B%2B-O(n)-space-and-time-solution-(8-lines)-with-detailed-explanation.)

原清单中的 901、2104 已归入对应学习分类；本文是学习索引，不代表相关 Go 实现已完成。

<a id="extended-practice"></a>

## 扩展练习：迁移点与边界

新增题目只补充学习路线与解法总结，不自动生成未请求的实现或测试。

| 题目（原题） | 如何迁移当前模板 | 必须确认的条件 / 边界 |
| --- | --- | --- |
| [503. 下一个更大元素II](https://leetcode.com/problems/next-greater-element-ii/) | 循环数组逻辑遍历两遍，取模访问，只为必要下标入栈 | 严格更大，相等不能作为答案 |
| [901. 股票价格跨度](https://leetcode.com/problems/online-stock-span/) | 流式栈存(price,span)，弹出不大于当前价格者并合并跨度 | 本题相等需要计入跨度，区别于739 |
| [2104. 子数组范围和](https://leetcode.com/problems/sum-of-subarray-ranges/) | 分别计算最大值贡献与最小值贡献再相减 | 两组边界都要统一重复值归属；结果int64 |
| [1856. 子数组最小乘积的最大值](https://leetcode.com/problems/maximum-subarray-min-product/) | 单调栈找最小值作用范围，前缀和算范围总和 | 正数保证扩展到边界最优；先比较真实乘积再取模 |
| [402. 移掉K位数字](https://leetcode.com/problems/remove-k-digits/) | 当前数字更小时弹掉更大前位，维护最小字典序子序列 | 仍有删除预算则删尾部，去前导0，空结果返回0 |
| [962. 最大宽度坡](https://leetcode.com/problems/maximum-width-ramp/) | 先筛严格递减的候选左端点，右端从右往左弹栈结算 | 不是最近更大元素；匹配后最远右端已经确定 |

**复习方法**：先独立写出状态、合法性或松弛条件，再指出相较本篇基础题改变的那一行；若前提不满足，主动换算法，不为统一模板牺牲正确性。

## 目录归并：单调性、结算方向与重复值归属

重复的 496、84、907 模板练习已分别并入 [l496/variants/template](l496/variants/template)、[l84/variants/template](l84/variants/template)、[l907/variants/template](l907/variants/template)，原测试和目录来源保留，实际源码现按本页局部模板实现。503 的排列分类练习保存在 [l503/variants/from_permutation](l503/variants/from_permutation)。不同包保留各自函数名与测试，各变体现使用同名模板独立实现，保留各自原测试。

| 子模型 | 栈保存什么 / 什么时候可结算 | 题目入口 |
| --- | --- | --- |
| 下一个严格更大 | 尚未找到答案的下标；当前值严格大于栈顶时弹出 | [496](l496)、[503](l503)、[739](l739) |
| 矩形边界 | 候选高度；遇更低高度时结算 `height*(right-left-1)` | [84](l84)、[85](l85) |
| 最值贡献 | 用一严格一非严格边界给相等值唯一归属 | [907](l907)、[2104](l2104) |
| 凹槽水层 | 当前更高柱子确定右边界；弹底后还须有左边界 | [42](l42) |
| 带预算的贪心栈 | 较优新字符可以淘汰较差前缀，但必须仍能完成答案 | [402](l402)、[316](l316)、[321](l321) |
| 非最近边界的候选淘汰 | 候选左端点、分块最大值、跨度等专有状态 | [962](l962)、[768](l768)、[901](l901)、[456](l456)、[1130](l1130)、[853](l853) |

最近严格更大位置的可复用骨架如下；它不负责矩形清栈或贡献归属：

```go
func nextGreaterIndices(nums []int) []int { return nextGreaterPositions(nums, false) }
```

496 把位置结果转换为对应值；739 转换为等待距离；503 为循环数组处理第二遍，但不能把自己或相等值当作严格更大值。84/907 即使使用相似出栈循环，结算内容与相等值策略也不同。

两处重新分类：

- 原 856 已归入 [stack/l856](../stack/l856)，原练习在其 `variants/from_monotonic_stack`。嵌套层的分数累加没有单调值不变量。
- 原 239 已归入 [sliding_window/l239](../sliding_window/l239)。它需要淘汰过期队首，是单调队列，不能只用栈顶操作表达窗口失效。

当前实际 Go 文件已按本页局部模板实现，独立变体仍运行各自原测试；不以排除练习文件制造通过。迁移清单见 [reorganization.json](../two_pointers/reorganization.json)。


## 逐题模板适配器补全

下列函数使用本页已声明的同名子模板；不是将不同不变量强行改为同一弹栈或移动条件。

### 42

```go
func trap(height []int) int { return trappedLayers(height) }
```

### 85

```go
func maximalRectangle(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}
	heights := make([]int, len(matrix[0]))
	best := 0
	for _, row := range matrix {
		for j, value := range row {
			if value == '1' {
				heights[j]++
			} else {
				heights[j] = 0
			}
		}
		if area := histogramArea(heights); area > best {
			best = area
		}
	}
	return best
}
```

### 316

```go
func removeDuplicateLetters(s string) string { return distinctSubsequence(s) }
```

### 321

```go
func maxNumber(nums1 []int, nums2 []int, k int) []int {
	best := []int{}
	for take := 0; take <= k; take++ {
		if take > len(nums1) || k-take > len(nums2) {
			continue
		}
		a := greedySubsequence(nums1, take, func(current, previous int) bool { return current > previous })
		b := greedySubsequence(nums2, k-take, func(current, previous int) bool { return current > previous })
		candidate := mergeLargest(a, b)
		if suffixGreater(candidate, best) {
			best = candidate
		}
	}
	return best
}
```

### 402

```go
func removeKdigits(num string, k int) string {
	digits := greedySubsequence([]byte(num), len(num)-k, func(current, previous byte) bool { return current < previous })
	start := 0
	for start < len(digits) && digits[start] == '0' {
		start++
	}
	if start == len(digits) {
		return "0"
	}
	return string(digits[start:])
}
```

### 456

```go
func find132pattern(nums []int) bool { return has132(nums) }
```

### 503

```go
func nextGreaterElements(nums []int) []int {
	next := nextGreaterPositions(nums, true)
	result := make([]int, len(next))
	for i, j := range next {
		result[i] = -1
		if j >= 0 {
			result[i] = nums[j]
		}
	}
	return result
}
```

### 768

```go
func maxChunksToSorted(arr []int) int { return sortedChunks(arr) }
```

### 853

```go
func carFleet(target int, position []int, speed []int) int {
	return countFleets(target, position, speed)
}
```

### 901

```go
func Constructor() StockSpanner { return StockSpanner{} }
func (s *StockSpanner) Next(price int) int {
	var span int
	s.pending, span = collapseSpan(s.pending, price)
	return span
}
```

### 962

```go
func maxWidthRamp(nums []int) int { return widestRamp(nums) }
```

### 1130

```go
func mctFromLeafValues(arr []int) int { return minimumLeafTreeCost(arr) }
```

### 2104

```go
func subArrayRanges(nums []int) int64 {
	minLeft, minRight := monotonicBoundaries(nums, func(previous, current int) bool { return previous >= current })
	maxLeft, maxRight := monotonicBoundaries(nums, func(previous, current int) bool { return previous <= current })
	var result int64
	for i, value := range nums {
		minima := int64(i-minLeft[i]) * int64(minRight[i]-i)
		maxima := int64(i-maxLeft[i]) * int64(maxRight[i]-i)
		result += int64(value) * (maxima - minima)
	}
	return result
}
```

### 1856

```go
func maxSumMinProduct(nums []int) int {
	prefix := make([]int64, len(nums)+1)
	for i, value := range nums {
		prefix[i+1] = prefix[i] + int64(value)
	}
	left, right := monotonicBoundaries(nums, func(previous, current int) bool { return previous >= current })
	var best int64
	for i, value := range nums {
		product := int64(value) * (prefix[right[i]] - prefix[left[i]+1])
		if product > best {
			best = product
		}
	}
	return int(best % 1_000_000_007)
}
```

### 2281

```go
func totalStrength(strength []int) int {
	const mod int64 = 1_000_000_007
	n := len(strength)
	prefix := make([]int64, n+1)
	prefixOfPrefix := make([]int64, n+2)
	for i, value := range strength {
		prefix[i+1] = (prefix[i] + int64(value)) % mod
	}
	for i, value := range prefix {
		prefixOfPrefix[i+1] = (prefixOfPrefix[i] + value) % mod
	}
	left, right := monotonicBoundaries(strength, func(previous, current int) bool { return previous >= current })
	var result int64
	for i, value := range strength {
		sumRight := (prefixOfPrefix[right[i]+1] - prefixOfPrefix[i+1] + mod) % mod
		sumLeft := (prefixOfPrefix[i+1] - prefixOfPrefix[left[i]+1] + mod) % mod
		weighted := (int64(i-left[i])*sumRight - int64(right[i]-i)*sumLeft) % mod
		if weighted < 0 {
			weighted += mod
		}
		result = (result + int64(value)*weighted) % mod
	}
	return int(result)
}
```

1856 依赖正数，边界内尽量扩展可最大化区间和，先比较 int64 真实乘积再取模。2281 使用同一最小值归属边界，再用前缀和的前缀和求左右端点选择的总区间和；减法结果要规范到非负模值，不与 907 的单纯区间个数贡献混淆。
