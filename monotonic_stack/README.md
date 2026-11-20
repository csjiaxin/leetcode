# 单调栈：识别题型、维护候选与结算边界

[返回总目录](../algorithm_templates/README.md) · [贪心专题](../greedy/readme.md)

单调栈不只是“找下一个更大元素”。它维护的是**仍有价值的候选**：新元素到来时，从栈顶连续地为旧元素确定答案、淘汰劣势候选，或结算已经形成的区间结构。

## 目录

1. [遇到新题，如何判断可以用单调栈？](#recognize)
2. [通用机制、复杂度与比较符号](#core)
3. [A. 最近更大 / 更小元素与跨度](#nearest)
4. [B. 左右边界与最大扩展区间](#boundaries)
5. [C. 子数组最值的贡献统计](#contribution)
6. [D. 凹槽与按层结算](#water)
7. [E. 贪心删选与字典序](#subsequence)
8. [F. 特殊候选、合并与模式识别](#special)
9. [G. 与单调队列、普通栈、排序扫描的区别](#alternatives)
10. [检查清单与学习路线](#pitfalls)
11. [本地实现与变体说明](#implementation)

本文将原先分散的基础题、分类笔记、扩展练习与适配器按题型合并，每个主题只保留一套主要解释。题号链接指向 LeetCode，“源码”链接指向仓库；扩展题单不代表本地已有实现。

<a id="recognize"></a>

## 1. 遇到新题，如何判断可以用单调栈？

### 1.1 五个识别问题

| 问题 | 典型信号 / 需要证明的事情 |
| --- | --- |
| 是否需要比较当前位置与某一侧元素？ | 最近更高、左侧第一个更小、遇到阻挡前能延伸多远 |
| 暴力是否对每个位置向左 / 右扫描，直到遇到某个边界？ | 这种 O(n²) 搜索经常可以合并成一次扫描 |
| 新元素能否让旧候选永久退出？ | 答案已经确定、被新候选支配、形成可结算结构，或在约束允许下撤销较差前缀 |
| 为什么只处理栈顶就够了？ | 当前应处理的候选必须集中在栈顶连续一段，而不是要从中间随机删除 |
| 相等元素如何处理？ | 严格与非严格关系、边界阻挡、重复值贡献归属，都必须先定义 |

**最关键的判据**：能否证明旧候选退出后不再需要单独保留？如果只弹栈却丢失未来必需的信息，就不能这样做；有时需要先把它的跨度、块最大值等信息合并到新记录中。

### 1.2 并非所有“更大 / 更小”都适用

| 需求 | 优先考虑 |
| --- | --- |
| 每个位置一侧最近的更大 / 更小值 | 单调栈 |
| 固定滑动窗口最大值 | 单调队列，需要从队首删除过期元素 |
| 任意区间最大值查询 | 线段树、稀疏表等，取决于是否修改 |
| 右侧比当前小的元素数量 | 树状数组、归并计数等，不只是找最近边界 |
| 从动态候选中任意取最小 / 最大值 | 堆、有序集合；不一定具备栈顶连续淘汰结构 |
| 括号、嵌套表达式 | 普通栈，不要求值单调 |

<a id="core"></a>

## 2. 通用机制、复杂度与比较符号

先确定栈的含义，再决定单调方向；不要先背“递增栈”这个名字。

```text
遍历当前元素 i：
    while 栈非空，且 i 让栈顶应该退出：
        j = 弹出栈顶
        为 j 结算答案，或把 j 的信息合并到新状态
    必要时读取剩余栈顶，为当前 i 计算答案
    若满足入栈条件，将 i 或聚合记录入栈
根据状态定义处理剩余候选
```

- **存下标**：需要距离、宽度、左右边界；比较值时通过数组取值。
- **存值 / 字符**：只关心数值关系或直接构造答案。
- **存 `(值, 数量)` 等记录**：一次弹出代表整段历史，例如 901。
- **摊还 O(n)**：每个元素最多入栈、出栈一次，while 的总执行次数是线性的。排序、后缀比较、多轮枚举等额外操作应另算，不能把整个算法都一概称为 O(n)。

### 2.1 比较符号表：只针对“从左到右，弹栈结算旧位置的右侧答案”

| 要找右侧第一个…… | 弹栈条件 | 栈底到栈顶的值 |
| --- | --- | --- |
| 严格更大 `>` | `nums[top] < nums[i]` | 非递增 |
| 大于等于 `>=` | `nums[top] <= nums[i]` | 严格递减 |
| 严格更小 `<` | `nums[top] > nums[i]` | 非递减 |
| 小于等于 `<=` | `nums[top] >= nums[i]` | 严格递增 |

换成反向扫描或“弹栈后读取栈顶回答当前元素”，需要重新推导，不可机械照搬。

### 2.2 什么时候结算谁？

- **弹栈时回答旧位置 j**：右侧第一个更大值、矩形、贡献统计。
- **弹栈后读栈顶回答当前 i**：左侧最近更大 / 更小值。
- **弹栈直接修改候选答案**：删除数字、构造最优子序列。

末尾也不一定要清栈：最近更大值中未解决的元素保持 -1；矩形需要最终右边界；贪心子序列可能需要截断尾部。清栈策略取决于任务。

本文 Go 模板只使用标准库；组合下面的 Go 代码块需添加 `package main`，不需要额外 import。泛型 helper 需要 Go 1.18+。乘积与累加按各题约束使用 `int64`，返回 `int` 的示例仍要求结果能由平台 int 表示。

<a id="nearest"></a>

## 3. A. 最近更大 / 更小元素与跨度

### 通用模板：右侧第一个严格更大位置

栈存等待答案的下标，对应值非递增。当前值能弹出 j，说明 j 之前一直没遇到更大值，因此当前就是第一个。

```go
// 无答案返回 -1；循环数组只在第一轮入栈，第二轮解决剩余候选。
func nextGreaterPositions(nums []int, circular bool) []int {
    n := len(nums)
    next := make([]int, n)
    for i := range next { next[i] = -1 }
    stack := make([]int, 0, n)
    rounds := n
    if circular { rounds *= 2 }
    for step := 0; step < rounds; step++ {
        i := step % n // n=0 时循环不会执行。
        for len(stack) > 0 && nums[stack[len(stack)-1]] < nums[i] {
            j := stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            next[j] = i
        }
        if step < n { stack = append(stack, i) }
    }
    return next
}
```

<a id="lc739"></a>

### 739. 每日温度：位置转距离

[原题](https://leetcode.com/problems/daily-temperatures/) · [源码](l739/739.go) · [测试](l739/739_test.go)

```go
func dailyTemperatures(temperatures []int) []int {
    next := nextGreaterPositions(temperatures, false)
    result := make([]int, len(next))
    for i, j := range next {
        if j >= 0 { result[i] = j-i }
    }
    return result
}
```

例如 `[30,30,40] → [2,1,0]`。相等不算更暖，不弹栈。时间、空间均 O(n)。

### 同类题的迁移

| 题目 | 怎样修改 / 关键边界 |
| --- | --- |
| <a id="lc496"></a>[496. 下一个更大元素 I](https://leetcode.com/problems/next-greater-element-i/) · [源码](l496/496.go) | 对 nums2 调 helper，再建“旧值→下一个更大值”映射回答 nums1；题目保证值唯一。时间 O(n+m)，不是找右侧最大值 |
| [503. 下一个更大元素 II](https://leetcode.com/problems/next-greater-element-ii/) · [源码](l503/503.go) | `circular=true`，将结果下标转成数值；相等值和自己都不能作为严格更大值。时间 O(n) |
| [1475. 商品折扣后的最终价格](https://leetcode.com/problems/final-prices-with-a-special-discount-in-a-shop/) | 找右侧第一个小于等于当前值，弹栈条件改成 `nums[top] >= nums[i]`，用找到的值作折扣 |
| [901. 股票价格跨度](https://leetcode.com/problems/online-stock-span/) · [源码](l901/901.go) | 找左侧最近严格更大值；将所有不大于当前价的历史跨度合并，见下方流式模板 |

```go
type priceSpan struct { price, span int }
type StockSpanner struct { pending []priceSpan }
func Constructor() StockSpanner { return StockSpanner{} }
func (s *StockSpanner) Next(price int) int {
    span := 1
    for len(s.pending) > 0 && s.pending[len(s.pending)-1].price <= price {
        span += s.pending[len(s.pending)-1].span
        s.pending = s.pending[:len(s.pending)-1]
    }
    s.pending = append(s.pending, priceSpan{price, span})
    return span
}
```

901 的价格从栈底到顶严格递减，相等也需要合并。单次最坏 O(n)，每次调用摊还 O(1)，累计空间 O(n)。

<a id="boundaries"></a>

## 4. B. 左右边界与最大扩展区间

信号：**以某个元素作为区间的最低高度 / 限制，向两边最多扩展多远？** 由逐个位置向两边搜索，转为统一计算边界。

### 通用边界 helper 与相等值约定

下方 helper 在弹栈时确定旧位置的右边界，当前元素入栈前读取剩余栈顶作为左边界。不存在时左为 -1、右为 n。

| pop(previous,current) | 得到的左边界 | 得到的右边界 |
| --- | --- | --- |
| `previous > current` | 最近小于等于 | 最近严格小于 |
| `previous >= current` | 最近严格小于 | 最近小于等于 |
| `previous < current` | 最近大于等于 | 最近严格大于 |
| `previous <= current` | 最近严格大于 | 最近大于等于 |

这个单遍 helper 会自然得到一侧严格、一侧非严格的边界。**它并不总是返回左右两侧严格更小的位置。**

```go
func monotonicBoundaries(nums []int, pop func(previous, current int) bool) ([]int, []int) {
    n := len(nums)
    left, right := make([]int, n), make([]int, n)
    stack := make([]int, 0, n)
    for i, x := range nums {
        right[i] = n
        for len(stack) > 0 && pop(nums[stack[len(stack)-1]], x) {
            j := stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            right[j] = i
        }
        left[i] = -1
        if len(stack) > 0 { left[i] = stack[len(stack)-1] }
        stack = append(stack, i)
    }
    return left, right
}
```

<a id="lc84"></a>

### 84. 柱状图中最大的矩形

[原题](https://leetcode.com/problems/largest-rectangle-in-histogram/) · [源码](l84/84.go) · [测试](l84/84_test.go)

枚举每根柱作为高度，计算它负责的最大宽度。本文统一采用**相等保留**：左边界小于等于、右边界严格小于，较早的同高柱最终负责完整宽度。虽然部分同高柱被左侧等高柱限制，但不会漏掉最大面积。

```go
func histogramArea(heights []int) int {
    left, right := monotonicBoundaries(heights, func(a, b int) bool { return a > b })
    var best int64
    for i, h := range heights {
        area := int64(h)*int64(right[i]-left[i]-1)
        if area > best { best = area }
    }
    return int(best)
}
func largestRectangleArea(heights []int) int { return histogramArea(heights) }
```

例：`[2,1,5,6,2,3]` 中高度 5 覆盖 `[5,6]`，面积 10；`[2,2]` 中最左柱负责宽 2，面积 4。右边界默认 n，循环后直接计算，无须再引入 `i=n` 虚拟柱或显式清栈。时间、空间均 O(n)。

### 85. 最大矩形：逐行转直方图

[原题](https://leetcode.com/problems/maximal-rectangle/) · [源码](l85/85.go)

每列记录到当前行连续的 1 的数量，遇 0 清零。每行调用 84，枚举以本行为底的矩形。矩阵按题意为矩形，元素为字节 `'0'`、`'1'`。

```go
func maximalRectangle(matrix [][]byte) int {
    if len(matrix) == 0 { return 0 }
    heights := make([]int, len(matrix[0]))
    best := 0
    for _, row := range matrix {
        for j, value := range row {
            if value == '1' { heights[j]++ } else { heights[j] = 0 }
        }
        if area := histogramArea(heights); area > best { best = area }
    }
    return best
}
```

r 行 c 列，时间 O(rc)，额外空间 O(c)。

### 边界扩展题

- [1793. 好子数组的最大分数](https://leetcode.com/problems/maximum-score-of-a-good-subarray/)：以每个数为最小值，检查负责区间是否包含指定位置 k，再计算最小值乘长度；也可使用双指针。
- [1856. 子数组最小乘积的最大值](https://leetcode.com/problems/maximum-subarray-min-product/)：最小值作用范围＋前缀和。**正数条件**使范围扩展越大，区间和越大；先比较真实乘积，最后才取模。重复最小值中仍有一个代表能覆盖最优区间。

```go
func maxSumMinProduct(nums []int) int {
    prefix := make([]int64, len(nums)+1)
    for i, x := range nums { prefix[i+1] = prefix[i]+int64(x) }
    left, right := monotonicBoundaries(nums, func(a, b int) bool { return a >= b })
    var best int64
    for i, x := range nums {
        product := int64(x)*(prefix[right[i]]-prefix[left[i]+1])
        if product > best { best = product }
    }
    return int(best % 1_000_000_007)
}
```

1856 时间、空间均 O(n)；真实乘积在原题约束下可由 int64 表示。若出现负数，“扩展到整个边界最优”的推理不再成立。[历史练习](l1856/legacy) 不等同于可直接提交的当前实现。

<a id="contribution"></a>

## 5. C. 子数组最值的贡献统计

信号：**所有子数组的最小值 / 最大值之和**。不要枚举 O(n²) 个子数组，而是枚举某个元素负责多少个子数组。

若 i 的负责边界为 L、R，则：

```text
左端点选择：L+1 ... i，共 i-L 种
右端点选择：i ... R-1，共 R-i 种
贡献次数 = (i-L)*(R-i)
贡献值   = nums[i]*(i-L)*(R-i)
```

### 重复值：为每个子数组指定唯一的最值代表

本文最小值统一用**左侧严格小于、右侧小于等于**（helper 的 `pop >=`），让相等最小值中最右的那个负责。`[2,2]`：左边 2 贡献 2，右边 2 贡献 4，合计 6；双元素子数组只算一次。

反过来的“左非严格、右严格”也成立，只要统一。对最近阻挡边界而言，**两侧都严格会重复计数，两侧都非严格会漏计**。

<a id="lc907"></a>

### 907. 子数组的最小值之和

[原题](https://leetcode.com/problems/sum-of-subarray-minimums/) · [源码](l907/907.go) · [测试](l907/907_test.go)

```go
func sumSubarrayMins(arr []int) int {
    left, right := monotonicBoundaries(arr, func(a, b int) bool { return a >= b })
    const mod int64 = 1_000_000_007
    var result int64
    for i, x := range arr {
        result = (result+int64(x)*int64(i-left[i])*int64(right[i]-i))%mod
    }
    return int(result)
}
```

时间、空间均 O(n)。先转 int64 再相乘，乘积在本题约束下可表示。

### 2104. 子数组范围和

[原题](https://leetcode.com/problems/sum-of-subarray-ranges/) · [源码](l2104/2104.go)

`所有子数组(max-min) = 最大值贡献和 - 最小值贡献和`。最大值用对称的左严格大于、右大于等于。返回 int64，不取模。

```go
func subArrayRanges(nums []int) int64 {
    minL, minR := monotonicBoundaries(nums, func(a, b int) bool { return a >= b })
    maxL, maxR := monotonicBoundaries(nums, func(a, b int) bool { return a <= b })
    var result int64
    for i, x := range nums {
        minima := int64(i-minL[i])*int64(minR[i]-i)
        maxima := int64(i-maxL[i])*int64(maxR[i]-i)
        result += int64(x)*(maxima-minima)
    }
    return result
}
```

时间、空间均 O(n)。

### 2281. 巫师的总力量和：贡献还要乘区间和

[原题](https://leetcode.com/problems/sum-of-total-strength-of-wizards/) · [历史练习](l2281/legacy)

每个子数组贡献为 `最小值 × 区间和`，边界与 907 相同，但不能只乘区间个数。令 `P[t]` 为前 t 项和，`Q[t]` 为前 t 个 P 的和：

```text
指定最小值 i，边界 L、R：
Σ(l=L+1..i) Σ(r=i..R-1) (P[r+1]-P[l])
= (i-L)*(Q[R+1]-Q[i+1]) - (R-i)*(Q[i+1]-Q[L+1])
```

```go
func totalStrength(strength []int) int {
    const mod int64 = 1_000_000_007
    n := len(strength)
    prefix := make([]int64, n+1)
    sums := make([]int64, n+2)
    for i, x := range strength { prefix[i+1] = (prefix[i]+int64(x))%mod }
    for i, x := range prefix { sums[i+1] = (sums[i]+x)%mod }
    left, right := monotonicBoundaries(strength, func(a, b int) bool { return a >= b })
    var result int64
    for i, x := range strength {
        sumRight := (sums[right[i]+1]-sums[i+1]+mod)%mod
        sumLeft := (sums[i+1]-sums[left[i]+1]+mod)%mod
        weighted := (int64(i-left[i])*sumRight-int64(right[i]-i)*sumLeft)%mod
        if weighted < 0 { weighted += mod }
        result = (result+int64(x)*weighted)%mod
    }
    return int(result)
}
```

时间、空间均 O(n)。减法需规范到非负模值；这是“前缀和的前缀和”，不要与 907 单纯计算端点选择数量混淆。

<a id="water"></a>

## 6. D. 凹槽与按层结算

### 42. 接雨水

[原题](https://leetcode.com/problems/trapping-rain-water/) · [源码](l42/42.go)

维护非递增高度栈。当前柱更高时，弹出的是**槽底**；弹出后的新栈顶是左墙，当前是右墙。有两侧墙才能结算一层水。

```text
本层高度 = min(左墙高度, 右墙高度) - 槽底高度
本层宽度 = right-left-1
本层水量 = 高度*宽度
```

```go
func trap(height []int) int {
    stack := []int{}
    water := 0
    for right, h := range height {
        for len(stack) > 0 && height[stack[len(stack)-1]] < h {
            bottom := stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            if len(stack) == 0 { break }
            left := stack[len(stack)-1]
            level := height[left]
            if h < level { level = h }
            water += (right-left-1)*(level-height[bottom])
        }
        stack = append(stack, right)
    }
    return water
}
```

时间 O(n)，栈空间 O(n)。与 84 不同，弹出的高度不是最终水层高度；例如 `[3,0,2]` 的水深为 2。未找到右墙的候选不结算。另有双指针 O(n) 时间、O(1) 空间解法。

<a id="subsequence"></a>

## 7. E. 贪心删选与字典序

信号：**保持相对顺序，删除若干元素，构造最小 / 最大子序列**。字典序由第一个不同位置决定，因此在能完成任务的前提下，让更优元素尽量靠前。

### 固定长度子序列模板

```go
// 保留恰好 k 个元素，要求 0 <= k <= len(values)。
// better(current,previous) 表示当前值更适合出现在较前位置。
func greedySubsequence[T any](values []T, k int, better func(current, previous T) bool) []T {
    drop := len(values)-k
    stack := make([]T, 0, len(values))
    for _, value := range values {
        for drop > 0 && len(stack) > 0 && better(value, stack[len(stack)-1]) {
            stack = stack[:len(stack)-1]
            drop--
        }
        stack = append(stack, value)
    }
    return stack[:k] // 还没用完的删除额度，从尾部扣除。
}
```

**不是一味弹掉较差栈顶**，还必须保证删除预算或剩余元素数量允许。预算耗尽后不再保证栈中整体单调。这个 helper 时间、空间均 O(n)，返回值与栈底层数组共享内存。

### 402. 移掉 K 位数字

[原题](https://leetcode.com/problems/remove-k-digits/) · [源码](l402/402.go)

```go
func removeKdigits(num string, k int) string {
    digits := greedySubsequence([]byte(num), len(num)-k, func(a, b byte) bool { return a < b })
    start := 0
    for start < len(digits) && digits[start] == '0' { start++ }
    if start == len(digits) { return "0" }
    return string(digits[start:])
}
```

`1432219, k=3 → 1219`：依次撤销阻挡较小前缀的 4、3、2。遍历后处理剩余额度，再去前导零；空结果返回 `"0"`。时间、空间均 O(n)。

### 316 / 1081：每个不同字符恰好一次

[316 原题](https://leetcode.com/problems/remove-duplicate-letters/) · [1081 原题](https://leetcode.com/problems/smallest-subsequence-of-distinct-characters/) · [316 源码](l316/316.go)

不同于固定删除预算：弹掉较大字符前，必须确认它之后还会出现，否则会丢失必选字符。另用 used 跳过已选字符，remaining 记录剩余次数。

```go
// 适用于原题的小写 ASCII 字符；按字节处理，不是通用 Unicode 字符模板。
func removeDuplicateLetters(s string) string {
    var remaining [256]int
    var used [256]bool
    for i := 0; i < len(s); i++ { remaining[s[i]]++ }
    stack := []byte{}
    for i := 0; i < len(s); i++ {
        x := s[i]
        remaining[x]--
        if used[x] { continue }
        for len(stack) > 0 && stack[len(stack)-1] > x && remaining[stack[len(stack)-1]] > 0 {
            used[stack[len(stack)-1]] = false
            stack = stack[:len(stack)-1]
        }
        used[x] = true
        stack = append(stack, x)
    }
    return string(stack)
}
```

时间 O(n)，原题字符集下额外空间 O(26)。例如 `cbacdcbc → acdb`，不能简单排序去重。

### 其他构造题

- [1673. 找出最具竞争力的子序列](https://leetcode.com/problems/find-the-most-competitive-subsequence/)：直接调用 `greedySubsequence(nums, k, func(a,b int) bool { return a < b })`。
- [321. 拼接最大数](https://leetcode.com/problems/create-maximum-number/) · [源码](l321/321.go)：枚举两个数组各取多少个，各自调用最大子序列模板，再合并。**合并时比较剩余后缀，不只是首元素**；例如 `[6,7]` 与 `[6,0,4]`，应先取前者的 6。

```go
func suffixGreater(a, b []int) bool {
    i := 0
    for i < len(a) && i < len(b) && a[i] == b[i] { i++ }
    return i < len(a) && (i == len(b) || a[i] > b[i])
}
func mergeLargest(a, b []int) []int {
    result := make([]int, 0, len(a)+len(b))
    for len(a) > 0 || len(b) > 0 {
        if suffixGreater(a, b) {
            result = append(result, a[0]); a = a[1:]
        } else {
            result = append(result, b[0]); b = b[1:]
        }
    }
    return result
}
func maxNumber(nums1, nums2 []int, k int) []int {
    best := []int{}
    for take := 0; take <= k; take++ {
        if take > len(nums1) || k-take > len(nums2) { continue }
        a := greedySubsequence(nums1, take, func(a, b int) bool { return a > b })
        b := greedySubsequence(nums2, k-take, func(a, b int) bool { return a > b })
        candidate := mergeLargest(a, b)
        if suffixGreater(candidate, best) { best = candidate }
    }
    return best
}
```

设数组长 m、n，可行拆分数 S≤k+1，朴素后缀比较使总时间为 O(S·(m+n+k²))，工作空间 O(m+n+k)，不是简单的 O(m+n)。更广泛的证明与选型见 [贪心分类指南](../greedy/readme.md#greedy-stack)。

<a id="special"></a>

## 8. F. 特殊候选、合并与模式识别

这些题共享“候选从栈顶连续退出”的机制，但不能直接套最近更大值的结算公式。

### 962. 最大宽度坡：最远匹配，不是最近匹配

[原题](https://leetcode.com/problems/maximum-width-ramp/) · [源码](l962/962.go)

找 `i<j 且 nums[i]<=nums[j]` 的最大 `j-i`。先从左到右，只保留严格递减的左端点候选：更晚且值更大的左端点，被更早且值更小的候选支配。再从右向左扫描右端点；一旦匹配，就已得到该左端点最远的合法右端，可以永久弹出。

```go
func maxWidthRamp(nums []int) int {
    candidates := []int{}
    for i, x := range nums {
        if len(candidates) == 0 || x < nums[candidates[len(candidates)-1]] {
            candidates = append(candidates, i)
        }
    }
    best := 0
    for right := len(nums)-1; right >= 0; right-- {
        for len(candidates) > 0 && nums[candidates[len(candidates)-1]] <= nums[right] {
            left := candidates[len(candidates)-1]
            candidates = candidates[:len(candidates)-1]
            if right-left > best { best = right-left }
        }
    }
    return best
}
```

时间、空间均 O(n)，没有合法正宽度时返回 0。

### 456. 132 模式：答案还需要已弹出的状态

[原题](https://leetcode.com/problems/132-pattern/) · [源码](l456/456.go)

从右向左扫描，栈值非递增。当前大值作为“3”，弹出的较小值作为“2”；另维护已形成合法“3 > 2”配对的最大 2。若更左侧出现更小值，就是“1”。下标顺序由反向扫描保证。

```go
func find132pattern(nums []int) bool {
    stack := []int{}
    middle := -int(^uint(0)>>1)-1 // int 最小值；不存在更小 int，不会误报。
    for i := len(nums)-1; i >= 0; i-- {
        if nums[i] < middle { return true }
        for len(stack) > 0 && stack[len(stack)-1] < nums[i] {
            middle = stack[len(stack)-1]
            stack = stack[:len(stack)-1]
        }
        stack = append(stack, nums[i])
    }
    return false
}
```

连续弹出的值递增；在尚未返回 true 的扫描中，更新不会降低已维护的有效 middle。三个大小关系均严格，不能把相等当作 132。时间、空间均 O(n)。

### 768. 最多能完成排序的块 II：合并的是块

[原题](https://leetcode.com/problems/max-chunks-to-make-sorted-ii/) · [源码](l768/768.go)

栈存各块最大值，非递减。当前值小于前块最大值，意味着跨块顺序冲突，必须向左合并所有冲突块。合并块最大值应保留为原最右块最大值，不能改成当前较小值。

```go
func maxChunksToSorted(arr []int) int {
    maxima := []int{}
    for _, x := range arr {
        if len(maxima) == 0 || maxima[len(maxima)-1] <= x {
            maxima = append(maxima, x)
            continue
        }
        mergedMax := maxima[len(maxima)-1]
        for len(maxima) > 0 && maxima[len(maxima)-1] > x {
            maxima = maxima[:len(maxima)-1]
        }
        maxima = append(maxima, mergedMax)
    }
    return len(maxima)
}
```

相等允许分块，因此弹出条件为 `>`。时间、空间均 O(n)。不要混淆 769：它额外保证输入为 `0..n-1` 的排列。

### 1130. 叶值的最小代价生成树：局部贪心消元

[原题](https://leetcode.com/problems/minimum-cost-tree-from-leaf-values/) · [源码](l1130/1130.go)

叶子中序顺序固定、叶值为正。维护严格递减栈，较小值遇到右侧不小于它的值时，可与两侧候选中更小者结合，计入乘积，较大的值继续代表子树最大叶值。直觉是较小值消失时尽量匹配较小代价；这需要题目结构支持，不是任意邻接乘积都能这样做。

```go
func mctFromLeafValues(arr []int) int {
    stack := []int{}
    cost := 0
    for _, x := range arr {
        for len(stack) > 0 && stack[len(stack)-1] <= x {
            middle := stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            neighbor := x
            if len(stack) > 0 && stack[len(stack)-1] < neighbor {
                neighbor = stack[len(stack)-1]
            }
            cost += middle*neighbor
        }
        stack = append(stack, x)
    }
    for len(stack) > 1 {
        last := len(stack)-1
        cost += stack[last]*stack[last-1]
        stack = stack[:last]
    }
    return cost
}
```

不要漏掉末尾递减栈的清算。时间、空间均 O(n)；也可先理解区间 DP 的 O(n³) 时间、O(n²) 空间解法，再学习贪心优化。

<a id="alternatives"></a>

## 9. G. 与单调队列、普通栈、排序扫描的区别

| 题目 | 核心状态与规则 | 为什么不必 / 不能照搬单调栈 |
| --- | --- | --- |
| [239. 滑动窗口最大值](https://leetcode.com/problems/sliding-window-maximum/) · [源码](../sliding_window/l239/239.go) | 双端队列存下标；队尾删除不大于新值的候选，队首删除 `index<=i-k` 的过期者，队首是当前最大值 | 既需尾部淘汰，又需头部过期删除。时间 O(n)，辅助空间 O(k) 需用环形队列或回收切片前缀；仅递增 head 保存历史可能占 O(n) |
| [856. 括号的分数](https://leetcode.com/problems/score-of-parentheses/) · [源码](../stack/l856/856.go) | 普通栈保存外层得分，右括号将内层转成 `max(1,2*内层分数)` 再合并 | 维护的是嵌套结构，没有值的单调不变量。时间 O(n)、栈空间 O(n)，也可用深度计数做到 O(1) 额外空间 |
| [853. 车队](https://leetcode.com/problems/car-fleet/) · [源码](l853/853.go) | 按位置从近到远排序，计算到达时间；`time<=前方车队时间` 则合并，否则新增车队 | 可写成栈，但只保留最近前方车队时间即可；总体 O(n log n)，排序后扫描 O(n)、扫描状态 O(1)。浮点除法前需转换类型 |

853 不能超车，恰好在终点追上也算同一车队，不能仅按速度排序。排序副本通常需要 O(n) 空间。239 用堆的简单懒删除方案最坏可占 O(n) 空间、耗时 O(n log n)，不要直接宣称 O(k) 堆空间。

<a id="pitfalls"></a>

## 10. 检查清单与学习路线

### 每道题先写清楚这六句话

1. 栈存下标、值，还是带聚合信息的记录？
2. 栈底到顶保持严格递增、非递减、严格递减还是非递增？
3. 当前元素为什么能让栈顶永久退出？是否需要保留它的聚合信息？
4. 弹栈结算旧元素、读取新栈顶回答当前元素，还是直接修改答案？
5. 弹栈后的新栈顶是什么边界？不存在时用什么默认值？
6. 相等元素归谁？是否清栈？乘积、模减法与返回类型会不会出错？

### 关键对照

| 场景 | 弹出者的意义 | 主要计算 / 相等处理 |
| --- | --- | --- |
| 739 / 496 | 等待答案的旧位置 | 严格更大才弹，距离为 `i-j` |
| 84 | 候选矩形高度 | 宽为 `R-L-1`；本文相等保留，早柱覆盖完整宽度 |
| 907 / 2104 | 被指定的最值代表 | 端点选择数为 `(i-L)*(R-i)`，一边严格一边非严格 |
| 42 | 槽底 | `min(两墙)-槽底`，弹出后必须还有左墙 |
| 402 / 316 | 较差的暂定前缀 | 更优新值＋仍能完成任务，不保证最终栈整体单调 |
| 901 / 768 | 已压缩的历史跨度 / 区块 | 弹出信息需要合并，不能直接丢弃 |

单调栈不等于排序：通常不能改变原来的相对位置。也不是每次 while 弹栈都用同一比较符号。空数组、单元素、全相等、严格递增、严格递减应作为基础测试；贡献类尤其要手算 `[2,2]`。

### 建议顺序

```text
最近边界：739 → 496 → 503 → 901 → 1475
区间扩展：84 → 85 → 1793 → 1856
贡献统计：907 → 2104 → 2281
贪心构造：402 → 316 / 1081 → 1673 → 321
结构结算：42
特殊不变量：962 → 768 → 456 → 1130
分类对照：239（队列）、856（普通栈）、853（排序扫描）
```

复习时先独立写状态和弹栈理由，再指出相较基础题改变的条件。**“新元素到来时，可从候选末尾连续且永久地淘汰或结算”是识别单调栈的主线，而不是题目里出现了“最大 / 最小”。**

<a id="implementation"></a>

## 11. 本地实现与变体说明

- 每份实际题解在自己的包内保留官方函数签名，不依赖仓库本地算法模板包。本文同类 helper 只声明一次，使用时复制依赖的 helper 与目标题函数即可。
- 496、84、907 的历史模板练习分别保存在 [l496/variants/template](l496/variants/template)、[l84/variants/template](l84/variants/template)、[l907/variants/template](l907/variants/template)，保留各自测试；503 的排列分类来源保存在 [l503/variants/from_permutation](l503/variants/from_permutation)。
- 856 已归到 [普通栈目录](../stack/l856)，原变体位于其 `variants/from_monotonic_stack`；239 已归到 [滑动窗口目录](../sliding_window/l239)。
- 1856、2281 目前仅有 legacy 练习，本文给出参考模板，不将历史代码当成已验证的主实现。本文代码可能与题解文件在 helper 命名或局部写法上不同，以本节边界约定理解即可。
- 仓库迁移记录见 [reorganization.json](../two_pointers/reorganization.json)。

参考：[LeetCode 单调栈题单](https://leetcode.com/tag/monotonic-stack/)。
